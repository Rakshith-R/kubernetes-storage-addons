/*
Copyright 2021.

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

type SpaceReclaimSpec struct {
	PVC string `json:"pvc,omitempty"`
	PV  string `json:"pv,omitempty"`
}

type OperationSpec struct {
	SpaceReclaim SpaceReclaimSpec `json:"spaceReclaim,omitempty"`
}

// CSIAddonsJobSpec defines the desired state of CSIAddonsJob
type CSIAddonsJobSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of CSIAddonsJob. Edit csiaddonsjob_types.go to remove/update
	Operation OperationSpec `json:"operation"`
}

// CSIAddonsJobStatus defines the observed state of CSIAddonsJob
type CSIAddonsJobStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// CSIAddonsJob is the Schema for the csiaddonsjobs API
type CSIAddonsJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CSIAddonsJobSpec   `json:"spec,omitempty"`
	Status CSIAddonsJobStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// CSIAddonsJobList contains a list of CSIAddonsJob
type CSIAddonsJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CSIAddonsJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CSIAddonsJob{}, &CSIAddonsJobList{})
}
