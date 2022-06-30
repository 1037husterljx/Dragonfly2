/*
 *     Copyright 2020 The Dragonfly Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package daemon

import (
	"errors"
	"net"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"d7y.io/dragonfly/v2/client/config"
	logger "d7y.io/dragonfly/v2/internal/dflog"
	"d7y.io/dragonfly/v2/manager/searcher"
	"d7y.io/dragonfly/v2/pkg/dfnet"
	"d7y.io/dragonfly/v2/pkg/net/fqdn"
	"d7y.io/dragonfly/v2/pkg/net/ip"
	"d7y.io/dragonfly/v2/pkg/net/sn"
	"d7y.io/dragonfly/v2/pkg/reachable"
	"d7y.io/dragonfly/v2/pkg/rpc/manager"
	"d7y.io/dragonfly/v2/pkg/rpc/manager/client"
	"d7y.io/dragonfly/v2/pkg/rpc/manager/client/mocks"
)

func TestDaemonSchedulersToAvailableNetAddrs(t *testing.T) {
	l, err := net.Listen("tcp", ":3000")
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()

	tests := []struct {
		name       string
		schedulers []*manager.Scheduler
		expect     func(t *testing.T, addrs []dfnet.NetAddr)
	}{
		{
			name: "available ip",
			schedulers: []*manager.Scheduler{
				{
					Ip:                 "127.0.0.1",
					Port:               int32(3000),
					SchedulerClusterId: 1,
				},
				{
					Ip:                 "127.0.0.1",
					Port:               int32(3001),
					SchedulerClusterId: 1,
				},
			},
			expect: func(t *testing.T, addrs []dfnet.NetAddr) {
				assert := assert.New(t)
				assert.EqualValues(addrs, []dfnet.NetAddr{{Type: dfnet.TCP, Addr: "127.0.0.1:3000"}})
			},
		},
		{
			name: "available host",
			schedulers: []*manager.Scheduler{
				{
					Ip:                 "foo",
					HostName:           "localhost",
					Port:               int32(3000),
					SchedulerClusterId: 1,
				},
				{
					Ip:                 "foo",
					HostName:           "localhost",
					Port:               int32(3001),
					SchedulerClusterId: 1,
				},
			},
			expect: func(t *testing.T, addrs []dfnet.NetAddr) {
				assert := assert.New(t)
				assert.EqualValues(addrs, []dfnet.NetAddr{{Type: dfnet.TCP, Addr: "localhost:3000"}})
			},
		},
		{
			name: "available ip and host",
			schedulers: []*manager.Scheduler{
				{
					Ip:                 "foo",
					HostName:           "localhost",
					Port:               int32(3000),
					SchedulerClusterId: 1,
				},
				{
					Ip:                 "foo",
					HostName:           "localhost",
					Port:               int32(3001),
					SchedulerClusterId: 1,
				},
				{
					Ip:                 "127.0.0.1",
					HostName:           "foo",
					Port:               int32(3001),
					SchedulerClusterId: 1,
				},
				{
					Ip:                 "127.0.0.1",
					HostName:           "foo",
					Port:               int32(3000),
					SchedulerClusterId: 1,
				},
				{
					Ip:                 "127.0.0.1",
					HostName:           "foo",
					Port:               int32(3001),
					SchedulerClusterId: 1,
				},
			},
			expect: func(t *testing.T, addrs []dfnet.NetAddr) {
				assert := assert.New(t)
				assert.EqualValues(addrs, []dfnet.NetAddr{
					{Type: dfnet.TCP, Addr: "localhost:3000"},
					{Type: dfnet.TCP, Addr: "127.0.0.1:3000"},
				})
			},
		},
		{
			name: "unreachable",
			schedulers: []*manager.Scheduler{
				{
					Ip:                 "foo",
					HostName:           "localhost",
					Port:               int32(3001),
					SchedulerClusterId: 1,
				},
				{
					Ip:                 "127.0.0.1",
					HostName:           "foo",
					Port:               int32(3001),
					SchedulerClusterId: 1,
				},
			},
			expect: func(t *testing.T, addrs []dfnet.NetAddr) {
				assert := assert.New(t)
				assert.EqualValues(addrs, []dfnet.NetAddr{})
			},
		},
		{
			name:       "empty schedulers",
			schedulers: []*manager.Scheduler{},
			expect: func(t *testing.T, addrs []dfnet.NetAddr) {
				assert := assert.New(t)
				assert.EqualValues(addrs, []dfnet.NetAddr{})
			},
		},
		{
			name: "available ip with different scheduler cluster",
			schedulers: []*manager.Scheduler{
				{
					Ip:                 "127.0.0.1",
					HostName:           "foo",
					Port:               int32(3000),
					SchedulerClusterId: 1,
				},
				{
					Ip:                 "127.0.0.1",
					HostName:           "foo",
					Port:               int32(3000),
					SchedulerClusterId: 1,
				},
				{
					Ip:                 "127.0.0.1",
					HostName:           "foo",
					Port:               int32(3001),
					SchedulerClusterId: 2,
				},
			},
			expect: func(t *testing.T, addrs []dfnet.NetAddr) {
				assert := assert.New(t)
				assert.EqualValues(addrs, []dfnet.NetAddr{
					{Type: dfnet.TCP, Addr: "127.0.0.1:3000"},
					{Type: dfnet.TCP, Addr: "127.0.0.1:3000"},
				})
			},
		},
		{
			name: "available host with different scheduler cluster",
			schedulers: []*manager.Scheduler{
				{
					Ip:                 "127.0.0.1",
					HostName:           "foo",
					Port:               int32(3001),
					SchedulerClusterId: 1,
				},
				{
					Ip:                 "foo",
					HostName:           "localhost",
					Port:               int32(3000),
					SchedulerClusterId: 2,
				},
				{
					Ip:                 "foo",
					HostName:           "localhost",
					Port:               int32(3000),
					SchedulerClusterId: 2,
				},
			},
			expect: func(t *testing.T, addrs []dfnet.NetAddr) {
				assert := assert.New(t)
				assert.EqualValues(addrs, []dfnet.NetAddr{
					{Type: dfnet.TCP, Addr: "localhost:3000"},
					{Type: dfnet.TCP, Addr: "localhost:3000"},
				})
			},
		},
		{
			name: "available host and ip with different scheduler cluster",
			schedulers: []*manager.Scheduler{
				{
					Ip:                 "foo",
					HostName:           "localhost",
					Port:               int32(3000),
					SchedulerClusterId: 2,
				},
				{
					Ip:                 "127.0.0.1",
					HostName:           "foo",
					Port:               int32(3001),
					SchedulerClusterId: 1,
				},
				{
					Ip:                 "127.0.0.1",
					HostName:           "goo",
					Port:               int32(3000),
					SchedulerClusterId: 2,
				},
				{
					Ip:                 "127.0.0.1",
					HostName:           "foo",
					Port:               int32(3000),
					SchedulerClusterId: 1,
				},
			},
			expect: func(t *testing.T, addrs []dfnet.NetAddr) {
				assert := assert.New(t)
				assert.EqualValues(addrs, []dfnet.NetAddr{
					{Type: dfnet.TCP, Addr: "localhost:3000"},
					{Type: dfnet.TCP, Addr: "127.0.0.1:3000"},
				})
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.expect(t, schedulersToAvailableNetAddrs(tc.schedulers))
		})
	}
}

func TestNewWithAddrs(t *testing.T) {
	l, err := net.Listen("tcp", ":3000")
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()

	type args struct {
		netAddrs       []dfnet.NetAddr
		SelectStrategy string
		hostOption     config.HostOption
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ipcheck",
			args: args{
				netAddrs: []dfnet.NetAddr{
					{
						Type: dfnet.TCP,
						Addr: "16.6.6.6:4538",
					},
				},
				SelectStrategy: config.ManagerSelectStrategyManagerReady,
				hostOption: config.HostOption{
					Hostname:       fqdn.FQDNHostname,
					ListenIP:       net.IPv4zero.String(),
					AdvertiseIP:    ip.IPv4,
					SN:             sn.SN,
					SecurityDomain: "",
					Location:       "",
					IDC:            "",
					NetTopology:    "",
				},
			},
			wantErr: true,
		},
		{
			name: "ipcheck2",
			args: args{
				netAddrs: []dfnet.NetAddr{
					{
						Type: dfnet.TCP,
						Addr: "16.6.6.6:7868",
					},
					{
						Type: dfnet.TCP,
						Addr: "127.0.0.1:3000",
					},
				},
				SelectStrategy: config.ManagerSelectStrategyManagerReady,
				hostOption: config.HostOption{
					Hostname:       fqdn.FQDNHostname,
					ListenIP:       net.IPv4zero.String(),
					AdvertiseIP:    ip.IPv4,
					SN:             sn.SN,
					SecurityDomain: "",
					Location:       "",
					IDC:            "",
					NetTopology:    "",
				},
			},
			wantErr: false,
		},
		{
			name: "SchedulerReady",
			args: args{
				netAddrs: []dfnet.NetAddr{
					{
						Type: dfnet.TCP,
						Addr: "16.6.6.6:5647",
					},
					{
						Type: dfnet.TCP,
						Addr: "127.0.0.1:3000",
					},
					{
						Type: dfnet.TCP,
						Addr: "127.0.0.1:3000",
					},
				},
				SelectStrategy: config.ManagerSelectStrategySchedulerReady,
				hostOption: config.HostOption{
					Hostname:       fqdn.FQDNHostname,
					ListenIP:       net.IPv4zero.String(),
					AdvertiseIP:    ip.IPv4,
					SN:             sn.SN,
					SecurityDomain: "",
					Location:       "",
					IDC:            "",
					NetTopology:    "",
				},
			},
			wantErr: false,
		},
		{
			name: "configcheck",
			args: args{
				netAddrs: []dfnet.NetAddr{
					{
						Type: dfnet.TCP,
						Addr: "127.0.0.1:3000",
					},
				},
				SelectStrategy: "xxxx",
				hostOption: config.HostOption{
					Hostname:       fqdn.FQDNHostname,
					ListenIP:       net.IPv4zero.String(),
					AdvertiseIP:    ip.IPv4,
					SN:             sn.SN,
					SecurityDomain: "",
					Location:       "",
					IDC:            "",
					NetTopology:    "",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewManagerWithAddrs_test(t, tt.args.netAddrs, tt.args.SelectStrategy, tt.args.hostOption)
			logger.Info(err)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewWithAddrs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

// NewWithAddrs creates manager client with addresses.
func NewManagerWithAddrs_test(t *testing.T, netAddrs []dfnet.NetAddr, SelectStrategy string, hostOption config.HostOption) (client.Client, error) {
	for i, netAddr := range netAddrs {
		reachable := reachable.New(&reachable.Config{Address: netAddr.Addr})
		if err := reachable.Check(); err == nil {
			if SelectStrategy == config.ManagerSelectStrategyManagerReady {
				logger.Infof("use %s address for manager grpc client", netAddr.Addr)
				return nil, nil
			} else if SelectStrategy == config.ManagerSelectStrategySchedulerReady {
				ctl := gomock.NewController(t)
				defer ctl.Finish()
				mockManagerClient := mocks.NewMockClient(ctl)
				if i == 0 {
					gomock.InOrder(
						mockManagerClient.EXPECT().ListSchedulers(gomock.Any()).Return(&manager.ListSchedulersResponse{Schedulers: []*manager.Scheduler{}}, nil).Times(1),
						mockManagerClient.EXPECT().Close().Return(nil).AnyTimes(),
					)
				} else if i == 1 {
					gomock.InOrder(
						mockManagerClient.EXPECT().ListSchedulers(gomock.Any()).Return(&manager.ListSchedulersResponse{
							Schedulers: []*manager.Scheduler{
								{
									Ip:   "127.0.0.1",
									Port: 4000,
								},
							},
						}, nil).Times(1),
						mockManagerClient.EXPECT().Close().Return(nil).AnyTimes(),
					)
				} else {
					gomock.InOrder(
						mockManagerClient.EXPECT().ListSchedulers(gomock.Any()).Return(&manager.ListSchedulersResponse{
							Schedulers: []*manager.Scheduler{
								{
									Ip:   "127.0.0.1",
									Port: 3000,
								},
							},
						}, nil).Times(1),
						mockManagerClient.EXPECT().Close().Return(nil).AnyTimes(),
					)
				}

				if managerClient, err := getManagerWithValidSchedulers_test(mockManagerClient, hostOption); err == nil {
					logger.Infof("use %s address for manager grpc client", netAddr.Addr)
					return managerClient, nil
				}
			} else {
				logger.Warn("wrong SelectStrategy")
				return nil, errors.New("wrong healthCheckMethod config")
			}
		}
		logger.Warnf("%s manager address can not reachable", netAddr.Addr)
	}

	return nil, errors.New("can not find available manager addresses")
}

// healthCheckMethod:ScheduleReady
func getManagerWithValidSchedulers_test(managerClient client.Client, hostOption config.HostOption) (client.Client, error) {

	listSchedulersResp, err := managerClient.ListSchedulers(&manager.ListSchedulersRequest{
		SourceType: manager.SourceType_PEER_SOURCE,
		HostName:   hostOption.Hostname,
		Ip:         hostOption.AdvertiseIP,
		HostInfo: map[string]string{
			searcher.ConditionSN:             hostOption.SN,
			searcher.ConditionSecurityDomain: hostOption.SecurityDomain,
			searcher.ConditionIDC:            hostOption.IDC,
			searcher.ConditionNetTopology:    hostOption.NetTopology,
			searcher.ConditionLocation:       hostOption.Location,
		},
	})
	if err != nil {
		return nil, err
	}

	if len(schedulersToAvailableNetAddrs(listSchedulersResp.Schedulers)) == 0 {
		managerClient.Close()
		logger.Warnf("listSchedulersResp.Schedulers:%v", listSchedulersResp.Schedulers)
		return nil, errors.New("manager returns Schedule as empty")
	}
	logger.Infof("listSchedulersResp.Schedulers:%v", listSchedulersResp.Schedulers)
	return managerClient, nil
}
