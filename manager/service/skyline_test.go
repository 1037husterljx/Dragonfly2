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
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSkyline(t *testing.T) {
	var domain, appName, account, accessKey = "skyline_domain", "skyline_appName", "skyline_account", "skyline_accessKey"
	os.Setenv(domain, "sky.alibaba-inc.com")
	os.Setenv(appName, "dragonfly")
	os.Setenv(account, "dragonfly-skyline")
	os.Setenv(accessKey, "FTvjhscxJq3P0sMg")
	defer func() {
		os.Unsetenv(domain)
		os.Unsetenv(appName)
		os.Unsetenv(account)
		os.Unsetenv(accessKey)
	}()
	skyline, err := NewSkylineAPI()
	assert.Nil(t, err)
	host1, err := skyline.GetHostInfoBySN("11F6ML1")
	assert.Nil(t, err)
	fmt.Println(host1)
	host2, err := skyline.GetHostInfoByIP("11.160.191.32")
	assert.Nil(t, err)
	fmt.Println(host2)
	host3, err := skyline.GetHostInfoByHostName("i22g06223.eu95sqa")
	assert.Nil(t, err)
	fmt.Println(host3)
	assert.EqualValues(t, host1, host2)
	assert.EqualValues(t, host1, host3)
}
