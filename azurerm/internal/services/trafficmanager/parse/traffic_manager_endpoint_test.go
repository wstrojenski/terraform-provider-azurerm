package parse

import (
	"testing"
)

func TestTrafficManagerEndpointId(t *testing.T) {
	testData := []struct {
		Name     string
		Input    string
		Expected *TrafficManagerEndpointId
	}{
		{
			Name:     "Empty",
			Input:    "",
			Expected: nil,
		},
		{
			Name:     "No Resource Groups Segment",
			Input:    "/subscriptions/00000000-0000-0000-0000-000000000000",
			Expected: nil,
		},
		{
			Name:     "No Resource Groups Value",
			Input:    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/",
			Expected: nil,
		},
		{
			Name:     "Resource Group ID",
			Input:    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/foo/",
			Expected: nil,
		},
		{
			Name:     "Missing Profiles Value",
			Input:    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Network/trafficmanagerprofiles/",
			Expected: nil,
		},
		{
			Name:     "Traffic Manager Profile ID",
			Input:    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Network/trafficmanagerprofiles/profName1",
			Expected: nil,
		},
		{
			Name:     "Missing Endpoints Value",
			Input:    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Network/trafficmanagerprofiles/profName1/type1/",
			Expected: nil,
		},
		{
			Name:  "Traffic Manager Endpoint ID",
			Input: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Network/trafficmanagerprofiles/profName1/type1/Endpoint1",
			Expected: &TrafficManagerEndpointId{
				ResourceGroup: "resGroup1",
				ProfileName:   "profName1",
				EndpointType:  "type1",
				Name:          "Endpoint1",
			},
		},
		{
			Name:     "Wrong Casing",
			Input:    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Cdn/profiles/profName1/type1/Endpoint1",
			Expected: nil,
		},
	}

	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Name)

		actual, err := TrafficManagerEndpointID(v.Input)
		if err != nil {
			if v.Expected == nil {
				continue
			}

			t.Fatalf("Expected a value but got an error: %s", err)
		}

		if actual.Name != v.Expected.Name {
			t.Fatalf("Expected %q but got %q for Name", v.Expected.Name, actual.Name)
		}
		if actual.EndpointType != v.Expected.EndpointType {
			t.Fatalf("Expected %q but got %q for EndpointType", v.Expected.EndpointType, actual.EndpointType)
		}
		if actual.ProfileName != v.Expected.ProfileName {
			t.Fatalf("Expected %q but got %q for ProfileName", v.Expected.ProfileName, actual.ProfileName)
		}
		if actual.ResourceGroup != v.Expected.ResourceGroup {
			t.Fatalf("Expected %q but got %q for Resource Group", v.Expected.ResourceGroup, actual.ResourceGroup)
		}
	}
}
