package azure

type SubscriptionsReq struct {
	TenantID     string `json:"tenant_id" validate:"required" description:"Tenant ID of Azure."`
	ClientID     string `json:"client_id" validate:"required" description:"Client ID of Azure."`
	ClientSecret string `json:"client_secret" validate:"required" description:"Client secret of Azure."`
}
