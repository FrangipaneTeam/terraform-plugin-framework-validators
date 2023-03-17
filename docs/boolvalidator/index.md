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

## Special

- [`Not`](not.md) - This validator is used to negate the result of another validator.
