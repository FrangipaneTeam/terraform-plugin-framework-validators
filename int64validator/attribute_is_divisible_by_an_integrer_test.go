package int64validator

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestAttributeIsDivisibleByAnIntegerValidator(t *testing.T) {
	t.Parallel()

	type testCase struct {
		req      validator.Int64Request
		in       path.Expression
		expError bool
	}

	testCases := map[string]testCase{
		"work": {
			req: validator.Int64Request{
				ConfigValue:    types.Int64Value(6),
				Path:           path.Root("bar"),
				PathExpression: path.MatchRoot("bar"),
				Config: tfsdk.Config{
					Schema: schema.Schema{
						Attributes: map[string]schema.Attribute{
							"foo": schema.Int64Attribute{},
							"bar": schema.Int64Attribute{},
						},
					},
					Raw: tftypes.NewValue(tftypes.Object{
						AttributeTypes: map[string]tftypes.Type{
							"foo": tftypes.Number,
							"bar": tftypes.Number,
						},
					}, map[string]tftypes.Value{
						"foo": tftypes.NewValue(tftypes.Number, 12),
						"bar": tftypes.NewValue(tftypes.Number, 6),
					}),
				},
			},
			in:       path.MatchRoot("foo"),
			expError: false,
		},
		"not work": {
			req: validator.Int64Request{
				ConfigValue:    types.Int64Value(6),
				Path:           path.Root("bar"),
				PathExpression: path.MatchRoot("bar"),
				Config: tfsdk.Config{
					Schema: schema.Schema{
						Attributes: map[string]schema.Attribute{
							"foo": schema.Int64Attribute{},
							"bar": schema.Int64Attribute{},
						},
					},
					Raw: tftypes.NewValue(tftypes.Object{
						AttributeTypes: map[string]tftypes.Type{
							"foo": tftypes.Number,
							"bar": tftypes.Number,
						},
					}, map[string]tftypes.Value{
						"foo": tftypes.NewValue(tftypes.Number, 13),
						"bar": tftypes.NewValue(tftypes.Number, 6),
					}),
				},
			},
			in:       path.MatchRoot("foo"),
			expError: true,
		},
	}

	for name, test := range testCases {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			res := &validator.Int64Response{}

			x := attributeIsDivisibleByAnInteger{
				PathExpression: test.in,
			}

			x.ValidateInt64(context.TODO(), test.req, res)

			if test.expError && !res.Diagnostics.HasError() {
				t.Fatal("expected error(s), got none")
			}

			if !test.expError && res.Diagnostics.HasError() {
				t.Fatalf("unexpected error(s): %s", res.Diagnostics)
			}
		})
	}
}
