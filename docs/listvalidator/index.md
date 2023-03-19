# ListValidator

List validator are used to validate the plan of a list attribute.
It will be used into the `Validators` field of the `schema.ListAttribute` struct.

## How to use it

```go
import (
    flistvalidator "github.com/FrangipaneTeam/terraform-plugin-framework-validators/listvalidator"
)
```

## List of Validators

## Special

- [`Not`](not.md) - This validator is used to negate the result of another validator.

## Generic

### String

Hashicorp provides a generic validator for strings. It uses the validators already defined in string to validate a list of strings.
It is available in the [hashicorp stringvalidator](https://github.com/hashicorp/terraform-plugin-framework-validators/tree/main) package.

Example of usage:

```go
// Used within a Schema method of a DataSource, Provider, or Resource
_ = schema.Schema{
    Attributes: map[string]schema.Attribute{
        "example_attr": schema.ListAttribute{
            ElementType: types.StringType,
            Required:    true,
            Validators: []validator.List{
                // Validate this List must contain string values which are at least 3 characters.
                listvalidator.ValueStringsAre(fstringvalidator.IsUUID()),
            },
        },
    },
}
```
