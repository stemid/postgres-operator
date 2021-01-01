/*
 Copyright 2021 Crunchy Data Solutions, Inc.
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

package naming

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

// AsSelector is a wrapper around metav1.LabelSelectorAsSelector() which converts
// the LabelSelector API type into something that implements labels.Selector.
func AsSelector(s metav1.LabelSelector) (labels.Selector, error) {
	return metav1.LabelSelectorAsSelector(&s)
}

// AnyCluster selects things for any PostgreSQL cluster.
func AnyCluster() metav1.LabelSelector {
	return metav1.LabelSelector{
		MatchExpressions: []metav1.LabelSelectorRequirement{
			{Key: LabelCluster, Operator: metav1.LabelSelectorOpExists},
		},
	}
}

// ClusterInstances selects things for PostgreSQL instances in cluster.
func ClusterInstances(cluster string) metav1.LabelSelector {
	return metav1.LabelSelector{
		MatchLabels: map[string]string{
			LabelCluster: cluster,
		},
		MatchExpressions: []metav1.LabelSelectorRequirement{
			{Key: LabelInstance, Operator: metav1.LabelSelectorOpExists},
		},
	}
}

// ClusterReplicas selects things for PostgreSQL replicas in cluster.
func ClusterReplicas(cluster string) metav1.LabelSelector {
	s := ClusterInstances(cluster)
	s.MatchLabels[LabelRole] = "replica"
	return s
}