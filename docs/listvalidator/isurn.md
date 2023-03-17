# `IsURN`

This validator is used to check if the list attribute contains valid URNs.

## How to use it

```go
// Schema defines the schema for the resource.
func (r *xResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = schema.Schema{
        (...)
            "vm_ids": schema.ListAttribute{
              Required:            true,
              ElementType:         types.StringType,
              MarkdownDescription: "Set of VM IDs assigned to this rule",
              Validators: []validator.List{
                flistvalidator.IsURN(),
              },
            },
`
