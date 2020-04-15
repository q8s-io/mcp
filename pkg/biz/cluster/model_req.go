package cluster

import (
	"encoding/base64"

	"github.com/q8s-io/mcp/pkg/biz/kubeconfig"
)

type AttachReq struct {
	Name       string `json:"name" validate:"required" description:"Name of attached cluster."`
	Kubeconfig string `json:"kubeconfig" validate:"required" description:"Kubeconfig info with Base64 format."`
	Context    string `json:"context" validate:"required" description:"Context of kubeconfig."`
}

func (r *AttachReq) toClusterKubeconfig() (*kubeconfig.Kubeconfig, error) {
	decodeBytes, err := base64.StdEncoding.DecodeString(r.Kubeconfig)
	if err != nil {
		return nil, err
	}

	return &kubeconfig.Kubeconfig{
		Kubeconfig: string(decodeBytes),
		Context:    r.Context,
	}, nil
}

func (r *AttachReq) toCluster() (*Cluster, error) {
	decodeBytes, err := base64.StdEncoding.DecodeString(r.Kubeconfig)
	if err != nil {
		return nil, err
	}

	return &Cluster{
		Name:     r.Name,
		Provider: AttachedProvider,
		Status:   StatusOk,
		Kubeconfig: kubeconfig.Kubeconfig{
			Kubeconfig: string(decodeBytes),
			Context:    r.Context,
		},
	}, nil
}
