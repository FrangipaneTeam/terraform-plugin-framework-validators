# StringPlanModifiers

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

- [`IsValidIP`](isvalidip.md) - This validator is used to check if the string is a valid IP address.
- [`IsValidNetmask`](isvalidnetmask.md) - This validator is used to check if the string is a valid netmask.

## String

- [`IsValidURN`](isvalidurn.md) - This validator is used to check if the string is a valid URN.
