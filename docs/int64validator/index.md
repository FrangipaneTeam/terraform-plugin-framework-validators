# Int64Validator

Int64 validator are used to validate the plan of a int64 attribute.
I will be used into the `Validators` field of the `schema.Int64Attribute` struct.

## How to use it

```go
import (
    fint64validator "github.com/FrangipaneTeam/terraform-plugin-framework-validators/int64validator"
)
```

## List of Validators

## Special

- [`Not`](not.md) - This validator is used to negate the result of another validator.
