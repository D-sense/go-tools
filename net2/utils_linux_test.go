// Copyright 2019 xgfone
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package net2

import (
	"testing"
)

func TestGetIP(t *testing.T) {
	if ips, err := GetIP("127.0.0.1"); err != nil || len(ips) != 1 || ips[0] != "127.0.0.1" {
		t.Fail()
	}

	if ips, err := GetIP("lo"); err != nil || ips[0] != "127.0.0.1" {
		t.Fail()
	}
}

func TestGetInterfaceByIP(t *testing.T) {
	if iface, err := GetInterfaceByIP("127.0.0.1"); err != nil || iface != "lo" {
		t.Fail()
	}
}
