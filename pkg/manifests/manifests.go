package manifests

import (
	"embed"
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/util/yaml"
)

//go:embed yamls
var yamls embed.FS

func GetDaemonSet() (*appsv1.DaemonSet, error) {
	f, err := yamls.ReadFile("daemonset.yaml")
	if err != nil {
		return nil, fmt.Errorf("failed to read from daemonset.yaml file; %w", err)
	}
	ds := &appsv1.DaemonSet{}
	if err = yaml.Unmarshal(f, ds); err != nil {
		return nil, fmt.Errorf("failed to unmarshal daemonset.yaml data; %w", err)
	}
	return ds, nil
}

func UpdateDaemonSet(ds *appsv1.DaemonSet, size int64) {
	rl := corev1.ResourceList{
		corev1.ResourceCPU: *resource.NewQuantity(size, resource.DecimalSI),
	}
	// shortcut
	podSpec := &ds.Spec.Template.Spec
	// TODO find by name
	podSpec.Containers[0].Resources = corev1.ResourceRequirements{Limits: rl}
}
