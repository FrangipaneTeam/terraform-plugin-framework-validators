package listvalidator_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/FrangipaneTeam/terraform-plugin-framework-validators/listvalidator"
)

func TestIsURNValidator(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		list                types.List
		expectedDiagnostics diag.Diagnostics
	}{
		"null-list": {
			list:                types.ListNull(types.StringType),
			expectedDiagnostics: nil,
		},
		"unknown-list": {
			list:                types.ListUnknown(types.StringType),
			expectedDiagnostics: nil,
		},
		"null-value": {
			list: types.ListValueMust(
				types.StringType,
				[]attr.Value{types.StringNull()},
			),
			expectedDiagnostics: diag.Diagnostics{
				diag.NewAttributeErrorDiagnostic(
					path.Root("test").AtListIndex(0),
					"Failed to parse URN at index 0",
					"Element at index 0 is null",
				),
			},
		},
		"invalid-urn": {
			list: types.ListValueMust(
				types.StringType,
				[]attr.Value{types.StringValue("4aeb40d8-038c-4e77-8181-a7054f583b12")},
			),
			expectedDiagnostics: diag.Diagnostics{
				diag.NewAttributeErrorDiagnostic(
					path.Root("test").AtListIndex(0),
					"Failed to parse URN at index 0",
					`This value is not a valid URN : "4aeb40d8-038c-4e77-8181-a7054f583b12"`,
				),
			},
		},
		"valid-urn": {
			list: types.ListValueMust(
				types.StringType,
				[]attr.Value{types.StringValue("urn:test:demo:4aeb40d8-038c-4e77-8181-a7054f583b12")},
			),
			expectedDiagnostics: nil,
		},
		"with-invalid-urn": {
			list: types.ListValueMust(
				types.StringType,
				[]attr.Value{types.StringValue("urn:test:demo:4aeb40d8-038c-4e77-8181-a7054f583b12"), types.StringValue("4aeb40d8-038c-4e77-8181-a7054f583b12")},
			),
			expectedDiagnostics: diag.Diagnostics{
				diag.NewAttributeErrorDiagnostic(
					path.Root("test").AtListIndex(1),
					"Failed to parse URN at index 1",
					`This value is not a valid URN : "4aeb40d8-038c-4e77-8181-a7054f583b12"`,
				),
			},
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			request := validator.ListRequest{
				Path:           path.Root("test"),
				PathExpression: path.MatchRoot("test"),
				ConfigValue:    test.list,
			}
			response := validator.ListResponse{}
			listvalidator.IsURN().ValidateList(context.Background(), request, &response)

			if diff := cmp.Diff(response.Diagnostics, test.expectedDiagnostics); diff != "" {
				t.Errorf("unexpected diagnostics difference: %s", diff)
			}
		})
	}
}
