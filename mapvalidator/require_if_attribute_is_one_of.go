package mapvalidator

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/FrangipaneTeam/terraform-plugin-framework-validators/internal"
)

// RequireIfAttributeIsOneOf checks if the path.Path attribute contains
// one of the exceptedValue attr.Value.
func RequireIfAttributeIsOneOf(path path.Expression, exceptedValue []attr.Value) validator.Map {
	return internal.RequireIfAttributeIsOneOf{
		PathExpression: path,
		ExceptedValues: exceptedValue,
	}
}
