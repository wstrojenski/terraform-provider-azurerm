package parse

import (
	"fmt"

	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
)

type TrafficManagerProfileId struct {
	ResourceGroup string
	Name          string
}

func TrafficManagerProfileID(input string) (*TrafficManagerProfileId, error) {
	id, err := azure.ParseAzureResourceID(input)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Unable to parse CDN Profile ID %q: %+v", input, err)
	}

	profile := TrafficManagerProfileId{
		ResourceGroup: id.ResourceGroup,
	}

	if profile.Name, err = id.PopSegment("trafficManagerProfiles"); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &profile, nil
}
