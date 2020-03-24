package validate

import (
	"fmt"
	"regexp"

	"github.com/Azure/azure-sdk-for-go/services/analysisservices/mgmt/2017-08-01/analysisservices"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func AnalysisServicesServerName(v interface{}, k string) (warnings []string, errors []error) {
	value := v.(string)

	if !regexp.MustCompile(`^[a-z][0-9a-z]{2,62}$`).Match([]byte(value)) {
		errors = append(errors, fmt.Errorf("%q must begin with a letter, be lowercase alphanumeric, and be between 3 and 63 characters in length", k))
	}

	return warnings, errors
}

func QuerypoolConnectionMode() schema.SchemaValidateFunc {
	connectionModes := make([]string, len(analysisservices.PossibleConnectionModeValues()))
	for i, v := range analysisservices.PossibleConnectionModeValues() {
		connectionModes[i] = string(v)
	}

	return validation.StringInSlice(connectionModes, true)
}
