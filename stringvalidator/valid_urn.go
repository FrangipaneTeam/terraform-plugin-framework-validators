package stringvalidator

import (
	"context"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = urnValidator{}

type urnValidator struct{}

// Description describes the validation in plain text formatting.
func (validator urnValidator) Description(_ context.Context) string {
	return "must be a valid URN"
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator urnValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

// Validate performs the validation.
func (validator urnValidator) ValidateString(
	ctx context.Context,
	request validator.StringRequest,
	response *validator.StringResponse,
) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	re := regexp.MustCompile(`(?m)urn:[A-Za-z0-9][A-Za-z0-9-]{0,31}:([A-Za-z0-9()+,\-.:=@;$_!*']|%[0-9A-Fa-f]{2})+`)

	if !re.MatchString(request.ConfigValue.ValueString()) {
		response.Diagnostics.AddAttributeError(
			request.Path,
			"Failed to parse URN",
			fmt.Sprintf("invalid URN: %s", request.ConfigValue.String()),
		)
	}
}

// IsValidURN returns a validator which ensures that the configured attribute
// Null (unconfigured) and unknown (known after apply) values are skipped.
func IsValidURN() validator.String {
	return &urnValidator{}
}
