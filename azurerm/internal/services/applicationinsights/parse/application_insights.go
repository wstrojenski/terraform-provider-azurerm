package parse

import (
	"fmt"

	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
)

type ApplicationInsightsId struct {
	ResourceGroup string
	Name          string
}

func ApplicationInsightsID(input string) (*ApplicationInsightsId, error) {
	id, err := azure.ParseAzureResourceID(input)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Unable to parse Application Insights ID %q: %+v", input, err)
	}

	server := ApplicationInsightsId{
		ResourceGroup: id.ResourceGroup,
	}

	if server.Name, err = id.PopSegment("components"); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &server, nil
}
