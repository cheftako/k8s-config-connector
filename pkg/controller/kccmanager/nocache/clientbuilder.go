// Copyright 2022 Google LLC
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
// limitations under the License.

package nocache

import (
	opv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"

	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var NoCacheClientFunc = func(config *rest.Config, options client.Options) (client.Client, error) {
	options.Cache = nil
	return client.New(config, options)
}

// Fine grained cache controls for ConfigConnector and ConfigConnectorContext.
var ByCCandCCC = map[client.Object]cache.ByObject{
	&opv1beta1.ConfigConnector{}:        {},
	&opv1beta1.ConfigConnectorContext{}: {},
}
