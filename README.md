# terraform-plugin-framework-validators

This repository contains the validators for the Terraform Plugin Framework.

## List of Validators

### String Validators

- `IsValidIP` - Validates that the value is a valid IP address.
- `IsValidNetmask` - Validates that the value is a valid netmask.

## Who to use

The validators are designed to be used with the `Validators` field.

### String Validator

```go
    import (
        fstringvalidator "github.com/FrangipaneTeam/terraform-plugin-framework-validators/stringvalidator"
    )

    [...]
    "gateway": schema.StringAttribute{
        [...]
        Validators: []validator.String{
            fstringvalidator.IsValidIP(),
        },
    },
```

## Maintainers

This repository is maintained by the [Frangipane Team](https://github.com/FrangipaneTeam)
