# SetValidator

Set validator are used to validate the plan of a set attribute.
It will be used into the `Validators` field of the `schema.SetAttribute` struct.

## How to use it

```go
import (
    fsetvalidator "github.com/FrangipaneTeam/terraform-plugin-framework-validators/setvalidator"
)
```

## List of Validators

Every `string` validators are available for maps thanks to a generic validator provided by Hashicorp. See the section below for more details.

- [`RequireIfAttributeIsOneOf`](../common/require_if_attribute_is_one_of.md) - This validator is used to require the attribute if another attribute is one of the given values.

### Special

- [`Not`](not.md) - This validator is used to negate the result of another validator.

## Generic

### String

Hashicorp provides a generic validator for strings. It uses the validators already defined in string to validate a list of strings.
It is available in the [hashicorp stringvalidator](https://github.com/hashicorp/terraform-plugin-framework-validators/tree/main) package.

Example of usage:

```go
_ = schema.Schema{
    Attributes: map[string]schema.Attribute{
        "example_attr": schema.SetAttribute{
            ElementType: types.StringType,
            Required:    true,
            Validators: []validator.Set{
                // Validate this Set must contain string values which are URNs.
                setvalidator.ValueStringsAre(fstringvalidator.IsURN())
            },
        },
    },
}
```
