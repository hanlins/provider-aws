//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2020 The Crossplane Authors.

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

// Code generated by controller-gen. DO NOT EDIT.

package manualv1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCCIDRBlock) DeepCopyInto(out *VPCCIDRBlock) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCCIDRBlock.
func (in *VPCCIDRBlock) DeepCopy() *VPCCIDRBlock {
	if in == nil {
		return nil
	}
	out := new(VPCCIDRBlock)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VPCCIDRBlock) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCCIDRBlockList) DeepCopyInto(out *VPCCIDRBlockList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VPCCIDRBlock, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCCIDRBlockList.
func (in *VPCCIDRBlockList) DeepCopy() *VPCCIDRBlockList {
	if in == nil {
		return nil
	}
	out := new(VPCCIDRBlockList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VPCCIDRBlockList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCCIDRBlockObservation) DeepCopyInto(out *VPCCIDRBlockObservation) {
	*out = *in
	if in.AssociationID != nil {
		in, out := &in.AssociationID, &out.AssociationID
		*out = new(string)
		**out = **in
	}
	if in.CIDRBlock != nil {
		in, out := &in.CIDRBlock, &out.CIDRBlock
		*out = new(string)
		**out = **in
	}
	if in.IPv6CIDRBlock != nil {
		in, out := &in.IPv6CIDRBlock, &out.IPv6CIDRBlock
		*out = new(string)
		**out = **in
	}
	if in.IPv6CIDRBlockState != nil {
		in, out := &in.IPv6CIDRBlockState, &out.IPv6CIDRBlockState
		*out = new(VPCCIDRBlockState)
		(*in).DeepCopyInto(*out)
	}
	if in.IPv6Pool != nil {
		in, out := &in.IPv6Pool, &out.IPv6Pool
		*out = new(string)
		**out = **in
	}
	if in.NetworkBorderGroup != nil {
		in, out := &in.NetworkBorderGroup, &out.NetworkBorderGroup
		*out = new(string)
		**out = **in
	}
	if in.CIDRBlockState != nil {
		in, out := &in.CIDRBlockState, &out.CIDRBlockState
		*out = new(VPCCIDRBlockState)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCCIDRBlockObservation.
func (in *VPCCIDRBlockObservation) DeepCopy() *VPCCIDRBlockObservation {
	if in == nil {
		return nil
	}
	out := new(VPCCIDRBlockObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCCIDRBlockParameters) DeepCopyInto(out *VPCCIDRBlockParameters) {
	*out = *in
	if in.AmazonProvidedIPv6CIDRBlock != nil {
		in, out := &in.AmazonProvidedIPv6CIDRBlock, &out.AmazonProvidedIPv6CIDRBlock
		*out = new(bool)
		**out = **in
	}
	if in.CIDRBlock != nil {
		in, out := &in.CIDRBlock, &out.CIDRBlock
		*out = new(string)
		**out = **in
	}
	if in.IPv6CIDRBlock != nil {
		in, out := &in.IPv6CIDRBlock, &out.IPv6CIDRBlock
		*out = new(string)
		**out = **in
	}
	if in.IPv6CIDRBlockNetworkBorderGroup != nil {
		in, out := &in.IPv6CIDRBlockNetworkBorderGroup, &out.IPv6CIDRBlockNetworkBorderGroup
		*out = new(string)
		**out = **in
	}
	if in.IPv6Pool != nil {
		in, out := &in.IPv6Pool, &out.IPv6Pool
		*out = new(string)
		**out = **in
	}
	if in.VPCID != nil {
		in, out := &in.VPCID, &out.VPCID
		*out = new(string)
		**out = **in
	}
	if in.VPCIDRef != nil {
		in, out := &in.VPCIDRef, &out.VPCIDRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.VPCIDSelector != nil {
		in, out := &in.VPCIDSelector, &out.VPCIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCCIDRBlockParameters.
func (in *VPCCIDRBlockParameters) DeepCopy() *VPCCIDRBlockParameters {
	if in == nil {
		return nil
	}
	out := new(VPCCIDRBlockParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCCIDRBlockSpec) DeepCopyInto(out *VPCCIDRBlockSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCCIDRBlockSpec.
func (in *VPCCIDRBlockSpec) DeepCopy() *VPCCIDRBlockSpec {
	if in == nil {
		return nil
	}
	out := new(VPCCIDRBlockSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCCIDRBlockState) DeepCopyInto(out *VPCCIDRBlockState) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.StatusMessage != nil {
		in, out := &in.StatusMessage, &out.StatusMessage
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCCIDRBlockState.
func (in *VPCCIDRBlockState) DeepCopy() *VPCCIDRBlockState {
	if in == nil {
		return nil
	}
	out := new(VPCCIDRBlockState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCCIDRBlockStatus) DeepCopyInto(out *VPCCIDRBlockStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCCIDRBlockStatus.
func (in *VPCCIDRBlockStatus) DeepCopy() *VPCCIDRBlockStatus {
	if in == nil {
		return nil
	}
	out := new(VPCCIDRBlockStatus)
	in.DeepCopyInto(out)
	return out
}
