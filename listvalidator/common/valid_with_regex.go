package common

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.List = RegexValidator{}

type RegexValidator struct {
	Desc  string
	Regex string

	ErrorSummary string
	ErrorDetail  string
}

// Description describes the validation in plain text formatting.
func (validator RegexValidator) Description(_ context.Context) string {
	return validator.Desc
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator RegexValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

// Validate performs the validation.
func (validator RegexValidator) ValidateList(
	ctx context.Context,
	request validator.ListRequest,
	response *validator.ListResponse,
) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}
	elems := request.ConfigValue.Elements()
	re := regexp.MustCompile(validator.Regex)
	for _, elem := range elems {
		if !re.MatchString(elem.String()) {
			response.Diagnostics.AddAttributeError(
				request.Path,
				validator.ErrorSummary,
				validator.ErrorDetail,
			)
		}
	}
}
