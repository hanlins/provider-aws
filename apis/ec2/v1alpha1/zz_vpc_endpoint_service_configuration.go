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

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// VPCEndpointServiceConfigurationParameters defines the desired state of VPCEndpointServiceConfiguration
type VPCEndpointServiceConfigurationParameters struct {
	// Region is which region the VPCEndpointServiceConfiguration will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// Indicates whether requests from service consumers to create an endpoint to
	// your service must be accepted. To accept a request, use AcceptVpcEndpointConnections.
	AcceptanceRequired *bool `json:"acceptanceRequired,omitempty"`
	// The Amazon Resource Names (ARNs) of one or more Gateway Load Balancers.
	GatewayLoadBalancerARNs []*string `json:"gatewayLoadBalancerARNs,omitempty"`
	// The Amazon Resource Names (ARNs) of one or more Network Load Balancers for
	// your service.
	NetworkLoadBalancerARNs []*string `json:"networkLoadBalancerARNs,omitempty"`
	// (Interface endpoint configuration) The private DNS name to assign to the
	// VPC endpoint service.
	PrivateDNSName *string `json:"privateDNSName,omitempty"`
	// The tags to associate with the service.
	TagSpecifications                               []*TagSpecification `json:"tagSpecifications,omitempty"`
	CustomVPCEndpointServiceConfigurationParameters `json:",inline"`
}

// VPCEndpointServiceConfigurationSpec defines the desired state of VPCEndpointServiceConfiguration
type VPCEndpointServiceConfigurationSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       VPCEndpointServiceConfigurationParameters `json:"forProvider"`
}

// VPCEndpointServiceConfigurationObservation defines the observed state of VPCEndpointServiceConfiguration
type VPCEndpointServiceConfigurationObservation struct {
	// Unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientToken *string `json:"clientToken,omitempty"`
	// Information about the service configuration.
	ServiceConfiguration *ServiceConfiguration `json:"serviceConfiguration,omitempty"`
}

// VPCEndpointServiceConfigurationStatus defines the observed state of VPCEndpointServiceConfiguration.
type VPCEndpointServiceConfigurationStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          VPCEndpointServiceConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VPCEndpointServiceConfiguration is the Schema for the VPCEndpointServiceConfigurations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type VPCEndpointServiceConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VPCEndpointServiceConfigurationSpec   `json:"spec"`
	Status            VPCEndpointServiceConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPCEndpointServiceConfigurationList contains a list of VPCEndpointServiceConfigurations
type VPCEndpointServiceConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPCEndpointServiceConfiguration `json:"items"`
}

// Repository type metadata.
var (
	VPCEndpointServiceConfigurationKind             = "VPCEndpointServiceConfiguration"
	VPCEndpointServiceConfigurationGroupKind        = schema.GroupKind{Group: Group, Kind: VPCEndpointServiceConfigurationKind}.String()
	VPCEndpointServiceConfigurationKindAPIVersion   = VPCEndpointServiceConfigurationKind + "." + GroupVersion.String()
	VPCEndpointServiceConfigurationGroupVersionKind = GroupVersion.WithKind(VPCEndpointServiceConfigurationKind)
)

func init() {
	SchemeBuilder.Register(&VPCEndpointServiceConfiguration{}, &VPCEndpointServiceConfigurationList{})
}
