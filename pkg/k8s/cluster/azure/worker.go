package azure

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	clusterazurev1 "sigs.k8s.io/cluster-api-provider-azure/api/v1alpha3"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1alpha3"
	bootstrapv1 "sigs.k8s.io/cluster-api/bootstrap/kubeadm/api/v1alpha3"
	bootstraptypesv1 "sigs.k8s.io/cluster-api/bootstrap/kubeadm/types/v1beta1"

	"github.com/q8s-io/mcp/pkg/k8s"
	"github.com/q8s-io/mcp/pkg/k8s/cluster"
)

type Workers struct {
	objects []runtime.Object
}

func NewWorkerComponents() cluster.Components {
	workers := &Workers{
		objects: make([]runtime.Object, 3),
	}
	workers.Setup()
	return workers
}

func (w *Workers) Create() {
	for _, component := range w.objects {
		k8s.GetManager().GetClient().Create(context.TODO(), component)
	}
}

func (w *Workers) Delete() {
	for _, component := range w.objects {
		k8s.GetManager().GetClient().Delete(context.TODO(), component)
	}
}

func (w *Workers) Setup() {
	namespace := "tenant-gzw"

	// MachineDeployment
	w.objects[0] = &clusterv1.MachineDeployment{
		ObjectMeta: v1.ObjectMeta{
			Namespace: namespace,
			Name:      "test-gzw-md-0",
		},
		Spec: clusterv1.MachineDeploymentSpec{
			ClusterName: "test-gzw",
			Replicas: func() *int32 {
				replica := int32(1)
				return &replica
			}(),
			Selector: v1.LabelSelector{
				MatchLabels: map[string]string{
					"matchLabels": "null",
				},
			},
			Template: clusterv1.MachineTemplateSpec{
				Spec: clusterv1.MachineSpec{
					ClusterName: "test-gzw",
					Version: func() *string {
						version := "v1.18.1"
						return &version
					}(),
					Bootstrap: clusterv1.Bootstrap{
						ConfigRef: &corev1.ObjectReference{
							APIVersion: "bootstrap.cluster.x-k8s.io/v1alpha3",
							Kind:       "KubeadmConfigTemplate",
							Namespace:  namespace,
							Name:       "test-gzw-md-0",
						},
					},
					InfrastructureRef: corev1.ObjectReference{
						APIVersion: "infrastructure.cluster.x-k8s.io/v1alpha3",
						Kind:       "AzureMachineTemplate",
						Namespace:  namespace,
						Name:       "test-gzw-md-0",
					},
				},
			},
		},
	}

	// AzureMachineTemplate
	w.objects[1] = &clusterazurev1.AzureMachineTemplate{
		ObjectMeta: v1.ObjectMeta{
			Namespace: namespace,
			Name:      "test-gzw-md-0",
		},
		Spec: clusterazurev1.AzureMachineTemplateSpec{
			Template: clusterazurev1.AzureMachineTemplateResource{
				Spec: clusterazurev1.AzureMachineSpec{
					Location: "chinanorth",
					OSDisk: clusterazurev1.OSDisk{
						OSType:     "Linux",
						DiskSizeGB: 30,
						ManagedDisk: clusterazurev1.ManagedDisk{
							StorageAccountType: "Standard_LRS",
						},
					},
					SSHPublicKey: "c3NoLXJzYSBBQUFBQjNOemFDMXljMkVBQUFBREFRQUJBQUFCQVFEVVFuRFFMNjJ4L0hRaGU2U1kwdEFkYTU1NEFhRE9DRXl2SDZPK09SMlRacE1IRVhtZElibkowdUZvVmJnLzBLZTZzVzFycHlmTVJrQkhaL2czZGthc0lxUmpvL1lGQzM4eDBqLzRnVUNaUDhIRlFDSW9ocmJkYldRQnZ5akVEajB0MW9uWEFxd0w5T2UxeGNJTHdBNjBySnFKV1YxbGsvbFRGYVArM0VaOEl0MmRpd2VCUmIraWVxTXFOMVczeTVBaHM5dFFNOERpSnZuTnBmTERaTmRwd3ZpbUJaNlBWdG0zcEhSVk0vdEhuTGtmMTcreHU5ZW8zdGt0aFRFcjhrNGdZb2lFcnRtemxqck9LZ3VQejZJbW5aSWUvU1Q0THMreU92RkZqOVlpekUvMDltY2xPa2FCcTllN3N6eWNaTlBsUGV2ZDYwTmpTYnF5S1JBNjhMeXYgZ2FvQGdhdWx6aHcK",
					VMSize:       "Standard_A2",
				},
			},
		},
	}

	// KubeadmConfigTemplate
	w.objects[2] = &bootstrapv1.KubeadmConfigTemplate{
		ObjectMeta: v1.ObjectMeta{
			Namespace: namespace,
			Name:      "test-gzw-md-0",
		},
		Spec: bootstrapv1.KubeadmConfigTemplateSpec{
			Template: bootstrapv1.KubeadmConfigTemplateResource{
				Spec: bootstrapv1.KubeadmConfigSpec{
					JoinConfiguration: &bootstraptypesv1.JoinConfiguration{
						NodeRegistration: bootstraptypesv1.NodeRegistrationOptions{
							Name: `{{ ds.meta_data["local_hostname"] }}`,
							KubeletExtraArgs: map[string]string{
								"cloud-config":   "/etc/kubernetes/azure.json",
								"cloud-provider": "azure",
							},
						},
					},
					Files: []bootstrapv1.File{
						{
							Owner:       "root:root",
							Path:        "/etc/kubernetes/azure.json",
							Permissions: "0644",
							Content: `{
	"cloud": "AzureChinaCloud",
	"tenantId": "c9572b54-e243-4caf-8684-cff70654c290",
	"subscriptionId": "593ea0a5-2089-4f6f-be30-ebe12fc78339",
	"aadClientId": "8736e6cc-dfd4-415f-98d4-c92f1de063f5",
	"aadClientSecret": "Jka@-iT8-H-yjyWvlzCOcbpo0huX36ns",
	"resourceGroup": "test-gzw",
	"securityGroupName": "test-gzw-node-nsg",
	"location": "chinanorth",
	"vmType": "standard",
	"vnetName": "test-gzw-vnet",
	"vnetResourceGroup": "test-gzw",
	"subnetName": "test-gzw-node-subnet",
	"routeTableName": "test-gzw-node-routetable",
	"loadBalancerSku": "standard",
	"maximumLoadBalancerRuleCount": 250,
	"useManagedIdentityExtension": false,
	"useInstanceMetadata": true
}`,
						},
					},
				},
			},
		},
	}
}
