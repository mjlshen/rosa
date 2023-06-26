/*
Copyright (c) 2020 Red Hat, Inc.

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

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

// ClusterAutoscaler represents the values of the 'cluster_autoscaler' type.
//
// Cluster-wide autoscaling configuration.
type ClusterAutoscaler struct {
	bitmap_                   uint32
	logVerbosity              int
	maxPodGracePeriod         int
	resourceLimits            *AutoscalerResourceLimits
	scaleDown                 *AutoscalerScaleDownConfig
	balanceSimilarNodeGroups  bool
	skipNodesWithLocalStorage bool
}

// Empty returns true if the object is empty, i.e. no attribute has a value.
func (o *ClusterAutoscaler) Empty() bool {
	return o == nil || o.bitmap_ == 0
}

// BalanceSimilarNodeGroups returns the value of the 'balance_similar_node_groups' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// BalanceSimilarNodeGroups enables/disables the
// `--balance-similar-node-groups` cluster-autocaler feature.
// This feature will automatically identify node groups with
// the same instance type and the same set of labels and try
// to keep the respective sizes of those node groups balanced.
func (o *ClusterAutoscaler) BalanceSimilarNodeGroups() bool {
	if o != nil && o.bitmap_&1 != 0 {
		return o.balanceSimilarNodeGroups
	}
	return false
}

// GetBalanceSimilarNodeGroups returns the value of the 'balance_similar_node_groups' attribute and
// a flag indicating if the attribute has a value.
//
// BalanceSimilarNodeGroups enables/disables the
// `--balance-similar-node-groups` cluster-autocaler feature.
// This feature will automatically identify node groups with
// the same instance type and the same set of labels and try
// to keep the respective sizes of those node groups balanced.
func (o *ClusterAutoscaler) GetBalanceSimilarNodeGroups() (value bool, ok bool) {
	ok = o != nil && o.bitmap_&1 != 0
	if ok {
		value = o.balanceSimilarNodeGroups
	}
	return
}

// LogVerbosity returns the value of the 'log_verbosity' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Sets the autoscaler log level.
// Default value is 1, level 4 is recommended for DEBUGGING and level 6 will enable almost everything.
func (o *ClusterAutoscaler) LogVerbosity() int {
	if o != nil && o.bitmap_&2 != 0 {
		return o.logVerbosity
	}
	return 0
}

// GetLogVerbosity returns the value of the 'log_verbosity' attribute and
// a flag indicating if the attribute has a value.
//
// Sets the autoscaler log level.
// Default value is 1, level 4 is recommended for DEBUGGING and level 6 will enable almost everything.
func (o *ClusterAutoscaler) GetLogVerbosity() (value int, ok bool) {
	ok = o != nil && o.bitmap_&2 != 0
	if ok {
		value = o.logVerbosity
	}
	return
}

// MaxPodGracePeriod returns the value of the 'max_pod_grace_period' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Gives pods graceful termination time before scaling down.
func (o *ClusterAutoscaler) MaxPodGracePeriod() int {
	if o != nil && o.bitmap_&4 != 0 {
		return o.maxPodGracePeriod
	}
	return 0
}

// GetMaxPodGracePeriod returns the value of the 'max_pod_grace_period' attribute and
// a flag indicating if the attribute has a value.
//
// Gives pods graceful termination time before scaling down.
func (o *ClusterAutoscaler) GetMaxPodGracePeriod() (value int, ok bool) {
	ok = o != nil && o.bitmap_&4 != 0
	if ok {
		value = o.maxPodGracePeriod
	}
	return
}

// ResourceLimits returns the value of the 'resource_limits' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Constraints of autoscaling resources.
func (o *ClusterAutoscaler) ResourceLimits() *AutoscalerResourceLimits {
	if o != nil && o.bitmap_&8 != 0 {
		return o.resourceLimits
	}
	return nil
}

// GetResourceLimits returns the value of the 'resource_limits' attribute and
// a flag indicating if the attribute has a value.
//
// Constraints of autoscaling resources.
func (o *ClusterAutoscaler) GetResourceLimits() (value *AutoscalerResourceLimits, ok bool) {
	ok = o != nil && o.bitmap_&8 != 0
	if ok {
		value = o.resourceLimits
	}
	return
}

// ScaleDown returns the value of the 'scale_down' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Configuration of scale down operation.
func (o *ClusterAutoscaler) ScaleDown() *AutoscalerScaleDownConfig {
	if o != nil && o.bitmap_&16 != 0 {
		return o.scaleDown
	}
	return nil
}

// GetScaleDown returns the value of the 'scale_down' attribute and
// a flag indicating if the attribute has a value.
//
// Configuration of scale down operation.
func (o *ClusterAutoscaler) GetScaleDown() (value *AutoscalerScaleDownConfig, ok bool) {
	ok = o != nil && o.bitmap_&16 != 0
	if ok {
		value = o.scaleDown
	}
	return
}

// SkipNodesWithLocalStorage returns the value of the 'skip_nodes_with_local_storage' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Enables/Disables `--skip-nodes-with-local-storage` CA feature flag. If true cluster autoscaler will never delete nodes with pods with local storage, e.g. EmptyDir or HostPath. true by default at autoscaler.
func (o *ClusterAutoscaler) SkipNodesWithLocalStorage() bool {
	if o != nil && o.bitmap_&32 != 0 {
		return o.skipNodesWithLocalStorage
	}
	return false
}

// GetSkipNodesWithLocalStorage returns the value of the 'skip_nodes_with_local_storage' attribute and
// a flag indicating if the attribute has a value.
//
// Enables/Disables `--skip-nodes-with-local-storage` CA feature flag. If true cluster autoscaler will never delete nodes with pods with local storage, e.g. EmptyDir or HostPath. true by default at autoscaler.
func (o *ClusterAutoscaler) GetSkipNodesWithLocalStorage() (value bool, ok bool) {
	ok = o != nil && o.bitmap_&32 != 0
	if ok {
		value = o.skipNodesWithLocalStorage
	}
	return
}

// ClusterAutoscalerListKind is the name of the type used to represent list of objects of
// type 'cluster_autoscaler'.
const ClusterAutoscalerListKind = "ClusterAutoscalerList"

// ClusterAutoscalerListLinkKind is the name of the type used to represent links to list
// of objects of type 'cluster_autoscaler'.
const ClusterAutoscalerListLinkKind = "ClusterAutoscalerListLink"

// ClusterAutoscalerNilKind is the name of the type used to nil lists of objects of
// type 'cluster_autoscaler'.
const ClusterAutoscalerListNilKind = "ClusterAutoscalerListNil"

// ClusterAutoscalerList is a list of values of the 'cluster_autoscaler' type.
type ClusterAutoscalerList struct {
	href  string
	link  bool
	items []*ClusterAutoscaler
}

// Len returns the length of the list.
func (l *ClusterAutoscalerList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Empty returns true if the list is empty.
func (l *ClusterAutoscalerList) Empty() bool {
	return l == nil || len(l.items) == 0
}

// Get returns the item of the list with the given index. If there is no item with
// that index it returns nil.
func (l *ClusterAutoscalerList) Get(i int) *ClusterAutoscaler {
	if l == nil || i < 0 || i >= len(l.items) {
		return nil
	}
	return l.items[i]
}

// Slice returns an slice containing the items of the list. The returned slice is a
// copy of the one used internally, so it can be modified without affecting the
// internal representation.
//
// If you don't need to modify the returned slice consider using the Each or Range
// functions, as they don't need to allocate a new slice.
func (l *ClusterAutoscalerList) Slice() []*ClusterAutoscaler {
	var slice []*ClusterAutoscaler
	if l == nil {
		slice = make([]*ClusterAutoscaler, 0)
	} else {
		slice = make([]*ClusterAutoscaler, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *ClusterAutoscalerList) Each(f func(item *ClusterAutoscaler) bool) {
	if l == nil {
		return
	}
	for _, item := range l.items {
		if !f(item) {
			break
		}
	}
}

// Range runs the given function for each index and item of the list, in order. If
// the function returns false the iteration stops, otherwise it continues till all
// the elements of the list have been processed.
func (l *ClusterAutoscalerList) Range(f func(index int, item *ClusterAutoscaler) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
