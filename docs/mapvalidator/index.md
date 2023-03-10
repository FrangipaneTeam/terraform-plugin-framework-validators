# MapValidator

Map validator are used to validate the plan of a map attribute.
I will be used into the `Validators` field of the `schema.MapAttribute` struct.

## How to use it

```go
import (
    fmapvalidator "github.com/FrangipaneTeam/terraform-plugin-framework-validators/mapvalidator"
)
```

## List of Validators

## Special

- [`Not`](not.md) - This validator is used to negate the result of another validator.
