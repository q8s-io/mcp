package dto

import (
	"encoding/base64"

	"github.com/q8s-io/mcp/pkg/domain/entity"
)

type KubeconfigResp struct {
	Kubeconfig string `json:"kubeconfig" description:"Kubeconfig info with Base64 format"`
	Context    string `json:"context" description:"Context of kubeconfig"`
}

func (r *KubeconfigResp) ConvertFromKubeconfig(kubeconfig *entity.Kubeconfig) error {
	kubeconfigBytes := base64.StdEncoding.EncodeToString([]byte(kubeconfig.Kubeconfig))

	r.Kubeconfig = kubeconfigBytes
	r.Context = kubeconfig.Context
	return nil
}
