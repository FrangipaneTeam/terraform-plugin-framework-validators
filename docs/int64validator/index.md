# Int64Validator

Int64 validator are used to validate the plan of a int64 attribute.
It will be used into the `Validators` field of the `schema.Int64Attribute` struct.

## How to use it

```go
import (
    fint64validator "github.com/FrangipaneTeam/terraform-plugin-framework-validators/int64validator"
)
```

## List of Validators

- [`RequireIfAttributeIsOneOf`](../common/require_if_attribute_is_one_of.md) - This validator is used to require the attribute if another attribute is one of the given values.
- [`AttributeIsDivisibleByAnInteger`](attribute_is_divisible_by_an_integer.md) - This validator is used to validate that the attribute is divisible by an integer.

## Special

- [`Not`](not.md) - This validator is used to negate the result of another validator.
