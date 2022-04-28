// Copyright 2021 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License

package util

import commonv1 "github.com/kubeflow/common/pkg/apis/common/v1"

const (
	DefaultGangSchedulerName = "volcano"
)

func IsGangSchedulerSet(replicas map[commonv1.ReplicaType]*commonv1.ReplicaSpec, schedulerName string) bool {
	if len(schedulerName) == 0 {
		schedulerName = DefaultGangSchedulerName
	}
	allReplicasHaveGangSchedulerSet := true
	for _, spec := range replicas {
		if !(spec.Template.Spec.SchedulerName != "" && spec.Template.Spec.SchedulerName == schedulerName) {
			allReplicasHaveGangSchedulerSet = false
		}
	}

	return allReplicasHaveGangSchedulerSet
}
