package stringvalidator

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/FrangipaneTeam/terraform-plugin-framework-validators/internal"
)

type OneOfWithDescriptionIfAttributeIsOneOfValues struct {
	Value       string
	Description string
}

// OneOfWithDescriptionIfAttributeIsOneOf checks that the String value is one of the expected values if the attribute is one of the exceptedValue.
// The description of the value is used to generate advanced
// Description and MarkdownDescription messages.
func OneOfWithDescriptionIfAttributeIsOneOf(path path.Expression, exceptedValue []attr.Value, values ...OneOfWithDescriptionIfAttributeIsOneOfValues) validator.String {
	frameworkValues := make([]internal.OneOfWithDescriptionIfAttributeIsOneOf, 0, len(values))

	for _, v := range values {
		frameworkValues = append(frameworkValues, internal.OneOfWithDescriptionIfAttributeIsOneOf{
			Value:       types.StringValue(v.Value),
			Description: v.Description,
		})
	}

	return internal.OneOfWithDescriptionIfAttributeIsOneOfValidator{
		Values:         frameworkValues,
		ExceptedValues: exceptedValue,
		PathExpression: path,
	}
}
