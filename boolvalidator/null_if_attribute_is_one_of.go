package boolvalidator

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/FrangipaneTeam/terraform-plugin-framework-validators/internal"
)

// NullIfAttributeIsOneOf checks if the path.Path attribute contains
// one of the exceptedValue attr.Value.
func NullIfAttributeIsOneOf(path path.Expression, exceptedValue []attr.Value) validator.Int64 {
	return internal.NullIfAttributeIsOneOf{
		PathExpression: path,
		ExceptedValues: exceptedValue,
	}
}
