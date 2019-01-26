// NOTE: Boilerplate only.  Ignore this file.

// Package v1beta1 contains API Schema definitions for the test v1beta1 API group
// +k8s:deepcopy-gen=package,register
// +groupName=test.opsI
package v1beta1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/runtime/scheme"
	"k8s.io/apimachinery/pkg/runtime"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	// SchemeGroupVersion is group version used to register these objects
	SchemeGroupVersion = schema.GroupVersion{Group: "test.opsI", Version: "v1beta1"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}



	schemeBuilder = runtime.NewSchemeBuilder(addKnownTypes, corev1.AddToScheme)



	Install = schemeBuilder.AddToScheme
	AddToScheme = schemeBuilder.AddToScheme

)

// Adds the list of known types to api.Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&OpsIAction{},
		&OpsIActionList{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
