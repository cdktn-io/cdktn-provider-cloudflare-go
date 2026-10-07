// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezerotrustdnslocations

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataCloudflareZeroTrustDnsLocationsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/zero_trust_dns_locations#account_id DataCloudflareZeroTrustDnsLocations#account_id}
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/zero_trust_dns_locations#direction DataCloudflareZeroTrustDnsLocations#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Filter the returned locations by one or more `field:value` pairs.
	//
	// Repeat the parameter to apply multiple filters; they are combined with
	// logical AND (a location must satisfy every filter to be returned).
	//
	// Supported fields and their matching behaviour:
	//   * `name` — case-insensitive substring match on the location name.
	//   * `id` — substring match on the location ID (UUID), with or without dashes.
	//   * `is_default` — whether it is the default for the account.
	//
	// Each entry must match one of the per-field patterns below:
	//   * the field must be one of `name`, `id`, or `is_default`;
	//   * `name`/`id` accept any value;
	//   * `is_default` only accepts `true` or `false`; any other value returns `400`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/zero_trust_dns_locations#filter DataCloudflareZeroTrustDnsLocations#filter}
	Filter *[]*string `field:"optional" json:"filter" yaml:"filter"`
	// Max items to fetch, default: 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/zero_trust_dns_locations#max_items DataCloudflareZeroTrustDnsLocations#max_items}
	MaxItems *float64 `field:"optional" json:"maxItems" yaml:"maxItems"`
	// Field to sort the returned locations by.
	//
	// When omitted, the order of
	// results is unspecified. Supported values:
	//   * `name` — sort alphabetically by location name.
	//   * `created_at` — sort by creation time; defaults to descending unless `direction` is set.
	//   * `updated_at` — sort by last-modified time; defaults to descending unless `direction` is set.
	// Available values: "name", "created_at", "updated_at".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/zero_trust_dns_locations#order_by DataCloudflareZeroTrustDnsLocations#order_by}
	OrderBy *string `field:"optional" json:"orderBy" yaml:"orderBy"`
	// Case-insensitive substring match on the location name. When combined with `filter`, both must match (logical AND).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/zero_trust_dns_locations#search DataCloudflareZeroTrustDnsLocations#search}
	Search *string `field:"optional" json:"search" yaml:"search"`
}

