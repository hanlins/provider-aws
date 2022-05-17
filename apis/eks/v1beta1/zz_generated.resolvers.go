/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	v1beta1 "github.com/crossplane/provider-aws/apis/ec2/v1beta1"
	v1beta11 "github.com/crossplane/provider-aws/apis/iam/v1beta1"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Cluster.
func (mg *Cluster) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: mg.Spec.ForProvider.ResourcesVpcConfig.SecurityGroupIDs,
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.ResourcesVpcConfig.SecurityGroupIDRefs,
		Selector:      mg.Spec.ForProvider.ResourcesVpcConfig.SecurityGroupIDSelector,
		To: reference.To{
			List:    &v1beta1.SecurityGroupList{},
			Managed: &v1beta1.SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourcesVpcConfig.SecurityGroupIDs")
	}
	mg.Spec.ForProvider.ResourcesVpcConfig.SecurityGroupIDs = mrsp.ResolvedValues
	mg.Spec.ForProvider.ResourcesVpcConfig.SecurityGroupIDRefs = mrsp.ResolvedReferences

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: mg.Spec.ForProvider.ResourcesVpcConfig.SubnetIDs,
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.ResourcesVpcConfig.SubnetIDRefs,
		Selector:      mg.Spec.ForProvider.ResourcesVpcConfig.SubnetIDSelector,
		To: reference.To{
			List:    &v1beta1.SubnetList{},
			Managed: &v1beta1.Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourcesVpcConfig.SubnetIDs")
	}
	mg.Spec.ForProvider.ResourcesVpcConfig.SubnetIDs = mrsp.ResolvedValues
	mg.Spec.ForProvider.ResourcesVpcConfig.SubnetIDRefs = mrsp.ResolvedReferences

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: mg.Spec.ForProvider.RoleArn,
		Extract:      v1beta11.RoleARN(),
		Reference:    mg.Spec.ForProvider.RoleArnRef,
		Selector:     mg.Spec.ForProvider.RoleArnSelector,
		To: reference.To{
			List:    &v1beta11.RoleList{},
			Managed: &v1beta11.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RoleArn")
	}
	mg.Spec.ForProvider.RoleArn = rsp.ResolvedValue
	mg.Spec.ForProvider.RoleArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this FargateProfile.
func (mg *FargateProfile) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: mg.Spec.ForProvider.ClusterName,
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClusterNameRef,
		Selector:     mg.Spec.ForProvider.ClusterNameSelector,
		To: reference.To{
			List:    &ClusterList{},
			Managed: &Cluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterName")
	}
	mg.Spec.ForProvider.ClusterName = rsp.ResolvedValue
	mg.Spec.ForProvider.ClusterNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: mg.Spec.ForProvider.PodExecutionRoleArn,
		Extract:      v1beta11.RoleARN(),
		Reference:    mg.Spec.ForProvider.PodExecutionRoleArnRef,
		Selector:     mg.Spec.ForProvider.PodExecutionRoleArnSelector,
		To: reference.To{
			List:    &v1beta11.RoleList{},
			Managed: &v1beta11.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PodExecutionRoleArn")
	}
	mg.Spec.ForProvider.PodExecutionRoleArn = rsp.ResolvedValue
	mg.Spec.ForProvider.PodExecutionRoleArnRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: mg.Spec.ForProvider.Subnets,
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.SubnetRefs,
		Selector:      mg.Spec.ForProvider.SubnetSelector,
		To: reference.To{
			List:    &v1beta1.SubnetList{},
			Managed: &v1beta1.Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Subnets")
	}
	mg.Spec.ForProvider.Subnets = mrsp.ResolvedValues
	mg.Spec.ForProvider.SubnetRefs = mrsp.ResolvedReferences

	return nil
}
