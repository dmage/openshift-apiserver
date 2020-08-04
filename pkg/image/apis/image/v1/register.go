package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
	corev1conversions "k8s.io/kubernetes/pkg/apis/core/v1"

	v1 "github.com/openshift/api/image/v1"
	"github.com/openshift/openshift-apiserver/pkg/image/apis/image"
	"github.com/openshift/openshift-apiserver/pkg/image/apis/image/docker10"
	"github.com/openshift/openshift-apiserver/pkg/image/apis/image/dockerpre012"
)

var (
	localSchemeBuilder = runtime.NewSchemeBuilder(
		image.Install,
		v1.Install,
		corev1conversions.AddToScheme,
		docker10.Install,
		dockerpre012.Install,

		addFieldSelectorKeyConversions,
		AddConversionFuncs,
		RegisterDefaults,
	)
	Install = localSchemeBuilder.AddToScheme
)
