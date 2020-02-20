package parse

import (
	"fmt"

	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
)

type AutomationJobScheduleId struct {
	ResourceGroup string
	AccountName   string
	JobScheduleID string
}

func AutomationJobScheduleID(input string) (*AutomationJobScheduleId, error) {
	id, err := azure.ParseAzureResourceID(input)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Unable to parse Automation Job Schedule ID %q: %+v", input, err)
	}

	schedule := AutomationJobScheduleId{
		ResourceGroup: id.ResourceGroup,
	}

	if schedule.AccountName, err = id.PopSegment("automationAccounts"); err != nil {
		return nil, err
	}

	if schedule.JobScheduleID, err = id.PopSegment("jobSchedules"); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &schedule, nil
}
