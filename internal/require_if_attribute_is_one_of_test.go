package internal_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"

	"github.com/FrangipaneTeam/terraform-plugin-framework-validators/internal"
)

func TestRequireIfAttributeIsOneOfValidator(t *testing.T) {
	t.Parallel()

	type testCase struct {
		req            internal.RequireIfAttributeIsOneOfRequest
		in             path.Expression
		exceptedValues []attr.Value
		expError       bool
	}

	testCases := map[string]testCase{
		"baseString": {
			req: internal.RequireIfAttributeIsOneOfRequest{
				ConfigValue:    types.StringValue("bar value"),
				Path:           path.Root("bar"),
				PathExpression: path.MatchRoot("bar"),
				Config: tfsdk.Config{
					Schema: schema.Schema{
						Attributes: map[string]schema.Attribute{
							"foo": schema.StringAttribute{},
							"bar": schema.StringAttribute{},
						},
					},
					Raw: tftypes.NewValue(tftypes.Object{
						AttributeTypes: map[string]tftypes.Type{
							"foo": tftypes.String,
							"bar": tftypes.String,
						},
					}, map[string]tftypes.Value{
						"foo": tftypes.NewValue(tftypes.String, "excepted value"),
						"bar": tftypes.NewValue(tftypes.String, "bar value"),
					}),
				},
			},
			in: path.MatchRoot("foo"),
			exceptedValues: []attr.Value{
				types.StringValue("excepted value"),
			},
			expError: false,
		},
		"baseInt64": {
			req: internal.RequireIfAttributeIsOneOfRequest{
				ConfigValue:    types.StringValue("bar value"),
				Path:           path.Root("bar"),
				PathExpression: path.MatchRoot("bar"),
				Config: tfsdk.Config{
					Schema: schema.Schema{
						Attributes: map[string]schema.Attribute{
							"foo": schema.Int64Attribute{},
							"bar": schema.StringAttribute{},
						},
					},
					Raw: tftypes.NewValue(tftypes.Object{
						AttributeTypes: map[string]tftypes.Type{
							"foo": tftypes.Number,
							"bar": tftypes.String,
						},
					}, map[string]tftypes.Value{
						"foo": tftypes.NewValue(tftypes.Number, 10),
						"bar": tftypes.NewValue(tftypes.String, "bar value"),
					}),
				},
			},
			in: path.MatchRoot("foo"),
			exceptedValues: []attr.Value{
				types.Int64Value(10),
			},
			expError: false,
		},
		"baseBool": {
			req: internal.RequireIfAttributeIsOneOfRequest{
				ConfigValue:    types.StringValue("bar value"),
				Path:           path.Root("bar"),
				PathExpression: path.MatchRoot("bar"),
				Config: tfsdk.Config{
					Schema: schema.Schema{
						Attributes: map[string]schema.Attribute{
							"foo": schema.Int64Attribute{},
							"bar": schema.StringAttribute{},
						},
					},
					Raw: tftypes.NewValue(tftypes.Object{
						AttributeTypes: map[string]tftypes.Type{
							"foo": tftypes.Bool,
							"bar": tftypes.String,
						},
					}, map[string]tftypes.Value{
						"foo": tftypes.NewValue(tftypes.Bool, true),
						"bar": tftypes.NewValue(tftypes.String, "bar value"),
					}),
				},
			},
			in: path.MatchRoot("foo"),
			exceptedValues: []attr.Value{
				types.BoolValue(true),
			},
			expError: false,
		},
		"path-attribute-is-null": {
			req: internal.RequireIfAttributeIsOneOfRequest{
				ConfigValue:    types.StringValue("bar value"),
				Path:           path.Root("bar"),
				PathExpression: path.MatchRoot("bar"),
				Config: tfsdk.Config{
					Schema: schema.Schema{
						Attributes: map[string]schema.Attribute{
							"foo": schema.StringAttribute{},
							"bar": schema.StringAttribute{},
						},
					},
					Raw: tftypes.NewValue(tftypes.Object{
						AttributeTypes: map[string]tftypes.Type{
							"foo": tftypes.String,
							"bar": tftypes.String,
						},
					}, map[string]tftypes.Value{
						"foo": tftypes.NewValue(tftypes.String, attr.NullValueString),
						"bar": tftypes.NewValue(tftypes.String, "bar value"),
					}),
				},
			},
			in: path.MatchRoot("foo"),
			exceptedValues: []attr.Value{
				types.StringValue("excepted value"),
			},
			expError: false,
		},
		"config-attribute-is-null": {
			req: internal.RequireIfAttributeIsOneOfRequest{
				ConfigValue:    types.StringNull(),
				Path:           path.Root("bar"),
				PathExpression: path.MatchRoot("bar"),
				Config: tfsdk.Config{
					Schema: schema.Schema{
						Attributes: map[string]schema.Attribute{
							"foo": schema.StringAttribute{},
							"bar": schema.StringAttribute{},
						},
					},
					Raw: tftypes.NewValue(tftypes.Object{
						AttributeTypes: map[string]tftypes.Type{
							"foo": tftypes.String,
							"bar": tftypes.String,
						},
					}, map[string]tftypes.Value{
						"foo": tftypes.NewValue(tftypes.String, "excepted value"),
						"bar": tftypes.NewValue(tftypes.String, attr.NullValueString),
					}),
				},
			},
			in: path.MatchRoot("foo"),
			exceptedValues: []attr.Value{
				types.StringValue("excepted value"),
			},
			expError: true,
		},
		"config-attribute-is-null-and-path-attribute-not-match": {
			req: internal.RequireIfAttributeIsOneOfRequest{
				ConfigValue:    types.StringNull(),
				Path:           path.Root("bar"),
				PathExpression: path.MatchRoot("bar"),
				Config: tfsdk.Config{
					Schema: schema.Schema{
						Attributes: map[string]schema.Attribute{
							"foo": schema.StringAttribute{},
							"bar": schema.StringAttribute{},
						},
					},
					Raw: tftypes.NewValue(tftypes.Object{
						AttributeTypes: map[string]tftypes.Type{
							"foo": tftypes.String,
							"bar": tftypes.String,
						},
					}, map[string]tftypes.Value{
						"foo": tftypes.NewValue(tftypes.String, "excepted value"),
						"bar": tftypes.NewValue(tftypes.String, attr.NullValueString),
					}),
				},
			},
			in: path.MatchRoot("foo"),
			exceptedValues: []attr.Value{
				types.StringValue("test value"),
			},
			expError: false,
		},
		"unknown": {
			req: internal.RequireIfAttributeIsOneOfRequest{
				ConfigValue:    types.StringUnknown(),
				Path:           path.Root("bar"),
				PathExpression: path.MatchRoot("bar"),
				Config: tfsdk.Config{
					Schema: schema.Schema{
						Attributes: map[string]schema.Attribute{
							"foo": schema.StringAttribute{},
							"bar": schema.StringAttribute{},
						},
					},
					Raw: tftypes.NewValue(tftypes.Object{
						AttributeTypes: map[string]tftypes.Type{
							"foo": tftypes.String,
							"bar": tftypes.String,
						},
					}, map[string]tftypes.Value{
						"foo": tftypes.NewValue(tftypes.String, "excepted value"),
						"bar": tftypes.NewValue(tftypes.String, attr.UnknownValueString),
					}),
				},
			},
			in: path.MatchRoot("foo"),
			exceptedValues: []attr.Value{
				types.StringValue("test value"),
			},
			expError: false,
		},
	}

	for name, test := range testCases {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			res := &internal.RequireIfAttributeIsOneOfResponse{}

			internal.RequireIfAttributeIsOneOf{
				PathExpression: test.in,
				ExceptedValues: test.exceptedValues,
			}.Validate(context.TODO(), test.req, res)

			if test.expError && !res.Diagnostics.HasError() {
				t.Fatal("expected error(s), got none")
			}

			if !test.expError && res.Diagnostics.HasError() {
				t.Fatalf("unexpected error(s): %s", res.Diagnostics)
			}
		})
	}
}
