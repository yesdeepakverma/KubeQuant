/*
Copyright 2026.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// KubeQuantRecommendationSpec defines the desired state of KubeQuantRecommendation
type KubeQuantRecommendationSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// The following markers will use OpenAPI v3 schema to validate the value
	// More info: https://book.kubebuilder.io/reference/markers/crd-validation.html

	// foo is an example field of KubeQuantRecommendation. Edit kubequantrecommendation_types.go to remove/update
	// +optional
	Namespace           *string `json:"namespace,omitempty"`  // namespace name where deployment/pod exists to monitor for
	Deployment          *string `json:"deployment,omitempty"` // name of the deployment to generate recommendation for. This and Pod values are mutual exclusive. The container name must be part of the deployment specification
	Pod                 *string `json:"pod,omitempty"`        // name of the pod to generate recommendation for. The container name must be part of this pod
	Container           *string `json:"container,omitempty"`  // name of the container to genrate the recommendation for
	TargetCpuPercentile *int8   `json:"targetCpuPercentile"`  // target cpu percentile p95, p80, to be in 90/80 without p
	TargetMemPercentile *int8   `json:"targetMemPercentile"`  // target cpu percentile p95, p80, to be in 90/80 without p
}

type Recommendation struct {
	RecommendedCpuRequests *int8 `json:"recommendedCpuRequests"` // recommended CPU Requests value
	RecommendedMemRequests *int8 `json:"recommendedMemRequests"` // recommended Mem Requests value
}

type RecommendationApproval struct {
	Approved       *bool        `json:"approved,omitempty"`
	ApprovedTime   *metav1.Time `json:"approvedTime,omitempty"`
	ApprovedReason *string      `json:"approvedReason,omitempty"`
	ApprovedBy     *string      `json:"approvedBy,omitempty"`
}

// KubeQuantRecommendationStatus defines the observed state of KubeQuantRecommendation.
type KubeQuantRecommendationStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// For Kubernetes API conventions, see:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#typical-status-properties

	// conditions represent the current state of the KubeQuantRecommendation resource.
	// Each condition has a unique type and reflects the status of a specific aspect of the resource.
	//
	// Standard condition types include:
	// - "Available": the resource is fully functional
	// - "Progressing": the resource is being created or updated
	// - "Degraded": the resource failed to reach or maintain its desired state
	//
	// The status of each condition is one of True, False, or Unknown.
	// +listType=map
	// +listMapKey=type
	// +optional
	Conditions     []metav1.Condition      `json:"conditions,omitempty"`
	Approval       *RecommendationApproval `json:"approval,omitempty"`
	Recommendation *Recommendation         `json:"recommendation,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// KubeQuantRecommendation is the Schema for the kubequantrecommendations API
type KubeQuantRecommendation struct {
	metav1.TypeMeta `json:",inline"`

	// metadata is a standard object metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitzero"`

	// spec defines the desired state of KubeQuantRecommendation
	// +required
	Spec KubeQuantRecommendationSpec `json:"spec"`

	// status defines the observed state of KubeQuantRecommendation
	// +optional
	Status KubeQuantRecommendationStatus `json:"status,omitzero"`
}

// +kubebuilder:object:root=true

// KubeQuantRecommendationList contains a list of KubeQuantRecommendation
type KubeQuantRecommendationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitzero"`
	Items           []KubeQuantRecommendation `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KubeQuantRecommendation{}, &KubeQuantRecommendationList{})
}
