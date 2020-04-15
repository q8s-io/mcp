package main

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1alpha3"
	ctrl "sigs.k8s.io/controller-runtime"
)

var (
	scheme = runtime.NewScheme()
)

func init() {
	clusterv1.AddToScheme(scheme)
}

func main() {
	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme: scheme,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	cluster := clusterv1.Cluster{}
	err = mgr.GetClient().Get(context.TODO(), types.NamespacedName{Namespace: "default", Name: "test"}, &cluster)
	if err != nil {
		fmt.Println(err)
		return
	}

	// reconcile

	stop := make(chan struct{})
	if err := mgr.Start(stop); err != nil {
		fmt.Println(err)
		return
	}
}
