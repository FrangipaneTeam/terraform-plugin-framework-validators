# `RequireIfAttributeIsOneOf`

!!! quote inline end "Released in v1.3.0"

This validator is used to require the attribute if another attribute is one of the given values.

## How to use it

```go
// Schema defines the schema for the resource.
func (r *xResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = schema.Schema{
        (...)
            "network_type": schema.StringAttribute{
                Optional:            true,
                MarkdownDescription: "Network type ...",
                Validators: []validator.String{
                    fstringvalidator.OneOf("public", "private"),
                },
            },
            "enabled": schema.StringAttribute{
                Optional:            true,
                MarkdownDescription: "Enable ...",
                Validators: []validator.String{
                    fboolvalidator.RequireIfAttributeIsOneOf(path.MatchRoot("network_type"),[]attr.Value{types.StringValue("private")})
                },
            },
```
