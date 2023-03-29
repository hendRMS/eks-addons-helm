/*
Copyright 2016 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package plugin

import (
	"testing"

	"k8s.io/api/core/v1"
)

func TestEnsureController(t *testing.T) {
	testCases := []struct {
		configMap *v1.ConfigMap
		expError  bool
	}{
		{
			&v1.ConfigMap{
				Data: map[string]string{
					"invalidmode": "",
				},
			},
			true,
		},
		{
			&v1.ConfigMap{
				Data: map[string]string{
					"toomanyentries1": "",
					"toomanyentries2": "",
				},
			},
			true,
		},
		{
			&v1.ConfigMap{
				Data: map[string]string{
					"linear": "{\"nodesPerReplica\":1}",
				},
			},
			false,
		},
	}

	for _, tc := range testCases {
		_, err := EnsureController(nil, tc.configMap)
		if err != nil && !tc.expError {
			t.Errorf("Expect no error, got error for configMap %v, error msg: %v", tc.configMap, err)
			continue
		} else if err == nil && tc.expError {
			t.Errorf("Expect error, got no error for configMap %v, error msg: %v", tc.configMap, err)
			continue
		}
	}
}
