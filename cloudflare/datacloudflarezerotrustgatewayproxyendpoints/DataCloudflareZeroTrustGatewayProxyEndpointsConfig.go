// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezerotrustgatewayproxyendpoints

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataCloudflareZeroTrustGatewayProxyEndpointsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/data-sources/zero_trust_gateway_proxy_endpoints#account_id DataCloudflareZeroTrustGatewayProxyEndpoints#account_id}.
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Sort direction.
	//
	// Only takes effect when `order_by` is also provided; it
	// is ignored otherwise. When `direction` is omitted the effective
	// direction is field-specific: `created_at` and `updated_at` default to
	// descending (newest first); `name` defaults to ascending.
	//   * `asc` — ascending.
	//   * `desc` — descending.
	// Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/data-sources/zero_trust_gateway_proxy_endpoints#direction DataCloudflareZeroTrustGatewayProxyEndpoints#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Filter the returned proxy endpoints by one or more `field:value` pairs.
	//
	// Repeat the parameter to apply multiple filters; they are combined with
	// logical AND (an endpoint must satisfy every filter to be returned).
	//
	// Supported fields and their matching behaviour:
	//   * `name` — case-insensitive substring match on the endpoint name.
	//   * `id` — substring match on the endpoint ID (UUID), with or without dashes.
	//   * `kind` — exact match on the endpoint kind. The value must be `ip` or `identity`; any other value returns `400`.
	//
	// Each entry must match one of the per-field patterns below: the field
	// must be one of `name`, `id`, or `kind`; `name`/`id` accept any value,
	// while `kind` only accepts `ip` or `identity`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/data-sources/zero_trust_gateway_proxy_endpoints#filter DataCloudflareZeroTrustGatewayProxyEndpoints#filter}
	Filter *[]*string `field:"optional" json:"filter" yaml:"filter"`
	// Max items to fetch, default: 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/data-sources/zero_trust_gateway_proxy_endpoints#max_items DataCloudflareZeroTrustGatewayProxyEndpoints#max_items}
	MaxItems *float64 `field:"optional" json:"maxItems" yaml:"maxItems"`
	// Field to sort the returned endpoints by.
	//
	// When omitted, the order of
	// results is unspecified. Supported values:
	//   * `name` — sort alphabetically by endpoint name.
	//   * `created_at` — sort by creation time; defaults to descending unless `direction` is set.
	//   * `updated_at` — sort by last-modified time; defaults to descending unless `direction` is set.
	// Available values: "name", "created_at", "updated_at".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/data-sources/zero_trust_gateway_proxy_endpoints#order_by DataCloudflareZeroTrustGatewayProxyEndpoints#order_by}
	OrderBy *string `field:"optional" json:"orderBy" yaml:"orderBy"`
	// Case-insensitive substring match on the endpoint name. When combined with `filter`, both must match (logical AND).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/data-sources/zero_trust_gateway_proxy_endpoints#search DataCloudflareZeroTrustGatewayProxyEndpoints#search}
	Search *string `field:"optional" json:"search" yaml:"search"`
}

