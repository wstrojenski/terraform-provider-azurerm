package parse

import (
	"testing"
)

func TestAutomationRunbookId(t *testing.T) {
	testData := []struct {
		Name     string
		Input    string
		Expected *AutomationRunbookId
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
			Name:     "Missing Automation Accounts Value",
			Input:    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Automation/automationAccounts/",
			Expected: nil,
		},
		{
			Name:     "Automation Account ID",
			Input:    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Automation/automationAccounts/account1",
			Expected: nil,
		},
		{
			Name:     "Missing Runbook Value",
			Input:    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Automation/automationAccounts/account1/runbooks/",
			Expected: nil,
		},
		{
			Name:  "Automation Runbook ID",
			Input: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Automation/automationAccounts/account1/runbooks/runbook1",
			Expected: &AutomationRunbookId{
				ResourceGroup: "resGroup1",
				AccountName:   "account1",
				Name:          "runbook1",
			},
		},
		{
			Name:     "Wrong Casing",
			Input:    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Automation/automationAccounts/account1/Runbooks/runbook1",
			Expected: nil,
		},
	}

	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Name)

		actual, err := AutomationRunbookID(v.Input)
		if err != nil {
			if v.Expected == nil {
				continue
			}

			t.Fatalf("Expected a value but got an error: %s", err)
		}

		if actual.Name != v.Expected.Name {
			t.Fatalf("Expected %q but got %q for Name", v.Expected.Name, actual.Name)
		}
		if actual.AccountName != v.Expected.AccountName {
			t.Fatalf("Expected %q but got %q for Account Name", v.Expected.AccountName, actual.AccountName)
		}

		if actual.ResourceGroup != v.Expected.ResourceGroup {
			t.Fatalf("Expected %q but got %q for Resource Group", v.Expected.ResourceGroup, actual.ResourceGroup)
		}
	}
}
