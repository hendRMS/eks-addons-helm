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

package controller

import (
	"k8s.io/api/core/v1"

	"github.com/kubernetes-sigs/cluster-proportional-autoscaler/pkg/autoscaler/k8sclient"
)

// Controller defines the interface every controller should implement
type Controller interface {
	// GetExpectedReplicas returns the expected replicas based on cluster status
	GetExpectedReplicas(*k8sclient.ClusterStatus) (int32, error)
	// SyncConfig syncs the ConfigMap with controller
	SyncConfig(*v1.ConfigMap) error
	// GetParamsVersion returns the latest parameters version from controller
	GetParamsVersion() string
	// GetControllerType returns the controller type
	GetControllerType() string
}
