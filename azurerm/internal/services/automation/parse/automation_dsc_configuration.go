package parse

import (
	"fmt"

	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
)

type AutomationDscConfigurationId struct {
	ResourceGroup string
	AccountName   string
	Name          string
}

func AutomationDscConfigurationID(input string) (*AutomationDscConfigurationId, error) {
	id, err := azure.ParseAzureResourceID(input)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Unable to parse Automation DSC Configuration ID %q: %+v", input, err)
	}

	configuration := AutomationDscConfigurationId{
		ResourceGroup: id.ResourceGroup,
	}

	if configuration.AccountName, err = id.PopSegment("automationAccounts"); err != nil {
		return nil, err
	}

	if configuration.Name, err = id.PopSegment("configurations"); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &configuration, nil
}
