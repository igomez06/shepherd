package configmaps

import (
	coreV1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// ConfigMapGroupVersionResource is the required Group Version Resource for accessing config maps in a cluster,
// using the dynamic client.
var ConfigMapGroupVersionResource = schema.GroupVersionResource{
	Group:    "",
	Version:  "v1",
	Resource: "configmaps",
}

// NewConfigmapTemplate is a constructor that creates a configmap template
func NewConfigmapTemplate(configmapName, namespace string, annotations, labels, data map[string]string) coreV1.ConfigMap {
	if annotations == nil {
		annotations = make(map[string]string)
	}
	if labels == nil {
		labels = make(map[string]string)
	}
	if data == nil {
		data = make(map[string]string)
	}

	return coreV1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:        configmapName,
			Namespace:   namespace,
			Annotations: annotations,
			Labels:      labels,
		},
		Data: data,
	}
}
