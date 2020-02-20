package parse

import (
	"fmt"

	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
)

type AutomationRunbookId struct {
	ResourceGroup string
	AccountName   string
	Name          string
}

func AutomationRunbookID(input string) (*AutomationRunbookId, error) {
	id, err := azure.ParseAzureResourceID(input)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Unable to parse Automation Runbook ID %q: %+v", input, err)
	}

	runbook := AutomationRunbookId{
		ResourceGroup: id.ResourceGroup,
	}

	if runbook.AccountName, err = id.PopSegment("automationAccounts"); err != nil {
		return nil, err
	}

	if runbook.Name, err = id.PopSegment("runbooks"); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &runbook, nil
}
