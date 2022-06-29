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

package service

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-http-utils/headers"

	"d7y.io/dragonfly/v2/manager/types"
	d7ystrings "d7y.io/dragonfly/v2/pkg/strings"
)

var hostInfo = []string{
	"ip", "sn", "host_name", "security_domain", "idc.abbreviation",
	//"tags", "vpc_id", "idc.country", "idc.area", "idc.province", "idc.city",
	//"room.abbreviation", "cabinet.cabinet_num", "cabinet.logic_region",
	//"dsw_cluster.name", "net_asw.name", "logic_pod.name", "net_pod.name",
}

const (
	HostInfoKyeSN       = "sn"
	HostInfoKeyIP       = "ip"
	HostInfoKeyHostname = "hostname"
)

type SkylineAPI struct {
	skylineClient
}

func NewSkylineAPI() (*SkylineAPI, error) {
	var skylineDomain = os.Getenv("skyline_domain")
	if skylineDomain == "" {
		return nil, fmt.Errorf("skyline domain environment variable is not set")
	}
	var skylineAppName = os.Getenv("skyline_appName")
	if skylineAppName == "" {
		return nil, fmt.Errorf("skyline appName environment variable is not set")
	}
	var skylineAccount = os.Getenv("skyline_account")
	if skylineAccount == "" {
		return nil, fmt.Errorf("skyline account environment variable is not set")
	}
	var skylineAccessKey = os.Getenv("skyline_accessKey")
	if skylineAccessKey == "" {
		return nil, fmt.Errorf("skyline accessKey environment variable is not set")
	}
	return &SkylineAPI{skylineClient{
		domain:    skylineDomain,
		appName:   skylineAppName,
		account:   skylineAccount,
		accessKey: skylineAccessKey,
		Client:    http.DefaultClient,
	}}, nil
}

func (skyline *SkylineAPI) GetHostInfoBySN(sn string) (*types.HostInfo, error) {
	return skyline.GetHostInfo(HostInfoKyeSN, sn)
}

func (skyline *SkylineAPI) GetHostInfoByIP(ip string) (*types.HostInfo, error) {
	return skyline.GetHostInfo(HostInfoKeyIP, ip)
}

func (skyline *SkylineAPI) GetHostInfoByHostName(hostname string) (*types.HostInfo, error) {
	return skyline.GetHostInfo(HostInfoKeyHostname, hostname)
}

func (skyline *SkylineAPI) GetHostInfo(key string, value string) (*types.HostInfo, error) {
	if d7ystrings.IsBlank(key) || d7ystrings.IsBlank(value) {
		return nil, fmt.Errorf("key: %s or value: %s is empty", key, value)
	}
	query := &QueryItem{
		From:      "server",
		Select:    strings.Join(hostInfo, ","),
		Condition: fmt.Sprintf("%s='%s'", key, value),
		NeedTotal: true,
		Page:      1,
		Num:       20,
	}
	resp, err := skyline.skylineClient.search(query)
	if err != nil {
		return nil, err
	}
	if !resp.Success {
		return nil, fmt.Errorf("errCode:%d, errMsg:%s", resp.ErrorCode, resp.ErrorMessage)
	}
	if resp.Data.TotalCount != 1 || resp.Data.HasMore {
		return nil, fmt.Errorf("want only one but got %d number of host", resp.Data.TotalCount)
	}
	hostJSON, err := json.Marshal(resp.Data.ItemList[0])
	if err != nil {
		return nil, err
	}
	var hostInfo = new(types.HostInfo)
	if err := json.Unmarshal(hostJSON, hostInfo); err != nil {
		return nil, err
	}
	return hostInfo, nil
}

type skylineClient struct {
	domain    string
	appName   string
	account   string
	accessKey string
	*http.Client
}

type Auth struct {
	AppName   string `json:"appName"`
	Account   string `json:"account"`
	Timestamp int64  `json:"timestamp"`
	Signature string `json:"signature"`
}

type QueryItem struct {
	From      string `json:"from"`
	Select    string `json:"select"`
	Condition string `json:"condition"`
	GroupBy   string `json:"groupBy"`
	NeedTotal bool   `json:"needTotal"`
	Page      int    `json:"page"`
	Num       int    `json:"num"`
}

type QueryRequest struct {
	Auth *Auth      `json:"auth"`
	Item *QueryItem `json:"item"`
}

type ResponseData struct {
	TotalCount int                 `json:"totalCount"`
	HasMore    bool                `json:"hasMore"`
	ItemList   []map[string]string `json:"itemList"`
}

type QueryResponse struct {
	Success      bool          `json:"success"`
	Data         *ResponseData `json:"value"`
	ErrorCode    int           `json:"errorCode"`
	ErrorCodeMsg string        `json:"errorCodeMsg"`
	ErrorMessage string        `json:"errorMessage"`
}

func (client *skylineClient) search(query *QueryItem) (*QueryResponse, error) {
	now := time.Now().Unix()
	queryRequest := &QueryRequest{
		Auth: &Auth{
			AppName:   client.appName,
			Account:   client.account,
			Timestamp: now,
			Signature: client.signature(now),
		},
		Item: query,
	}
	b, err := json.Marshal(queryRequest)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%s/item/query", client.domain), bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	req.Header.Set(headers.ContentType, "application/json")
	req.Header.Set("account", client.account)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var queryResponse = new(QueryResponse)
	if err := json.Unmarshal(buf, queryResponse); err != nil {
		return nil, err
	}
	return queryResponse, nil
}

// signature Signature for skyline
func (client *skylineClient) signature(timestamp int64) string {
	str := fmt.Sprintf("%s%s%d", client.account, client.accessKey, timestamp)
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
