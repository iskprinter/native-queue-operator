/*
Copyright 2023.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// NativeQueueSpec defines the desired state of NativeQueue
type NativeQueueSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of NativeQueue. Edit nativequeue_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// NativeQueueStatus defines the observed state of NativeQueue
type NativeQueueStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// NativeQueue is the Schema for the nativequeues API
type NativeQueue struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NativeQueueSpec   `json:"spec,omitempty"`
	Status NativeQueueStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// NativeQueueList contains a list of NativeQueue
type NativeQueueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NativeQueue `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NativeQueue{}, &NativeQueueList{})
}
