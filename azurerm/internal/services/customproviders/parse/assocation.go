package parse

import (
	"fmt"

	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
)

type AssociationId struct {
	ResourceId string
	Name       string
}

func AssociationID(input string) (*AssociationId, error) {
	id, err := azure.ParseAzureResourceID(input)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Unable to parse Association ID %q: %+v", input, err)
	}

	service := AssociationId{
		ResourceId: fmt.Sprintf("/subscriptions/%s/resourceGroups/%s", id.SubscriptionID, id.ResourceGroup),
	}

	if service.Name, err = id.PopSegment("associations"); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &service, nil
}
