package parse

import (
	"fmt"

	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
)

type AutomationCredentialId struct {
	ResourceGroup string
	AccountName   string
	Name          string
}

func AutomationCredentialID(input string) (*AutomationCredentialId, error) {
	id, err := azure.ParseAzureResourceID(input)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Unable to parse Automation Credential ID %q: %+v", input, err)
	}

	credential := AutomationCredentialId{
		ResourceGroup: id.ResourceGroup,
	}

	if credential.AccountName, err = id.PopSegment("automationAccounts"); err != nil {
		return nil, err
	}

	if credential.Name, err = id.PopSegment("credentials"); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &credential, nil
}
