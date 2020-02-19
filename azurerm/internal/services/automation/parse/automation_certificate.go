package parse

import (
	"fmt"

	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
)

type AutomationCertificateId struct {
	ResourceGroup string
	AccountName   string
	Name          string
}

func AutomationCertificateID(input string) (*AutomationCertificateId, error) {
	id, err := azure.ParseAzureResourceID(input)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Unable to parse Automation Certificate ID %q: %+v", input, err)
	}

	certificate := AutomationCertificateId{
		ResourceGroup: id.ResourceGroup,
	}

	if certificate.AccountName, err = id.PopSegment("automationAccounts"); err != nil {
		return nil, err
	}

	if certificate.Name, err = id.PopSegment("certificates"); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &certificate, nil
}
