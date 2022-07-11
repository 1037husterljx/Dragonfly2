//go:build linux
// +build linux

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

package dfpath

var DefaultWorkHome = "/home/staragent/plugins/dragonfly2/local_data"
var DefaultCacheDir = "/var/cache/dragonfly"
var DefaultConfigDir = "/etc/dragonfly"

// DefaultAliConfigDir ali内部SA插件形式部署
var DefaultAliConfigDir = "/home/staragent/plugins/dragonfly2/config"
var DefaultLogDir = "/var/log/dragonfly"
var DefaultDataDir = "/var/lib/dragonfly"
var DefaultPluginDir = "/usr/local/dragonfly/plugins"
