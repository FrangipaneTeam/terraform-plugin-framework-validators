package internal_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
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
		req             internal.RequireIfAttributeIsOneOfRequest
		in              path.Expression
		inPath          path.Path
		exceptedValues  []attr.Value
		expError        bool
		expErrorMessage string
	}

	testCases := map[string]testCase{
		"baseString": {
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
			in:     path.MatchRoot("foo"),
			inPath: path.Root("foo"),
			exceptedValues: []attr.Value{
				types.StringValue("excepted value"),
			},
			expError:        true,
			expErrorMessage: "If foo attribute is set and the value is one of \"excepted value\", this attribute is required",
		},
		"extendedString": {
			req: internal.RequireIfAttributeIsOneOfRequest{
				ConfigValue:    types.StringNull(),
				Path:           path.Root("foobar").AtListIndex(0).AtName("bar2"),
				PathExpression: path.MatchRoot("foobar").AtListIndex(0).AtName("bar2"),
				Config: tfsdk.Config{
					Schema: schema.Schema{
						Attributes: map[string]schema.Attribute{
							"foo": schema.StringAttribute{},
							"bar": schema.StringAttribute{},
							"foobar": schema.ListNestedAttribute{
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"bar1": schema.StringAttribute{},
										"bar2": schema.StringAttribute{},
									},
								},
							},
						},
					},
					Raw: tftypes.NewValue(tftypes.Object{
						AttributeTypes: map[string]tftypes.Type{
							"foo": tftypes.String,
							"bar": tftypes.String,
							"foobar": tftypes.List{
								ElementType: tftypes.Object{
									AttributeTypes: map[string]tftypes.Type{
										"bar1": tftypes.String,
										"bar2": tftypes.String,
									},
								},
							},
						},
					}, map[string]tftypes.Value{
						"foo": tftypes.NewValue(tftypes.String, "excepted value"),
						"bar": tftypes.NewValue(tftypes.String, "bar value"),
						"foobar": tftypes.NewValue(tftypes.List{
							ElementType: tftypes.Object{
								AttributeTypes: map[string]tftypes.Type{
									"bar1": tftypes.String,
									"bar2": tftypes.String,
								},
							},
						}, []tftypes.Value{
							tftypes.NewValue(tftypes.Object{
								AttributeTypes: map[string]tftypes.Type{
									"bar1": tftypes.String,
									"bar2": tftypes.String,
								},
							}, map[string]tftypes.Value{
								"bar1": tftypes.NewValue(tftypes.String, "bar1 excepted value"),
								"bar2": tftypes.NewValue(tftypes.String, attr.NullValueString),
							}),
						},
						),
					}),
				},
			},
			in:     path.MatchRoot("foobar").AtListIndex(0).AtName("bar1"),
			inPath: path.Root("foobar").AtListIndex(0).AtName("bar1"),
			exceptedValues: []attr.Value{
				types.StringValue("bar1 excepted value"),
			},
			expError:        true,
			expErrorMessage: "If foobar[0].bar1 attribute is set and the value is one of \"bar1 excepted value\", this attribute is required",
		},
		"baseInt64": {
			req: internal.RequireIfAttributeIsOneOfRequest{
				ConfigValue:    types.StringNull(),
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
						"foo": tftypes.NewValue(tftypes.Number, int64(10)),
						"bar": tftypes.NewValue(tftypes.String, attr.NullValueString),
					}),
				},
			},
			in:     path.MatchRoot("foo"),
			inPath: path.Root("foo"),
			exceptedValues: []attr.Value{
				types.Int64Value(10),
			},
			expError:        true,
			expErrorMessage: "If foo attribute is set and the value is one of 10, this attribute is required",
		},
		"baseBool": {
			req: internal.RequireIfAttributeIsOneOfRequest{
				ConfigValue:    types.StringNull(),
				Path:           path.Root("bar"),
				PathExpression: path.MatchRoot("bar"),
				Config: tfsdk.Config{
					Schema: schema.Schema{
						Attributes: map[string]schema.Attribute{
							"foo": schema.BoolAttribute{},
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
						"bar": tftypes.NewValue(tftypes.String, attr.NullValueString),
					}),
				},
			},
			in:     path.MatchRoot("foo"),
			inPath: path.Root("foo"),
			exceptedValues: []attr.Value{
				types.BoolValue(true),
			},
			expError:        true,
			expErrorMessage: "If foo attribute is set and the value is one of true, this attribute is required",
		},
		"path-attribute-is-null": {
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
						"foo": tftypes.NewValue(tftypes.String, attr.NullValueString),
						"bar": tftypes.NewValue(tftypes.String, attr.NullValueString),
					}),
				},
			},
			in:     path.MatchRoot("foo"),
			inPath: path.Root("foo"),
			exceptedValues: []attr.Value{
				types.StringValue("excepted value"),
			},
			expError: false,
		},
		"config-attribute-is-set": {
			req: internal.RequireIfAttributeIsOneOfRequest{
				ConfigValue:    types.StringValue("excepted value"),
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
			in:     path.MatchRoot("foo"),
			inPath: path.Root("foo"),
			exceptedValues: []attr.Value{
				types.StringValue("excepted value"),
			},
			expError: false,
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
			in:     path.MatchRoot("foo"),
			inPath: path.Root("foo"),
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
			in:     path.MatchRoot("foo"),
			inPath: path.Root("foo"),
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

			if test.expError && res.Diagnostics.HasError() {
				if !res.Diagnostics.Contains(diag.NewAttributeErrorDiagnostic(
					test.inPath,
					fmt.Sprintf("Invalid configuration for attribute %s", test.req.Path),
					test.expErrorMessage,
				)) {
					t.Fatal("expected error(s), got none. Error message is : ", res.Diagnostics.Errors())
				}
			}

			if !test.expError && res.Diagnostics.HasError() {
				t.Fatalf("unexpected error(s): %s", res.Diagnostics)
			}
		})
	}
}
