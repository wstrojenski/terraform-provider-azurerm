package parse

import (
	"fmt"

	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
)

type TrafficManagerEndpointId struct {
	ResourceGroup string
	ProfileName   string
	EndpointType  string
	Name          string
}

func TrafficManagerEndpointID(input string) (*TrafficManagerEndpointId, error) {
	id, err := azure.ParseAzureResourceID(input)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Unable to parse CDN Endpoint ID %q: %+v", input, err)
	}

	endpoint := TrafficManagerEndpointId{
		ResourceGroup: id.ResourceGroup,
	}

	if endpoint.ProfileName, err = id.PopSegment("trafficManagerProfiles"); err != nil {
		return nil, err
	}

	if endpoint.EndpointType, err = id.PopSegment(""); err != nil {
		return nil, err
	}

	if endpoint.Name, err = id.PopSegment(""); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &endpoint, nil
}
