package k8s

import (
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/klog"
	infraazurev1 "sigs.k8s.io/cluster-api-provider-azure/api/v1alpha3"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1alpha3"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

var (
	scheme = runtime.NewScheme()

	mgr manager.Manager
)

func loadScheme() {
	clientgoscheme.AddToScheme(scheme)
	clusterv1.AddToScheme(scheme)
	infraazurev1.AddToScheme(scheme)
}

func Start() error {
	loadScheme()

	var err error
	mgr, err = ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme: scheme,
	})
	if err != nil {
		klog.Errorf("error to get k8s manager, %v", err)
	}
	return err
}

func GetManager() manager.Manager {
	return mgr
}
