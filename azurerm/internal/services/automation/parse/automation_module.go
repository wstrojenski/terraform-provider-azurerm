package parse

import (
	"fmt"

	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
)

type AutomationModuleId struct {
	ResourceGroup string
	AccountName   string
	Name          string
}

func AutomationModuleID(input string) (*AutomationModuleId, error) {
	id, err := azure.ParseAzureResourceID(input)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Unable to parse Automation Module ID %q: %+v", input, err)
	}

	module := AutomationModuleId{
		ResourceGroup: id.ResourceGroup,
	}

	if module.AccountName, err = id.PopSegment("automationAccounts"); err != nil {
		return nil, err
	}

	if module.Name, err = id.PopSegment("modules"); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &module, nil
}
