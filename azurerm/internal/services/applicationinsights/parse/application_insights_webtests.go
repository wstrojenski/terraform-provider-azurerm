package parse

import (
	"fmt"

	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
)

type ApplicationInsightsWebtestsId struct {
	ResourceGroup string
	Name          string
}

func ApplicationInsightsWebtestsID(input string) (*ApplicationInsightsWebtestsId, error) {
	id, err := azure.ParseAzureResourceID(input)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Unable to parse Application Insights ID %q: %+v", input, err)
	}

	server := ApplicationInsightsWebtestsId{
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
