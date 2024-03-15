// Copyright 2019 Altinity Ltd and/or its affiliates. All rights reserved.
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

package creator

import (
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"

	api "github.com/altinity/clickhouse-operator/pkg/apis/clickhouse.altinity.com/v1"
)

func getOwnerReferences(chi *api.ClickHouseInstallation) []meta.OwnerReference {
	if chi.Runtime.Attributes.SkipOwnerRef {
		return nil
	}
	return _getOwnerReference(&chi.ObjectMeta)
}

func _getOwnerReference(objectMeta *meta.ObjectMeta) []meta.OwnerReference {
	controller := true
	block := true
	return []meta.OwnerReference{
		{
			APIVersion:         api.SchemeGroupVersion.String(),
			Kind:               api.ClickHouseInstallationCRDResourceKind,
			Name:               objectMeta.GetName(),
			UID:                objectMeta.GetUID(),
			Controller:         &controller,
			BlockOwnerDeletion: &block,
		},
	}
}
