# StringValidator

String validator are used to validate the plan of a string attribute.
I will be used into the `Validators` field of the `schema.StringAttribute` struct.

## How to use it

```go
import (
    fstringvalidator "github.com/FrangipaneTeam/terraform-plugin-framework-validators/stringvalidator"
)
```

## List of Validators

### Network

- [`IsIP`](isip.md) - This validator is used to check if the string is a valid IP address.
- [`IsNetmask`](isnetmask.md) - This validator is used to check if the string is a valid netmask.

## String

- [`IsURN`](isurn.md) - This validator is used to check if the string is a valid URN.
- [`IsUUID`](isuuid.md) - This validator is used to check if the string is a valid UUID.

## Special

- [`Not`](not.md) - This validator is used to negate the result of another validator.
