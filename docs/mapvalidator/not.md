# `Not`

!!! quote inline end "Released in v1.0.0"

This validator is used to check if the validators passed as arguments are NOT met.

## How to use it

```go
// Schema defines the schema for the resource.
func (r *xResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = schema.Schema{
        (...)
            "map_of": schema.MapAttribute{
                Optional:            true,
                MarkdownDescription: "Map of ...",
                Validators: []validator.Map{
                    fmapvalidator.Not(mapvalidator.Between(10,20))
                },
            },
```
