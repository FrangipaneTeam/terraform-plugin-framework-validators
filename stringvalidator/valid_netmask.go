package stringvalidator

import (
	"context"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = netmaskValidator{}

type netmaskValidator struct {
}

// Description describes the validation in plain text formatting.
func (validator netmaskValidator) Description(_ context.Context) string {
	return "must be a valid netmask"
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator netmaskValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

// Validate performs the validation.
func (validator netmaskValidator) ValidateString(
	ctx context.Context,
	request validator.StringRequest,
	response *validator.StringResponse,
) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	var re = regexp.MustCompile(`(?m)^(((255\.){3}(255|254|252|248|240|224|192|128|0+))|((255\.){2}(255|254|252|248|240|224|192|128|0+)\.0)|((255\.)(255|254|252|248|240|224|192|128|0+)(\.0+){2})|((255|254|252|248|240|224|192|128|0+)(\.0+){3}))$`)

	if !re.MatchString(request.ConfigValue.ValueString()) {
		response.Diagnostics.AddAttributeError(
			request.Path,
			"Failed to parse netmask",
			fmt.Sprintf("invalid value: %s", request.ConfigValue.String()),
		)
	}

}

// IsValidNetmask returns a validator which ensures that the configured attribute
// value is a valid Netmask.
// Null (unconfigured) and unknown (known after apply) values are skipped.
func IsValidNetmask() validator.String {
	return &netmaskValidator{}
}
