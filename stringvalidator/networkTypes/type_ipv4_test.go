package networktypes_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	networktypes "github.com/FrangipaneTeam/terraform-plugin-framework-validators/stringvalidator/networkTypes"
)

func TestValidIPV4Validator(t *testing.T) {
	t.Parallel()

	type testCase struct {
		val         types.String
		expectError bool
	}
	tests := map[string]testCase{
		"unknown": {
			val: types.StringUnknown(),
		},
		"null": {
			val: types.StringNull(),
		},
		"valid": {
			val: types.StringValue("192.168.1.1"),
		},
		"invalid": {
			val:         types.StringValue("192.168.1"),
			expectError: true,
		},
		"ipv6": {
			val:         types.StringValue("2001:0db8:85a3:0000:0000:8a2e:0370:7334"),
			expectError: true,
		},
		"multiple byte characters": {
			// Rightwards Arrow Over Leftwards Arrow (U+21C4; 3 bytes)
			val:         types.StringValue("⇄"),
			expectError: true,
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			request := validator.StringRequest{
				ConfigValue: test.val,
			}
			response := validator.StringResponse{}
			networktypes.IsIPV4().ValidateString(context.TODO(), request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", response.Diagnostics)
			}
		})
	}
}

// TestValidIPValidatorDescription
func TestValidIPV4ValidatorDescription(t *testing.T) {
	t.Parallel()

	type testCase struct {
		description string
	}
	tests := map[string]testCase{
		"description": {
			description: "a valid IPV4 address (192.168.0.1).",
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			validator := networktypes.IsIPV4()
			if validator.Description(context.Background()) != test.description {
				t.Fatalf("got unexpected description: %s != %s", validator.Description(context.Background()), test.description)
			}
		})
	}
}

// TestValidIPValidatorMarkdownDescription
func TestValidIPV4ValidatorMarkdownDescription(t *testing.T) {
	t.Parallel()

	type testCase struct {
		description string
	}
	tests := map[string]testCase{
		"description": {
			description: "a valid IPV4 address (`192.168.0.1`).",
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			validator := networktypes.IsIPV4()
			if validator.MarkdownDescription(context.Background()) != test.description {
				t.Fatalf("got unexpected description: %s != %s", validator.MarkdownDescription(context.Background()), test.description)
			}
		})
	}
}
