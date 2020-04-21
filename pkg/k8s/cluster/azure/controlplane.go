package azure

import (
	"context"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	capzv1 "sigs.k8s.io/cluster-api-provider-azure/api/v1alpha3"
	cabpkv1 "sigs.k8s.io/cluster-api/bootstrap/kubeadm/api/v1alpha3"
	bootstrapv1 "sigs.k8s.io/cluster-api/bootstrap/kubeadm/types/v1beta1"
	controlplanev1 "sigs.k8s.io/cluster-api/controlplane/kubeadm/api/v1alpha3"

	"github.com/q8s-io/mcp/pkg/k8s"
	"github.com/q8s-io/mcp/pkg/k8s/cluster"
)

type ControlPlanes struct {
	objects []runtime.Object
}

func NewControlPlaneComponents() cluster.Components {
	controlplanes := &ControlPlanes{
		objects: make([]runtime.Object, 2),
	}
	controlplanes.Setup()
	return controlplanes
}

func (c *ControlPlanes) Create() {
	for _, component := range c.objects {
		k8s.GetManager().GetClient().Create(context.TODO(), component)
	}
}

func (c *ControlPlanes) Delete() {
	for _, component := range c.objects {
		k8s.GetManager().GetClient().Delete(context.TODO(), component)
	}
}

func (c *ControlPlanes) Setup() {
	c.objects[0] = &controlplanev1.KubeadmControlPlane{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-ly-control-plane",
			Namespace: "default",
		},
		Spec: controlplanev1.KubeadmControlPlaneSpec{
			InfrastructureTemplate: corev1.ObjectReference{
				APIVersion: "infrastructure.cluster.x-k8s.io/v1alpha3",
				Kind:       "AzureMachineTemplate",
				Name:       "test-ly-control-plane",
			},
			KubeadmConfigSpec: cabpkv1.KubeadmConfigSpec{
				ClusterConfiguration: &bootstrapv1.ClusterConfiguration{
					APIServer: bootstrapv1.APIServer{
						ControlPlaneComponent: bootstrapv1.ControlPlaneComponent{
							ExtraArgs: map[string]string{
								"cloud-config":   "/etc/kubernetes/azure.json",
								"cloud-provider": "azure",
							},
							ExtraVolumes: []bootstrapv1.HostPathMount{
								{
									HostPath:  "/etc/kubernetes/azure.json",
									MountPath: "/etc/kubernetes/azure.json",
									Name:      "cloud-config",
									ReadOnly:  true,
								},
							},
						},
						TimeoutForControlPlane: &metav1.Duration{
							Duration: 20 * time.Minute,
						},
					},
					ControllerManager: bootstrapv1.ControlPlaneComponent{
						ExtraArgs: map[string]string{
							"allocate-node-cidrs": "false",
							"cloud-config":        "/etc/kubernetes/azure.json",
							"cloud-provider":      "azure",
						},
						ExtraVolumes: []bootstrapv1.HostPathMount{
							{
								HostPath:  "/etc/kubernetes/azure.json",
								MountPath: "/etc/kubernetes/azure.json",
								Name:      "cloud-config",
								ReadOnly:  true,
							},
						},
					},
				},
				Files: []cabpkv1.File{
					{
						Content: `{
	"cloud": "AzureChinaCloud",
	"tenantId": "c9572b54-e243-4caf-8684-cff70654c290",
	"subscriptionId": "593ea0a5-2089-4f6f-be30-ebe12fc78339",
	"aadClientId": "8736e6cc-dfd4-415f-98d4-c92f1de063f5",
	"aadClientSecret": "Jka@-iT8-H-yjyWvlzCOcbpo0huX36ns",
	"resourceGroup": "test-ly",
	"securityGroupName": "test-ly-node-nsg",
	"location": "chinanorth",
	"vmType": "standard",
	"vnetName": "test-ly-vnet",
	"vnetResourceGroup": "test-ly",
	"subnetName": "test-ly-node-subnet",
	"routeTableName": "test-ly-node-routetable",
	"userAssignedID": "test-ly",
	"loadBalancerSku": "standard",
	"maximumLoadBalancerRuleCount": 250,
	"useManagedIdentityExtension": false,
	"useInstanceMetadata": true
}`,
						Owner:       "root:root",
						Path:        "/etc/kubernetes/azure.json",
						Permissions: "0644",
					},
				},
				InitConfiguration: &bootstrapv1.InitConfiguration{
					NodeRegistration: bootstrapv1.NodeRegistrationOptions{
						KubeletExtraArgs: map[string]string{
							"cloud-config":   "/etc/kubernetes/azure.json",
							"cloud-provider": "azure",
						},
						Name: `{{ ds.meta_data["local_hostname"] }}`,
					},
				},
				JoinConfiguration: &bootstrapv1.JoinConfiguration{
					NodeRegistration: bootstrapv1.NodeRegistrationOptions{
						KubeletExtraArgs: map[string]string{
							"cloud-config":   "/etc/kubernetes/azure.json",
							"cloud-provider": "azure",
						},
						Name: `{{ ds.meta_data["local_hostname"] }}`,
					},
				},
			},
			Replicas: func() *int32 {
				replicas := int32(1)
				return &replicas
			}(),
			Version: "v1.16.7",
		},
	}

	AzureMachineTemplate := capzv1.AzureMachineTemplate{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-ly-control-plane-template",
			Namespace: "default",
		},
		Spec: capzv1.AzureMachineTemplateSpec{
			Template: capzv1.AzureMachineTemplateResource{
				Spec: capzv1.AzureMachineSpec{
					Location: "chinanorth",
					OSDisk: capzv1.OSDisk{
						DiskSizeGB: 30,
						ManagedDisk: capzv1.ManagedDisk{
							StorageAccountType: "Standard_LRS",
						},
						OSType: "Linux",
					},
					SSHPublicKey: "c3NoLXJzYSBBQUFBQjNOemFDMXljMkVBQUFBREFRQUJBQUFCQVFEVVFuRFFMNjJ4L0hRaGU2U1kwdEFkYTU1NEFhRE9DRXl2SDZPK09SMlRacE1IRVhtZElibkowdUZvVmJnLzBLZTZzVzFycHlmTVJrQkhaL2czZGthc0lxUmpvL1lGQzM4eDBqLzRnVUNaUDhIRlFDSW9ocmJkYldRQnZ5akVEajB0MW9uWEFxd0w5T2UxeGNJTHdBNjBySnFKV1YxbGsvbFRGYVArM0VaOEl0MmRpd2VCUmIraWVxTXFOMVczeTVBaHM5dFFNOERpSnZuTnBmTERaTmRwd3ZpbUJaNlBWdG0zcEhSVk0vdEhuTGtmMTcreHU5ZW8zdGt0aFRFcjhrNGdZb2lFcnRtemxqck9LZ3VQejZJbW5aSWUvU1Q0THMreU92RkZqOVlpekUvMDltY2xPa2FCcTllN3N6eWNaTlBsUGV2ZDYwTmpTYnF5S1JBNjhMeXYgZ2FvQGdhdWx6aHcK",
					VMSize:       "Standard_A2",
				},
			},
		},
	}
}
