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

package types

// HostInfo 数据字段 see https://yuque.antfin.com/baserun/isee_user_manual/la3480
type HostInfo struct {
	IP             string   `json:"ip,omitempty"`
	Sn             string   `json:"sn,omitempty"`
	HostName       string   `json:"host_name,omitempty"`
	Idc            string   `json:"idc.abbreviation,omitempty"` // 机房
	SecurityDomain string   `json:"security_domain,omitempty"`
	VpcID          string   `json:"vpc_id,omitempty"`
	Site           string   `json:"cabinet.logic_region,omitempty"` // 逻辑站点，机房内划分多个逻辑站点
	Room           string   `json:"room.abbreviation,omitempty"`
	Rack           string   `json:"cabinet.cabinet_num,omitempty"` // 机柜信息
	DswName        string   `json:"dsw_cluster.name,omitempty"`
	LogicPodName   string   `json:"logic_pod.name,omitempty"`
	PodName        string   `json:"net_pod.name,omitempty"`
	AswName        string   `json:"net_asw.name,omitempty"`
	Country        string   `json:"idc.country,omitempty"`
	Area           string   `json:"idc.area,omitempty"`
	Province       string   `json:"idc.province,omitempty"`
	City           string   `json:"idc.city,omitempty"`
	Tags           []string `json:"tags,omitempty"`
}
