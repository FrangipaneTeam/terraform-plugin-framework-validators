package int64validator

import (
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/FrangipaneTeam/terraform-plugin-framework-validators/internal"
)

type OneOfWithDescriptionValues struct {
	Value       int64
	Description string
}

// OneOfWithDescription checks that the int64 held in the attribute
// is one of the given `values`.
// The description of the value is used to generate advanced
// Description and MarkdownDescription messages.
func OneOfWithDescription(values ...OneOfWithDescriptionValues) validator.Int64 {
	frameworkValues := make([]internal.OneOfWithDescription, 0, len(values))

	for _, v := range values {
		frameworkValues = append(frameworkValues, internal.OneOfWithDescription{
			Value:       types.Int64Value(v.Value),
			Description: v.Description,
		})
	}

	return internal.OneOfWithDescriptionValidator{
		Values: frameworkValues,
	}
}
