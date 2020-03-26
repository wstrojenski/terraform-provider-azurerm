package client

import (
	"github.com/Azure/azure-sdk-for-go/services/preview/customproviders/mgmt/2018-09-01-preview/customproviders"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/common"
)

type Client struct {
	CustomProviderClient *customproviders.CustomResourceProviderClient
	AssociationClient    *customproviders.AssociationsClient
}

func NewClient(o *common.ClientOptions) *Client {
	CustomProviderClient := customproviders.NewCustomResourceProviderClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&CustomProviderClient.Client, o.ResourceManagerAuthorizer)

	AssociationClient := customproviders.NewAssociationsClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&AssociationClient.Client, o.ResourceManagerAuthorizer)

	return &Client{
		CustomProviderClient: &CustomProviderClient,
		AssociationClient:    &AssociationClient,
	}
}
