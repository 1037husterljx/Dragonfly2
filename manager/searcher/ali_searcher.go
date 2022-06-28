package searcher

import (
	"context"
	"fmt"
	"time"

	cachev8 "github.com/go-redis/cache/v8"
	"gorm.io/gorm"

	logger "d7y.io/dragonfly/v2/internal/dflog"
	"d7y.io/dragonfly/v2/manager/cache"
	"d7y.io/dragonfly/v2/manager/database"
	"d7y.io/dragonfly/v2/manager/model"
	"d7y.io/dragonfly/v2/manager/service"
	"d7y.io/dragonfly/v2/manager/types"
	"d7y.io/dragonfly/v2/pkg/rpc/manager"
	"d7y.io/dragonfly/v2/pkg/util"
)

var (
	defaultSecurityGroupName = "Buffer"

	rulesCache = &map[string][]model.SecurityGroup{}
)

type aliSearcher struct {
	*database.Database
	skylineAPI *service.SkylineAPI
	cache      *cache.Cache
}

func NewAliSearcher(db *database.Database, cache *cache.Cache) (Searcher, error) {
	skylineAPI, err := service.NewSkylineAPI()
	if err != nil {
		return nil, err
	}
	s := &aliSearcher{
		Database:   db,
		skylineAPI: skylineAPI,
		cache:      cache,
	}

	if err := s.refreshRulesCache(); err != nil {
		return nil, err
	}

	// 定时刷新rules缓存
	util.NewPeriodRoutine(
		context.Background(),
		time.Minute*5,
		func() {
			if err := s.refreshRulesCache(); err != nil {
				logger.Errorf("securityRules cache refresh failed:%v", err)
			}

		})

	logger.Info("aliSearcher init success")
	return s, nil
}

// FindSchedulerClusters 根据req查询cmdb得到hostInfo，获取相关securityRule,根据rule获取securityGroup, 返回securityGroup下的schedulers
func (s *aliSearcher) FindSchedulerClusters(ctx context.Context, empty []model.SchedulerCluster, req *manager.ListSchedulersRequest) ([]model.SchedulerCluster, error) {
	hostInfo, err := s.findHostInfo(ctx, req.HostInfo[ConditionSN], req.Ip, req.HostName)
	if err != nil {
		logger.Warnf("get hostInfo error :%v", err)
	}

	rules := s.getSecurityRules(ctx, hostInfo)

	return s.findSchedulersWithRules(rules)
}

// findSchedulersWithRules 按rules查找securityGroup
func (s *aliSearcher) findSchedulersWithRules(rules []string) ([]model.SchedulerCluster, error) {
	schedulerClusters := make([]model.SchedulerCluster, 0)
	for _, rule := range rules {
		if securityGroups, ok := (*rulesCache)[rule]; ok && len(securityGroups) > 0 {
			for _, sg := range securityGroups {
				if len(sg.SchedulerClusters) > 0 {
					for _, sc := range sg.SchedulerClusters {
						_ = s.DB.Find(&sc.Schedulers, "scheduler_cluster_id = ? and state = ?", sc.ID, "active").Error
						if len(sc.Schedulers) > 0 {
							schedulerClusters = append(schedulerClusters, sc)
						}
					}
				}
			}
			// 找到第一个scheduler不为空的schedulerCluster
			if len(schedulerClusters) > 0 {
				return schedulerClusters, nil
			}
		}
	}
	return nil, fmt.Errorf("can't find any securityGroup of rules:[%v]", rules)
}

// getSecurityRules 按优先级返回securityRules
// 例如:
// 0 : "ALI_TEST-NT12"
// 1 : "ALI_TEST"
// 2 : "Buffer"
func (s *aliSearcher) getSecurityRules(ctx context.Context, info *types.HostInfo) []string {
	if info == nil {
		return []string{defaultSecurityGroupName}
	}

	rules := make([]string, 0)
	if info.SecurityDomain != "" {
		if info.Idc != "" {
			rules = append(rules, info.SecurityDomain+"-"+info.Idc)
		}
		rules = append(rules, info.SecurityDomain)
	}
	rules = append(rules, defaultSecurityGroupName)
	return rules
}

// findHostInfo lookup hostInfo from cache or skylineAPI
func (s *aliSearcher) findHostInfo(ctx context.Context, sn, ip, hostname string) (*types.HostInfo, error) {
	var hostInfo *types.HostInfo
	lookups := map[string]string{
		service.HostInfoKyeSN:       sn,
		service.HostInfoKeyIP:       ip,
		service.HostInfoKeyHostname: hostname,
	}
	// 依次查询sn、ip、hostname
	for lookupKey, lookupValue := range lookups {
		if lookupValue == "" {
			continue
		}
		cachekey := cache.MakeCacheKey(cache.HostInfoNamespace, lookupKey+lookupValue)

		if err := s.cache.Get(ctx, cachekey, hostInfo); err == nil {
			return hostInfo, nil
		}

		hostInfo, err := s.skylineAPI.GetHostInfo(lookupKey, lookupValue)

		if err == nil {
			_ = s.cache.Once(&cachev8.Item{
				Ctx:   ctx,
				Key:   cachekey,
				Value: hostInfo,
				TTL:   time.Hour * 24,
			})
			return hostInfo, nil
		}
		logger.Debugf("get hostinfo of %s:%s failed", lookupKey, lookupValue)
	}
	return nil, fmt.Errorf("can't find hostinfo of {sn: %s, ip: %s, hostname: %s} from skyline", sn, ip, hostname)
}

// refreshRulesCache 生成security rules 内存缓存
// 因为schedulers是动态变化的，所以只缓存到schedulerCluster层
func (s *aliSearcher) refreshRulesCache() error {
	rules := &[]model.SecurityRule{}
	err := s.DB.
		Preload("SecurityGroups").
		Preload("SecurityGroups.SchedulerClusters").
		Find(rules).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	cache := map[string][]model.SecurityGroup{}
	for _, rule := range *rules {
		if len(rule.SecurityGroups) > 0 {
			cache[rule.Domain] = rule.SecurityGroups
		}
	}
	rulesCache = &cache
	return nil
}
