package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:shortName=rsp
type ResourcePolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ResourcePolicySpec   `json:"spec"`
	Status ResourcePolicyStatus `json:"status,omitempty"`
}

type ResourcePolicySpec struct {
	// +optional
	// +nullable
	// +listType=atomic
	Units []Unit `json:"units,omitempty" protobuf:"bytes,1,rep,name=units"`

	Selector       map[string]string `json:"selector,omitempty" protobuf:"bytes,2,rep,name=selector"`
	MatchLabelKeys []string          `json:"matchLabelKeys,omitempty" protobuf:"bytes,3,rep,name=matchLabelKeys"`
}

type Unit struct {
	Max          *int32          `json:"max,omitempty" protobuf:"varint,1,opt,name=max"`
	MaxResources v1.ResourceList `json:"maxResources,omitempty" protobuf:"bytes,2,rep,name=maxResources"`

	NodeSelector map[string]string `json:"nodeSelector,omitempty" protobuf:"bytes,3,rep,name=nodeSelector"`

	PodLabelsToAdd      map[string]string `json:"podLabels,omitempty" protobuf:"bytes,4,rep,name=podLabels"`
	PodAnnotationsToAdd map[string]string `json:"podAnnotations,omitempty" protobuf:"bytes,5,rep,name=podAnnotations"`
}

type ResourcePolicyStatus struct {
	Pods           []int64      `json:"pods,omitempty"`
	LastUpdateTime *metav1.Time `json:"lastUpdateTime,omitempty"`
}

// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ResourcePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []ResourcePolicy `json:"items"`
}
