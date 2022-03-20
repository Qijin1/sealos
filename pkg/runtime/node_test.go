/*
Copyright 2022 cuisongliu@qq.com.

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

package runtime

import (
	"testing"

	"github.com/fanux/sealos/pkg/utils/logger"
)

func TestNode(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "default",
			wantErr: false,
		},
	}
	//logger.Cfg(true,false)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ForceDelete = true
			logger.Cfg(true, false)
			k, _ := NewDefaultRuntime("default")
			if err := k.JoinMasters([]string{"192.168.64.18"}); (err != nil) != tt.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
			}
			t.Log("process success")
		})
	}
}