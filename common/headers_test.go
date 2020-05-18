/**
 * (C) Copyright IBM Corp. 2019.
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

package common

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestGetSystemInfo(t *testing.T) {
	var sysinfo = GetSystemInfo()
	assert.NotNil(t, sysinfo)
	assert.True(t, strings.Contains(sysinfo, "lang="))
	assert.True(t, strings.Contains(sysinfo, "arch="))
	assert.True(t, strings.Contains(sysinfo, "os="))
	assert.True(t, strings.Contains(sysinfo, "go.version="))
}

func TestGetSdkHeaders(t *testing.T) {
	var headers = GetSdkHeaders("myService", "v123", "myOperation")
	assert.NotNil(t, headers)

	var foundIt bool

	_, foundIt = headers[headerNameUserAgent]
	assert.True(t, foundIt)
	t.Logf("user agent: %s\n", headers[headerNameUserAgent])
}