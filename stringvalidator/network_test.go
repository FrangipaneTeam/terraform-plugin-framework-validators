package stringvalidator_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/FrangipaneTeam/terraform-plugin-framework-validators/stringvalidator"
)

func TestNetworkValidator(t *testing.T) {
	t.Parallel()

	type testCase struct {
		typesOfNetwork []stringvalidator.NetworkValidatorType
		val            types.String
		ComparatorOR   bool
		expectError    bool
	}
	tests := map[string]testCase{
		"unknown": {
			val: types.StringUnknown(),
		},
		"null": {
			val: types.StringNull(),
		},
		"valid-ipv4": {
			val: types.StringValue("1.1.1.1"),
			typesOfNetwork: []stringvalidator.NetworkValidatorType{
				stringvalidator.IPV4,
			},
		},
		"valid-ipv4-with-netmask": {
			val: types.StringValue("192.168.0.1/255.255.255.0"),
			typesOfNetwork: []stringvalidator.NetworkValidatorType{
				stringvalidator.IPv4WithNetmask,
			},
		},
		"valid-ipv4-with-cidr": {
			val: types.StringValue("192.168.0.1/24"),
			typesOfNetwork: []stringvalidator.NetworkValidatorType{
				stringvalidator.IPV4WithCIDR,
			},
		},
		"valid-rfc1918": {
			val: types.StringValue("192.168.0.1"),
			typesOfNetwork: []stringvalidator.NetworkValidatorType{
				stringvalidator.RFC1918,
			},
		},
		"invalid-rfc1918-valid-ipv4-with-cidr-comparatorOR": {
			val: types.StringValue("192.168.0.1/24"),
			typesOfNetwork: []stringvalidator.NetworkValidatorType{
				stringvalidator.RFC1918,
				stringvalidator.IPV4WithCIDR,
			},
			ComparatorOR: true,
			expectError:  false,
		},
		"invalid-rfc1918-valid-ipv4-with-cidr-comparatorAND": {
			val: types.StringValue("192.168.0.1/24"),
			typesOfNetwork: []stringvalidator.NetworkValidatorType{
				stringvalidator.RFC1918,
				stringvalidator.IPV4WithCIDR,
			},
			ComparatorOR: false,
			expectError:  true,
		},
		"no-types-of-networks": {
			val:         types.StringValue("1.1.1.1"),
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
			stringvalidator.IsNetwork(test.typesOfNetwork, test.ComparatorOR).ValidateString(context.TODO(), request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", response.Diagnostics)
			}
		})
	}
}
