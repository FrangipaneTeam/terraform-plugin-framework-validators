# ListValidator

List validator are used to validate the plan of a list attribute.
I will be used into the `Validators` field of the `schema.ListAttribute` struct.

## How to use it

```go
import (
    flistvalidator "github.com/FrangipaneTeam/terraform-plugin-framework-validators/listvalidator"
)
```

## List of Validators

## Special

- [`Not`](not.md) - This validator is used to negate the result of another validator.
