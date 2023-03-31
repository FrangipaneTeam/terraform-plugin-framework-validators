# BoolValidator

Bool validator are used to validate the plan of a boolean attribute.
It will be used into the `Validators` field of the `schema.BoolAttribute` struct.

## How to use it

```go
import (
    fboolvalidator "github.com/FrangipaneTeam/terraform-plugin-framework-validators/boolvalidator"
)
```

## List of Validators

- [`RequireIfAttributeIsOneOf`](../common/require_if_attribute_is_one_of.md) - This validator is used to require the attribute if another attribute is one of the given values.

## Special

- [`Not`](not.md) - This validator is used to negate the result of another validator.
