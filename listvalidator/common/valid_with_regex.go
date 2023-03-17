package common

import (
	"context"
	"fmt"
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
	for i, elem := range elems {
		if elem.IsNull() {
			response.Diagnostics.AddAttributeError(
				request.Path.AtListIndex(i),
				fmt.Sprintf("%s at index %d", validator.ErrorSummary, i),
				fmt.Sprintf("Element at index %d is null", i),
			)
			continue
		}

		if !re.MatchString(elem.String()) {
			response.Diagnostics.AddAttributeError(
				request.Path.AtListIndex(i),
				fmt.Sprintf("%s at index %d", validator.ErrorSummary, i),
				fmt.Sprintf("%s : %s", validator.ErrorDetail, elem.String()),
			)
		}
	}
}
