// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezerotrustgatewaypolicy


type DataCloudflareZeroTrustGatewayPolicyFilter struct {
	// Sort direction.
	//
	// When `order_by` is omitted, this controls the direction
	// of the existing precedence ordering. Shared rules remain first in either
	// direction. Accepted values are `asc` and `desc`.
	// Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/data-sources/zero_trust_gateway_policy#direction DataCloudflareZeroTrustGatewayPolicy#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Filter the returned rules by one or more `field:value` pairs. Repeat the parameter to combine filters with logical AND.
	//
	// Supported fields are `name`, `id`, `action`, `enabled`, `source_account`,
	// `is_shared`, `filters`, and `expression` (max 1024 bytes). The `source_account`
	// value is matched as a normalized UUID substring. The `filters` value must
	// be one of the rule filter names and matches a member of the rule's `filters`
	// array. The `expression` filter performs a case-insensitive literal
	// substring match across traffic, identity, and device posture expressions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/data-sources/zero_trust_gateway_policy#filter DataCloudflareZeroTrustGatewayPolicy#filter}
	Filter *[]*string `field:"optional" json:"filter" yaml:"filter"`
	// Field to sort the returned rules by. Supported values are `name`, `created_at`, `updated_at`, and `precedence`. Available values: "name", "created_at", "updated_at", "precedence".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/data-sources/zero_trust_gateway_policy#order_by DataCloudflareZeroTrustGatewayPolicy#order_by}
	OrderBy *string `field:"optional" json:"orderBy" yaml:"orderBy"`
	// Case-insensitive substring search across rule name and description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/data-sources/zero_trust_gateway_policy#search DataCloudflareZeroTrustGatewayPolicy#search}
	Search *string `field:"optional" json:"search" yaml:"search"`
}

