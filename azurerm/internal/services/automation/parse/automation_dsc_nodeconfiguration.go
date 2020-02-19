package parse

import (
	"fmt"

	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
)

type AutomationDscNodeConfigurationId struct {
	ResourceGroup string
	AccountName   string
	Name          string
}

func AutomationDscNodeConfigurationID(input string) (*AutomationDscNodeConfigurationId, error) {
	id, err := azure.ParseAzureResourceID(input)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Unable to parse Automation DSC Node Configuration ID %q: %+v", input, err)
	}

	configuration := AutomationDscNodeConfigurationId{
		ResourceGroup: id.ResourceGroup,
	}

	if configuration.AccountName, err = id.PopSegment("automationAccounts"); err != nil {
		return nil, err
	}

	if configuration.Name, err = id.PopSegment("nodeConfigurations"); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &configuration, nil
}
