package azure

import (
	"context"
	"errors"

	corev1 "k8s.io/api/core/v1"
	k8serros "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog"
	infraazurev1 "sigs.k8s.io/cluster-api-provider-azure/api/v1alpha3"
	"sigs.k8s.io/cluster-api/api/v1alpha3"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func CreateMaterStep1(client client.Client) error {
	cluster := v1alpha3.Cluster{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-ly",
			Namespace: "default",
		},
		Spec: v1alpha3.ClusterSpec{
			ClusterNetwork: &v1alpha3.ClusterNetwork{
				Pods: &v1alpha3.NetworkRanges{
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

	azurecluster := infraazurev1.AzureCluster{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-ly",
			Namespace: "default",
		},
		Spec: infraazurev1.AzureClusterSpec{
			Location: "chinanorth",
			NetworkSpec: infraazurev1.NetworkSpec{
				Vnet: infraazurev1.VnetSpec{
					Name: "test-ly-vnet",
				},
			},
			ResourceGroup: "test-ly",
		},
	}

	if err := CreateOrDelete(client, &cluster); err != nil {
		return err
	}

	if err := CreateOrDelete(client, &azurecluster); err != nil {
		return err
	}
	return nil
}

func CreateOrDelete(client client.Client, object runtime.Object) error {
	err := client.Create(context.TODO(), object)
	if err != nil {
		klog.Errorf("failed to create %s err:%s", object.GetObjectKind(), err)
		if k8serros.IsAlreadyExists(err) {
			klog.Info("%s already exists, attempt to delete %s", object.GetObjectKind(), object.GetObjectKind())
			if err = client.Delete(context.TODO(), object); err != nil {
				klog.Errorf("failed to delete %s err:%s", object.GetObjectKind(), err)
				return err
			} else {
				klog.Infof("failed to create %s but successfully deleted. need to create it again",
					object.GetObjectKind())
				return errors.New("_")
			}
		}
	}
	return err
}
