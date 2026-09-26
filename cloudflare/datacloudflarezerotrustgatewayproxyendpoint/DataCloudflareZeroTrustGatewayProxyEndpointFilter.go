// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezerotrustgatewayproxyendpoint


type DataCloudflareZeroTrustGatewayProxyEndpointFilter struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/data-sources/zero_trust_gateway_proxy_endpoint#direction DataCloudflareZeroTrustGatewayProxyEndpoint#direction}
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/data-sources/zero_trust_gateway_proxy_endpoint#filter DataCloudflareZeroTrustGatewayProxyEndpoint#filter}
	Filter *[]*string `field:"optional" json:"filter" yaml:"filter"`
	// Field to sort the returned endpoints by.
	//
	// When omitted, the order of
	// results is unspecified. Supported values:
	//   * `name` — sort alphabetically by endpoint name.
	//   * `created_at` — sort by creation time; defaults to descending unless `direction` is set.
	//   * `updated_at` — sort by last-modified time; defaults to descending unless `direction` is set.
	// Available values: "name", "created_at", "updated_at".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/data-sources/zero_trust_gateway_proxy_endpoint#order_by DataCloudflareZeroTrustGatewayProxyEndpoint#order_by}
	OrderBy *string `field:"optional" json:"orderBy" yaml:"orderBy"`
	// Case-insensitive substring match on the endpoint name. When combined with `filter`, both must match (logical AND).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/data-sources/zero_trust_gateway_proxy_endpoint#search DataCloudflareZeroTrustGatewayProxyEndpoint#search}
	Search *string `field:"optional" json:"search" yaml:"search"`
}

