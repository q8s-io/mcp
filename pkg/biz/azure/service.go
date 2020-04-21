package azure

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/profiles/preview/preview/subscription/mgmt/subscription"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"k8s.io/klog"
)

type Service struct{}

var serviceInstance *Service

func GetService() *Service {
	if serviceInstance == nil {
		serviceInstance = &Service{}
	}
	return serviceInstance
}

// required params: AZURE_TENANT_ID、AZURE_CLIENT_ID、AZURE_CLIENT_SECRET
// AZURE_ENVIRONMENT
func (s *Service) subscriptions(subscriptionReq *SubscriptionsReq) (SubscriptionsListResp, error) {
	settings := auth.EnvironmentSettings{
		Values: map[string]string{
			auth.TenantID:     subscriptionReq.TenantID,
			auth.ClientID:     subscriptionReq.ClientID,
			auth.ClientSecret: subscriptionReq.ClientSecret,
		},
	}
	environment, err := azure.EnvironmentFromName(environmentName_CN)
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

	azClient := subscription.NewSubscriptionsClientWithBaseURI(defaultBaseURI_CN)
	azClient.Authorizer = authorizer

	result, err := azClient.List(context.TODO())
	if err != nil {
		return nil, err
	}

	subscriptionList := make(SubscriptionsListResp, len(result.Values()))
	for index, subscription := range result.Values() {
		subscriptionList[index] = *subscription.SubscriptionID
	}
	return subscriptionList, nil
}
