package azure

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	capzv1 "sigs.k8s.io/cluster-api-provider-azure/api/v1alpha3"
	capiv1 "sigs.k8s.io/cluster-api/api/v1alpha3"

	"github.com/q8s-io/mcp/pkg/k8s"
	"github.com/q8s-io/mcp/pkg/k8s/cluster"
)

type Clusters struct {
	objects []runtime.Object
}

func NewClusterComponents() cluster.Components {
	clusters := &Clusters{
		objects: make([]runtime.Object, 2),
	}
	clusters.Setup()
	return clusters
}

func (c *Clusters) Create() {
	for _, component := range c.objects {
		k8s.GetManager().GetClient().Create(context.TODO(), component)
	}
}

func (c *Clusters) Delete() {
	for _, component := range c.objects {
		k8s.GetManager().GetClient().Delete(context.TODO(), component)
	}
}

func (c *Clusters) Setup() {
	c.objects[0] = &capiv1.Cluster{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-ly",
			Namespace: "default",
		},
		Spec: capiv1.ClusterSpec{
			ClusterNetwork: &capiv1.ClusterNetwork{
				Pods: &capiv1.NetworkRanges{
					CIDRBlocks: []string{"192.168.0.0/16"},
				},
			},
			ControlPlaneRef: &corev1.ObjectReference{
				APIVersion: "controlplane.cluster.x-k8s.io/v1alpha3",
				Kind:       "KubeadmControlPlane",
				Name:       "test-ly-control-plane",
			},
			InfrastructureRef: &corev1.ObjectReference{
				APIVersion: "infrastructure.cluster.x-k8s.io/v1alpha3",
				Kind:       "AzureCluster",
				Name:       "test-ly",
			},
		},
	}

	c.objects[1] = &capzv1.AzureCluster{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-ly",
			Namespace: "default",
		},
		Spec: capzv1.AzureClusterSpec{
			Location: "chinanorth",
			NetworkSpec: capzv1.NetworkSpec{
				Vnet: capzv1.VnetSpec{
					Name: "test-ly-vnet",
				},
			},
			ResourceGroup: "test-ly",
		},
	}
}
