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

// APIParameters defines the desired state of API
type APIParameters struct {
	// Region is which region the API will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`

	APIKeySelectionExpression *string `json:"apiKeySelectionExpression,omitempty"`

	CORSConfiguration *CORS `json:"corsConfiguration,omitempty"`

	CredentialsARN *string `json:"credentialsARN,omitempty"`

	Description *string `json:"description,omitempty"`

	DisableExecuteAPIEndpoint *bool `json:"disableExecuteAPIEndpoint,omitempty"`

	DisableSchemaValidation *bool `json:"disableSchemaValidation,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name"`

	// +kubebuilder:validation:Required
	ProtocolType *string `json:"protocolType"`

	RouteKey *string `json:"routeKey,omitempty"`

	RouteSelectionExpression *string `json:"routeSelectionExpression,omitempty"`

	Tags map[string]*string `json:"tags,omitempty"`

	Target *string `json:"target,omitempty"`

	Version             *string `json:"version,omitempty"`
	CustomAPIParameters `json:",inline"`
}

// APISpec defines the desired state of API
type APISpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       APIParameters `json:"forProvider"`
}

// APIObservation defines the observed state of API
type APIObservation struct {
	APIEndpoint *string `json:"apiEndpoint,omitempty"`

	APIGatewayManaged *bool `json:"apiGatewayManaged,omitempty"`

	APIID *string `json:"apiID,omitempty"`

	CreatedDate *metav1.Time `json:"createdDate,omitempty"`

	ImportInfo []*string `json:"importInfo,omitempty"`

	Warnings []*string `json:"warnings,omitempty"`
}

// APIStatus defines the observed state of API.
type APIStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          APIObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// API is the Schema for the APIS API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type API struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              APISpec   `json:"spec"`
	Status            APIStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// APIList contains a list of APIS
type APIList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []API `json:"items"`
}

// Repository type metadata.
var (
	APIKind             = "API"
	APIGroupKind        = schema.GroupKind{Group: Group, Kind: APIKind}.String()
	APIKindAPIVersion   = APIKind + "." + GroupVersion.String()
	APIGroupVersionKind = GroupVersion.WithKind(APIKind)
)

func init() {
	SchemeBuilder.Register(&API{}, &APIList{})
}
