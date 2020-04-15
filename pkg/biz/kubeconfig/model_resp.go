package kubeconfig

import (
	"encoding/base64"
)

type ClusterKubeconfigResp struct {
	Kubeconfig string `json:"kubeconfig" description:"Kubeconfig info with Base64 format"`
	Context    string `json:"context" description:"Context of kubeconfig"`
}

func (r *ClusterKubeconfigResp) convertFromKubeconfig(kubeconfig *Kubeconfig) error {
	kubeconfigBytes := base64.StdEncoding.EncodeToString([]byte(kubeconfig.Kubeconfig))

	r.Kubeconfig = kubeconfigBytes
	r.Context = kubeconfig.Context
	return nil
}
