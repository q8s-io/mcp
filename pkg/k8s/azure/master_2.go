package azure

import (
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	infraazurev1 "sigs.k8s.io/cluster-api-provider-azure/api/v1alpha3"
	cabpkv1 "sigs.k8s.io/cluster-api/bootstrap/kubeadm/api/v1alpha3"
	kubeadmv1beta1 "sigs.k8s.io/cluster-api/bootstrap/kubeadm/types/v1beta1"
	controlplane "sigs.k8s.io/cluster-api/controlplane/kubeadm/api/v1alpha3"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func CreateMasterStep2(client client.Client) error {
	var replicas int32 = 1
	args1 := make(map[string]string)
	args1["cloud-config"] = "/etc/kubernetes/azure.json"
	args1["cloud-provider"] = "azure"

	args2 := make(map[string]string)
	args2["allocate-node-cidrs"] = "false"
	args2["cloud-config"] = "/etc/kubernetes/azure.json"
	args2["cloud-provider"] = "azure"

	KubeadmControlPlane := controlplane.KubeadmControlPlane{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-ly-control-plane",
			Namespace: "default",
		},
		Spec: controlplane.KubeadmControlPlaneSpec{
			InfrastructureTemplate: corev1.ObjectReference{
				APIVersion: "infrastructure.cluster.x-k8s.io/v1alpha3",
				Kind:       "AzureMachineTemplate",
				Name:       "test-ly-control-plane",
			},
			KubeadmConfigSpec: cabpkv1.KubeadmConfigSpec{
				ClusterConfiguration: &kubeadmv1beta1.ClusterConfiguration{
					APIServer: kubeadmv1beta1.APIServer{
						ControlPlaneComponent: kubeadmv1beta1.ControlPlaneComponent{
							ExtraArgs: args1,
							ExtraVolumes: []kubeadmv1beta1.HostPathMount{
								{
									HostPath:  "/etc/kubernetes/azure.json",
									MountPath: "/etc/kubernetes/azure.json",
									Name:      "cloud-config",
									ReadOnly:  true,
								},
							},
						},
						TimeoutForControlPlane: &metav1.Duration{
							Duration: time.Duration(20 * time.Minute)},
					},
					ControllerManager: kubeadmv1beta1.ControlPlaneComponent{
						ExtraArgs: args2,
						ExtraVolumes: []kubeadmv1beta1.HostPathMount{
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
				InitConfiguration: &kubeadmv1beta1.InitConfiguration{
					NodeRegistration: kubeadmv1beta1.NodeRegistrationOptions{
						KubeletExtraArgs: args1,
						Name:             `{{ ds.meta_data["local_hostname"] }}`,
					},
				},
				JoinConfiguration: &kubeadmv1beta1.JoinConfiguration{
					NodeRegistration: kubeadmv1beta1.NodeRegistrationOptions{
						KubeletExtraArgs: args1,
						Name:             `{{ ds.meta_data["local_hostname"] }}`,
					},
				},
			},
			Replicas: &replicas,
			Version:  "v1.16.7",
		},
	}
	if err := CreateOrDelete(client, &KubeadmControlPlane); err != nil {
		return err
	}

	AzureMachineTemplate := infraazurev1.AzureMachineTemplate{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-ly-control-plane-template",
			Namespace: "default",
		},
		Spec: infraazurev1.AzureMachineTemplateSpec{
			Template: infraazurev1.AzureMachineTemplateResource{
				Spec: infraazurev1.AzureMachineSpec{
					Location: "chinanorth",
					OSDisk: infraazurev1.OSDisk{
						DiskSizeGB: 30,
						ManagedDisk: infraazurev1.ManagedDisk{
							StorageAccountType: "Standard_LRS",
						},
						OSType: "Linux",
					},
					SSHPublicKey: " c3NoLXJzYSBBQUFBQjNOemFDMXljMkVBQUFBREFRQUJBQUFCQVFEVVFuRFFMNjJ4" +
						"L0hRaGU2U1kwdEFkYTU1NEFhRE9DRXl2SDZPK09SMlRacE1IRVhtZElibkowdUZvVmJnLzBLZTZz" +
						"VzFycHlmTV JrQkhaL2czZGthc0lxUmpvL1lGQzM4eDBqLzRnVUNaUDhIRlFDSW9ocmJkYldRQnZ" +
						"5akVEajB0MW9uWEFxd0w5T2UxeGNJTHdBNjBySnFKV1YxbGsvbFRGYVArM0VaOEl0MmRpd2VCUmI" +
						"raWVxTXFO MVczeTVBaHM5dFFNOERpSnZuTnBmTERaTmRwd3ZpbUJaNlBWdG0zcEhSVk0vdEhuTG" +
						"tmMTcreHU5ZW8zdGt0aFRFcjhrNGdZb2lFcnRtemxqck9LZ3VQejZJbW5aSWUvU1Q0THMreU92Rk" +
						"ZqOVlpek UvMDltY2xPa2FCcTllN3N6eWNaTlBsUGV2ZDYwTmpTYnF5S1JBNjhMeXYgZ2FvQGdhd" +
						"Wx6aHcK",
					VMSize: "Standard_A2",
				},
			},
		},
	}
	if err := CreateOrDelete(client, &AzureMachineTemplate); err != nil {
		return err
	}

	return nil
}
