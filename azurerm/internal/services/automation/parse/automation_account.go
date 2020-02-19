package parse

import (
	"fmt"

	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
)

type AutomationAccountId struct {
	ResourceGroup string
	Name          string
}

func AutomationAccountID(input string) (*AutomationAccountId, error) {
	id, err := azure.ParseAzureResourceID(input)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Unable to parse Automation Account ID %q: %+v", input, err)
	}

	account := AutomationAccountId{
		ResourceGroup: id.ResourceGroup,
	}

	if account.Name, err = id.PopSegment("automationAccounts"); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &account, nil
}
