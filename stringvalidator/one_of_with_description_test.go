package stringvalidator_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/FrangipaneTeam/terraform-plugin-framework-validators/stringvalidator"
)

func TestOneOfWithDescriptionValidator(t *testing.T) {
	t.Parallel()

	type testCase struct {
		in        types.String
		validator validator.String
		expErrors int
	}

	testCases := map[string]testCase{
		"simple-match": {
			in: types.StringValue("foo"),
			validator: stringvalidator.OneOfWithDescription(
				stringvalidator.OneOfWithDescriptionValues{
					Value:       "foo",
					Description: "foo description",
				},
				stringvalidator.OneOfWithDescriptionValues{
					Value:       "bar",
					Description: "bar description",
				},
				stringvalidator.OneOfWithDescriptionValues{
					Value:       "baz",
					Description: "baz description",
				},
			),
			expErrors: 0,
		},
		"simple-mismatch-case-insensitive": {
			in: types.StringValue("foo"),
			validator: stringvalidator.OneOfWithDescription(
				stringvalidator.OneOfWithDescriptionValues{
					Value:       "FOO",
					Description: "foo description",
				},
				stringvalidator.OneOfWithDescriptionValues{
					Value:       "bar",
					Description: "bar description",
				},
				stringvalidator.OneOfWithDescriptionValues{
					Value:       "baz",
					Description: "baz description",
				},
			),
			expErrors: 1,
		},
		"simple-mismatch": {
			in: types.StringValue("foz"),
			validator: stringvalidator.OneOfWithDescription(
				stringvalidator.OneOfWithDescriptionValues{
					Value:       "foo",
					Description: "foo description",
				},
				stringvalidator.OneOfWithDescriptionValues{
					Value:       "bar",
					Description: "bar description",
				},
				stringvalidator.OneOfWithDescriptionValues{
					Value:       "baz",
					Description: "baz description",
				},
			),
			expErrors: 1,
		},
		"skip-validation-on-null": {
			in: types.StringNull(),
			validator: stringvalidator.OneOfWithDescription(
				stringvalidator.OneOfWithDescriptionValues{
					Value:       "foo",
					Description: "foo description",
				},
				stringvalidator.OneOfWithDescriptionValues{
					Value:       "bar",
					Description: "bar description",
				},
				stringvalidator.OneOfWithDescriptionValues{
					Value:       "baz",
					Description: "baz description",
				},
			),
			expErrors: 0,
		},
		"skip-validation-on-unknown": {
			in: types.StringUnknown(),
			validator: stringvalidator.OneOfWithDescription(
				stringvalidator.OneOfWithDescriptionValues{
					Value:       "foo",
					Description: "foo description",
				},
				stringvalidator.OneOfWithDescriptionValues{
					Value:       "bar",
					Description: "bar description",
				},
				stringvalidator.OneOfWithDescriptionValues{
					Value:       "baz",
					Description: "baz description",
				},
			),
			expErrors: 0,
		},
	}

	for name, test := range testCases {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			req := validator.StringRequest{
				ConfigValue: test.in,
			}
			res := validator.StringResponse{}
			test.validator.ValidateString(context.TODO(), req, &res)

			if test.expErrors > 0 && !res.Diagnostics.HasError() {
				t.Fatalf("expected %d error(s), got none", test.expErrors)
			}

			if test.expErrors > 0 && test.expErrors != res.Diagnostics.ErrorsCount() {
				t.Fatalf("expected %d error(s), got %d: %v", test.expErrors, res.Diagnostics.ErrorsCount(), res.Diagnostics)
			}

			if test.expErrors == 0 && res.Diagnostics.HasError() {
				t.Fatalf("expected no error(s), got %d: %v", res.Diagnostics.ErrorsCount(), res.Diagnostics)
			}
		})
	}
}
