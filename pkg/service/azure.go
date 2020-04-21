package service

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/profiles/preview/preview/subscription/mgmt/subscription"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"k8s.io/klog"

	"github.com/q8s-io/mcp/pkg/constants"
	"github.com/q8s-io/mcp/pkg/dto"
)

type Azure struct{}

func NewAzureService() *Azure {
	return &Azure{}
}

// required params: AZURE_TENANT_ID、AZURE_CLIENT_ID、AZURE_CLIENT_SECRET
// AZURE_ENVIRONMENT
func (a *Azure) Subscriptions(subscriptionReq *dto.AzureSubscriptionsReq) (dto.AzureSubscriptionsResp, error) {
	settings := auth.EnvironmentSettings{
		Values: map[string]string{
			auth.TenantID:     subscriptionReq.TenantID,
			auth.ClientID:     subscriptionReq.ClientID,
			auth.ClientSecret: subscriptionReq.ClientSecret,
		},
	}
	environment, err := azure.EnvironmentFromName(constants.EnvironmentNameCN)
	if err != nil {
		klog.Errorf("invalid environment name, %v", err)
		return nil, err
	}
	settings.Environment = environment
	settings.Values[auth.Resource] = settings.Environment.ResourceManagerEndpoint

	authorizer, err := settings.GetAuthorizer()
	if err != nil {
		return nil, err
	}

	azClient := subscription.NewSubscriptionsClientWithBaseURI(constants.DefaultBaseURICN)
	azClient.Authorizer = authorizer

	result, err := azClient.List(context.TODO())
	if err != nil {
		return nil, err
	}

	subscriptionList := make(dto.AzureSubscriptionsResp, len(result.Values()))
	for index, subscription := range result.Values() {
		subscriptionList[index] = *subscription.SubscriptionID
	}
	return subscriptionList, nil
}
