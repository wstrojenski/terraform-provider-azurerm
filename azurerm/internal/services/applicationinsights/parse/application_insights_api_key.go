package parse

import (
	"fmt"

	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
)

type ApplicationInsightsApiKeyId struct {
	ResourceGroup string
	KeyID         string
	Name          string
}

func ApplicationInsightsApiKeyID(input string) (*ApplicationInsightsApiKeyId, error) {
	id, err := azure.ParseAzureResourceID(input)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Unable to parse Application Insights API Key ID %q: %+v", input, err)
	}

	component := ApplicationInsightsApiKeyId{
		ResourceGroup: id.ResourceGroup,
	}

	if component.KeyID, err = id.PopSegment("apikeys"); err != nil {
		return nil, err
	}

	if component.Name, err = id.PopSegment("components"); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &component, nil
}
