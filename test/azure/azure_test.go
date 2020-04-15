package azure

//import (
//	"context"
//	"os"
//	"testing"
//
//	"github.com/Azure/azure-sdk-for-go/services/preview/subscription/mgmt/2018-03-01-preview/subscription"
//	"github.com/Azure/go-autorest/autorest/azure/auth"
//
//	azure "sigs.k8s.io/cluster-api-provider-azure/cloud"
//	"sigs.k8s.io/cluster-api-provider-azure/cloud/services/availabilityzones"
//	"sigs.k8s.io/cluster-api-provider-azure/cloud/services/groups"
//)
//
//const (
//	subscriptionID = "593ea0a5-2089-4f6f-be30-ebe12fc78339"
//	defaultBaseURI = "https://management.chinacloudapi.cn"
//)
//
//func init() {
//	os.Setenv("AZURE_SUBSCRIPTION_ID", subscriptionID)
//	os.Setenv("AZURE_TENANT_ID", "c9572b54-e243-4caf-8684-cff70654c290")
//	os.Setenv("AZURE_CLIENT_ID", "8736e6cc-dfd4-415f-98d4-c92f1de063f5")
//	os.Setenv("AZURE_CLIENT_SECRET", "Jka@-iT8-H-yjyWvlzCOcbpo0huX36ns")
//	os.Setenv("AZURE_ENVIRONMENT", "AZURECHINACLOUD")
//}
//
//func TestEnvironment(t *testing.T) {
//	settings, err := auth.GetSettingsFromEnvironment()
//	if err != nil {
//		t.Error(err)
//	}
//
//	t.Log(settings.Values)
//	t.Log(settings.Environment.ResourceManagerEndpoint)
//}
//
//func TestLocations(t *testing.T) {
//	settings, err := auth.GetSettingsFromEnvironment()
//	if err != nil {
//		t.Error(err)
//	}
//	authorizer, err := settings.GetAuthorizer()
//	if err != nil {
//		t.Error(err)
//	}
//
//	azClient := subscription.NewSubscriptionsClientWithBaseURI(defaultBaseURI)
//	azClient.Authorizer = authorizer
//	azClient.AddToUserAgent(azure.UserAgent)
//
//	result, err := azClient.ListLocations(context.TODO(), subscriptionID)
//	if err != nil {
//		t.Error(err)
//	}
//	for _, location := range *result.Value {
//		t.Log(*location.Name)
//	}
//}
//
//func TestAvailabilityZones(t *testing.T) {
//	auth, err := auth.NewAuthorizerFromEnvironment()
//	if err != nil {
//		t.Error(err)
//	}
//
//	azClient := availabilityzones.NewClient(subscriptionID, auth)
//	list, err := azClient.ListComplete(context.TODO(), "")
//	if err != nil {
//		t.Error(err)
//	}
//
//	for list.NotDone() {
//		resSku := list.Value()
//		t.Log("sku name: " + *resSku.Name)
//		for _, location := range *resSku.Locations {
//			t.Log(location)
//		}
//
//		err = list.NextWithContext(context.TODO())
//		if err != nil {
//			t.Error(err)
//		}
//	}
//}
//
//func TestGroups(t *testing.T) {
//	groupName := "test"
//
//	auth, err := auth.NewAuthorizerFromEnvironment()
//	if err != nil {
//		t.Error(err)
//	}
//
//	azClient := groups.NewClient(subscriptionID, auth)
//	group, err := azClient.Get(context.TODO(), groupName)
//	if err != nil {
//		t.Error(err)
//	}
//	t.Log(group)
//}
