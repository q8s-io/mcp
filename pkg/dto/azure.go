package dto

type AzureSubscriptionsResp []string

type AzureSubscriptionsReq struct {
	TenantID     string `json:"tenantId" validate:"required" description:"Tenant ID of Azure."`
	ClientID     string `json:"clientId" validate:"required" description:"Client ID of Azure."`
	ClientSecret string `json:"clientSecret" validate:"required" description:"Client secret of Azure."`
}
