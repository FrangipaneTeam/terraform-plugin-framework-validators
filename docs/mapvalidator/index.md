# MapValidator

Map validator are used to validate the plan of a map attribute.
It will be used into the `Validators` field of the `schema.MapAttribute` struct.

## How to use it

```go
import (
    fmapvalidator "github.com/FrangipaneTeam/terraform-plugin-framework-validators/mapvalidator"
)
```

## List of Validators

## Special

- [`Not`](not.md) - This validator is used to negate the result of another validator.

### String

Hashicorp provides a generic validator for strings. It uses the validators already defined in string to validate a list of strings.
It is available in the [hashicorp stringvalidator](https://github.com/hashicorp/terraform-plugin-framework-validators/tree/main) package.

Example of usage:

```go
    // Used within a Schema method of a DataSource, Provider, or Resource
    _ = schema.Schema{
        Attributes: map[string]schema.Attribute{
            "example_attr": schema.MapAttribute{
                ElementType: types.StringType,
                Required:    true,
                Validators: []validator.Map{
                    // Validate this Map must contain string values which are at least 3 characters.
                    mapvalidator.ValueStringsAre(fstringvalidator.IsUUID()),
                },
            },
        },
    }
```
