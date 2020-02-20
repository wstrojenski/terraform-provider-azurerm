package parse

import (
	"fmt"

	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
)

type AutomationScheduleId struct {
	ResourceGroup string
	AccountName   string
	Name          string
}

func AutomationScheduleID(input string) (*AutomationScheduleId, error) {
	id, err := azure.ParseAzureResourceID(input)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Unable to parse Automation Schedule ID %q: %+v", input, err)
	}

	schedule := AutomationScheduleId{
		ResourceGroup: id.ResourceGroup,
	}

	if schedule.AccountName, err = id.PopSegment("automationAccounts"); err != nil {
		return nil, err
	}

	if schedule.Name, err = id.PopSegment("schedules"); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &schedule, nil
}
