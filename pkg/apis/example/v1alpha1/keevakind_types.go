package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// KeevakindSpec defines the desired state of Keevakind
// +k8s:openapi-gen=true
type KeevakindSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Count int32  `json:"count"`
	Group string `json:"group"`
	Image string `json:"image"`
	Port  int32  `json:"port"`
}

// KeevakindStatus defines the observed state of Keevakind
// +k8s:openapi-gen=true
type KeevakindStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	PodNames []string `json:"podnames"`
	AppGroup string   `json:"appgroup"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Keevakind is the Schema for the keevakinds API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type Keevakind struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KeevakindSpec   `json:"spec,omitempty"`
	Status KeevakindStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KeevakindList contains a list of Keevakind
type KeevakindList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Keevakind `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Keevakind{}, &KeevakindList{})
}
