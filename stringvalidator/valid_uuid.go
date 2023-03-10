package stringvalidator

import (
	"context"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = uuidValidator{}

type uuidValidator struct{}

// Description describes the validation in plain text formatting.
func (validator uuidValidator) Description(_ context.Context) string {
	return "must be a valid UUID"
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator uuidValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

// Validate performs the validation.
func (validator uuidValidator) ValidateString(
	ctx context.Context,
	request validator.StringRequest,
	response *validator.StringResponse,
) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	re := regexp.MustCompile(`(?m)^\w{8}-\w{4}-\w{4}-\w{4}-\w{12}$`)

	if !re.MatchString(request.ConfigValue.ValueString()) {
		response.Diagnostics.AddAttributeError(
			request.Path,
			"Failed to parse UUID",
			fmt.Sprintf("invalid UUID: %s", request.ConfigValue.String()),
		)
	}
}

// IsValidUUID returns a validator which ensures that the configured attribute
// value is a valid (v4) UUID.
// Null (unconfigured) and unknown (known after apply) values are skipped.
func IsValidUUID() validator.String {
	return &uuidValidator{}
}
