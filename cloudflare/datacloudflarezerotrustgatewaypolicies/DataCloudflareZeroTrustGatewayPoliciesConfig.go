// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezerotrustgatewaypolicies

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataCloudflareZeroTrustGatewayPoliciesConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Specify the Cloudflare account identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/zero_trust_gateway_policies#account_id DataCloudflareZeroTrustGatewayPolicies#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Sort direction.
	//
	// When `order_by` is omitted, this controls the direction
	// of the existing precedence ordering. Shared rules remain first in either
	// direction. Accepted values are `asc` and `desc`.
	// Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/zero_trust_gateway_policies#direction DataCloudflareZeroTrustGatewayPolicies#direction}
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/zero_trust_gateway_policies#filter DataCloudflareZeroTrustGatewayPolicies#filter}
	Filter *[]*string `field:"optional" json:"filter" yaml:"filter"`
	// Max items to fetch, default: 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/zero_trust_gateway_policies#max_items DataCloudflareZeroTrustGatewayPolicies#max_items}
	MaxItems *float64 `field:"optional" json:"maxItems" yaml:"maxItems"`
	// Field to sort the returned rules by. Supported values are `name`, `created_at`, `updated_at`, and `precedence`. Available values: "name", "created_at", "updated_at", "precedence".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/zero_trust_gateway_policies#order_by DataCloudflareZeroTrustGatewayPolicies#order_by}
	OrderBy *string `field:"optional" json:"orderBy" yaml:"orderBy"`
	// Case-insensitive substring search across rule name and description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/zero_trust_gateway_policies#search DataCloudflareZeroTrustGatewayPolicies#search}
	Search *string `field:"optional" json:"search" yaml:"search"`
}

