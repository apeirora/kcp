//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The KCP Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	resource "k8s.io/apimachinery/pkg/api/resource"
	runtime "k8s.io/apimachinery/pkg/runtime"

	conditionsv1alpha1 "github.com/kcp-dev/kcp/third_party/conditions/apis/conditions/v1alpha1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadCluster) DeepCopyInto(out *WorkloadCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadCluster.
func (in *WorkloadCluster) DeepCopy() *WorkloadCluster {
	if in == nil {
		return nil
	}
	out := new(WorkloadCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkloadCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadClusterCondition) DeepCopyInto(out *WorkloadClusterCondition) {
	*out = *in
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = new(conditionsv1alpha1.Condition)
		(*in).DeepCopyInto(*out)
	}
	in.LastHeartbeatTime.DeepCopyInto(&out.LastHeartbeatTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadClusterCondition.
func (in *WorkloadClusterCondition) DeepCopy() *WorkloadClusterCondition {
	if in == nil {
		return nil
	}
	out := new(WorkloadClusterCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in WorkloadClusterConditions) DeepCopyInto(out *WorkloadClusterConditions) {
	{
		in := &in
		*out = make(WorkloadClusterConditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadClusterConditions.
func (in WorkloadClusterConditions) DeepCopy() WorkloadClusterConditions {
	if in == nil {
		return nil
	}
	out := new(WorkloadClusterConditions)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadClusterList) DeepCopyInto(out *WorkloadClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WorkloadCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadClusterList.
func (in *WorkloadClusterList) DeepCopy() *WorkloadClusterList {
	if in == nil {
		return nil
	}
	out := new(WorkloadClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkloadClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadClusterSpec) DeepCopyInto(out *WorkloadClusterSpec) {
	*out = *in
	if in.EvictAfter != nil {
		in, out := &in.EvictAfter, &out.EvictAfter
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadClusterSpec.
func (in *WorkloadClusterSpec) DeepCopy() *WorkloadClusterSpec {
	if in == nil {
		return nil
	}
	out := new(WorkloadClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadClusterStatus) DeepCopyInto(out *WorkloadClusterStatus) {
	*out = *in
	if in.Allocatable != nil {
		in, out := &in.Allocatable, &out.Allocatable
		*out = new(v1.ResourceList)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[v1.ResourceName]resource.Quantity, len(*in))
			for key, val := range *in {
				(*out)[key] = val.DeepCopy()
			}
		}
	}
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		*out = new(v1.ResourceList)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[v1.ResourceName]resource.Quantity, len(*in))
			for key, val := range *in {
				(*out)[key] = val.DeepCopy()
			}
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(WorkloadClusterConditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SyncedResources != nil {
		in, out := &in.SyncedResources, &out.SyncedResources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.LastHeartbeat != nil {
		in, out := &in.LastHeartbeat, &out.LastHeartbeat
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadClusterStatus.
func (in *WorkloadClusterStatus) DeepCopy() *WorkloadClusterStatus {
	if in == nil {
		return nil
	}
	out := new(WorkloadClusterStatus)
	in.DeepCopyInto(out)
	return out
}
