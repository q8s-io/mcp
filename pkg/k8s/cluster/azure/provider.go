package azure

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"

	"github.com/q8s-io/mcp/pkg/k8s"
	"github.com/q8s-io/mcp/pkg/k8s/cluster"
)

type Providers struct {
	objects []runtime.Object
}

func NewProviderComponents() cluster.Components {
	providers := &Providers{
		objects: make([]runtime.Object, 8),
	}
	providers.Setup()
	return providers
}

func (p *Providers) Create() {
	for _, component := range p.objects {
		k8s.GetManager().GetClient().Create(context.TODO(), component)
	}
}

func (p *Providers) Delete() {
	for _, component := range p.objects {
		k8s.GetManager().GetClient().Delete(context.TODO(), component)
	}
}

func (p *Providers) Setup() {
	namespace := "tenant-gzw"
	// Namespace
	p.objects[0] = &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: namespace,
			Labels: map[string]string{
				"cluster.x-k8s.io/provider": "infrastructure-azure",
			},
		},
	}

	// ClusterRoleBinding
	p.objects[1] = &rbacv1.ClusterRoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name: namespace + "-manager-rolebinding",
			Labels: map[string]string{
				"cluster.x-k8s.io/provider": "infrastructure-azure",
			},
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "ClusterRole",
			Name:     "capz-manager-role",
		},
		Subjects: []rbacv1.Subject{
			{
				Kind:      "ServiceAccount",
				Name:      "default",
				Namespace: namespace,
			},
		},
	}

	// ClusterRoleBinding
	p.objects[2] = &rbacv1.ClusterRoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name: namespace + "-proxy-rolebinding",
			Labels: map[string]string{
				"cluster.x-k8s.io/provider": "infrastructure-azure",
			},
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "ClusterRole",
			Name:     "capz-proxy-role",
		},
		Subjects: []rbacv1.Subject{
			{
				Kind:      "ServiceAccount",
				Name:      "default",
				Namespace: namespace,
			},
		},
	}

	// Role
	p.objects[3] = &rbacv1.Role{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      "capz-leader-election-role",
			Labels: map[string]string{
				"cluster.x-k8s.io/provider": "infrastructure-azure",
			},
		},
		Rules: []rbacv1.PolicyRule{
			{
				APIGroups: []string{""},
				Resources: []string{"configmaps"},
				Verbs:     []string{"get", "list", "watch", "create", "update", "patch", "delete"},
			},
			{
				APIGroups: []string{""},
				Resources: []string{"configmaps/status"},
				Verbs:     []string{"get", "update", "patch"},
			},
		},
	}

	// RoleBinding
	p.objects[4] = &rbacv1.RoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      "capz-leader-election-rolebinding",
			Labels: map[string]string{
				"cluster.x-k8s.io/provider": "infrastructure-azure",
			},
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "Role",
			Name:     "capz-leader-election-role",
		},
		Subjects: []rbacv1.Subject{
			{
				Kind:      "ServiceAccount",
				Name:      "default",
				Namespace: namespace,
			},
		},
	}

	// Secret
	p.objects[5] = &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      "capz-manager-bootstrap-credentials",
			Labels: map[string]string{
				"cluster.x-k8s.io/provider": "infrastructure-azure",
			},
		},
		Data: map[string][]byte{
			"client-id":       []byte("ODczNmU2Y2MtZGZkNC00MTVmLTk4ZDQtYzkyZjFkZTA2M2Y1"),
			"client-secret":   []byte("SmthQC1pVDgtSC15anlXdmx6Q09jYnBvMGh1WDM2bnM="),
			"subscription-id": []byte("NTkzZWEwYTUtMjA4OS00ZjZmLWJlMzAtZWJlMTJmYzc4MzM5"),
			"tenant-id":       []byte("Yzk1NzJiNTQtZTI0My00Y2FmLTg2ODQtY2ZmNzA2NTRjMjkw"),
		},
	}

	// Service
	p.objects[6] = &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      "capz-controller-manager-metrics-service",
			Labels: map[string]string{
				"cluster.x-k8s.io/provider": "infrastructure-azure",
				"control-plane":             "capz-controller-manager",
			},
			Annotations: map[string]string{
				"prometheus.io/port":   "8443",
				"prometheus.io/scheme": "https",
				"prometheus.io/scrape": "true",
			},
		},
		Spec: v1.ServiceSpec{
			Ports: []v1.ServicePort{
				{
					Name:       "https",
					Port:       8443,
					TargetPort: intstr.FromString("https"),
				},
			},
			Selector: map[string]string{
				"cluster.x-k8s.io/provider": "infrastructure-azure",
				"control-plane":             "capz-controller-manager",
			},
		},
	}

	// Deployment
	p.objects[7] = &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      "capz-controller-manager",
			Labels: map[string]string{
				"cluster.x-k8s.io/provider": "infrastructure-azure",
				"control-plane":             "capz-controller-manager",
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: func() *int32 {
				replicas := int32(1)
				return &replicas
			}(),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"cluster.x-k8s.io/provider": "infrastructure-azure",
					"control-plane":             "capz-controller-manager",
				},
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"cluster.x-k8s.io/provider": "infrastructure-azure",
						"control-plane":             "capz-controller-manager",
					},
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:  "kube-rbac-proxy",
							Image: "gcr.io/kubebuilder/kube-rbac-proxy:v0.4.1",
							Args: []string{
								"--secure-listen-address=0.0.0.0:8443",
								"--upstream=http://127.0.0.1:8080/",
								"--logtostderr=true",
								"--v=10",
							},
							Ports: []v1.ContainerPort{
								{
									Name:          "https",
									ContainerPort: 8443,
								},
							},
						},
						{
							Name:            "manager",
							Image:           "docker.io/q8sio/cluster-api-azure-controller-cn:v0.4.1",
							ImagePullPolicy: v1.PullAlways,
							Args: []string{
								"--metrics-addr=127.0.0.1:8080",
								"--enable-leader-election",
							},
							Env: []v1.EnvVar{
								{
									Name:  "AZURE_ENVIRONMENT",
									Value: "AZURECHINACLOUD",
								},
								{
									Name: "AZURE_SUBSCRIPTION_ID",
									ValueFrom: &v1.EnvVarSource{
										SecretKeyRef: &v1.SecretKeySelector{
											Key: "subscription-id",
											LocalObjectReference: v1.LocalObjectReference{
												Name: "capz-manager-bootstrap-credentials",
											},
										},
									},
								},
								{
									Name: "AZURE_TENANT_ID",
									ValueFrom: &v1.EnvVarSource{
										SecretKeyRef: &v1.SecretKeySelector{
											Key: "tenant-id",
											LocalObjectReference: v1.LocalObjectReference{
												Name: "capz-manager-bootstrap-credentials",
											},
										},
									},
								},
								{
									Name: "AZURE_CLIENT_ID",
									ValueFrom: &v1.EnvVarSource{
										SecretKeyRef: &v1.SecretKeySelector{
											Key: "client-id",
											LocalObjectReference: v1.LocalObjectReference{
												Name: "capz-manager-bootstrap-credentials",
											},
										},
									},
								},
								{
									Name: "AZURE_CLIENT_SECRET",
									ValueFrom: &v1.EnvVarSource{
										SecretKeyRef: &v1.SecretKeySelector{
											Key: "client-secret",
											LocalObjectReference: v1.LocalObjectReference{
												Name: "capz-manager-bootstrap-credentials",
											},
										},
									},
								},
							},
							Ports: []v1.ContainerPort{
								{
									Name:          "healthz",
									ContainerPort: 9440,
									Protocol:      v1.ProtocolTCP,
								},
							},
							LivenessProbe: &v1.Probe{
								Handler: v1.Handler{
									HTTPGet: &v1.HTTPGetAction{
										Path: "/healthz",
										Port: intstr.FromString("healthz"),
									},
								},
							},
						},
					},
					TerminationGracePeriodSeconds: func() *int64 {
						terminationGracePeriodSeconds := int64(10)
						return &terminationGracePeriodSeconds
					}(),
				},
			},
		},
	}
}
