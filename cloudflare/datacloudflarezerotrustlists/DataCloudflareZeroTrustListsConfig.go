// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezerotrustlists

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataCloudflareZeroTrustListsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/zero_trust_lists#account_id DataCloudflareZeroTrustLists#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Sort direction.
	//
	// Applies to the field named in `order_by`; when `order_by`
	// is omitted it applies to the default `created_at` ordering. When
	// `direction` is omitted the default is field-specific: explicitly choosing
	// `created_at` or `updated_at` defaults to descending (newest first); `name`
	// and `item_count` default to ascending; and the default `created_at`
	// ordering used when `order_by` is omitted is ascending (for backwards
	// compatibility).
	//   * `asc` — ascending.
	//   * `desc` — descending.
	// Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/zero_trust_lists#direction DataCloudflareZeroTrustLists#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Filter the returned lists by one or more `field:value` pairs.
	//
	// Repeat the parameter to apply multiple filters; they are combined with
	// logical AND (a list must satisfy every filter to be returned).
	//
	// Supported fields and their matching behaviour:
	//   * `name` — case-insensitive substring match on the list name.
	//   * `id` — substring match on the list ID (UUID), with or without dashes.
	//   * `type` — exact match on the list type. Supersedes the legacy `type` query
	//     parameter when both are supplied. Must be one of the valid type values.
	//   * `item_count` — exact integer match on the number of items in the list.
	//
	// Each entry must match one of the per-field patterns below: the field must be
	// one of `name`, `id`, `type`, or `item_count`; `name`/`id` accept any value,
	// `type` is restricted to the valid list type values, and `item_count` must be
	// a non-negative integer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/zero_trust_lists#filter DataCloudflareZeroTrustLists#filter}
	Filter *[]*string `field:"optional" json:"filter" yaml:"filter"`
	// Max items to fetch, default: 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/zero_trust_lists#max_items DataCloudflareZeroTrustLists#max_items}
	MaxItems *float64 `field:"optional" json:"maxItems" yaml:"maxItems"`
	// Field to sort the returned lists by.
	//
	// When omitted, results are ordered by
	// `created_at` in ascending order (i.e. creation order) for backwards
	// compatibility. Supported values:
	//   * `name` — sort alphabetically by list name.
	//   * `created_at` — sort by creation time; defaults to descending unless `direction` is set.
	//   * `updated_at` — sort by last-modified time; defaults to descending unless `direction` is set.
	//   * `item_count` — sort by number of items in the list.
	// Available values: "name", "created_at", "updated_at", "item_count".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/zero_trust_lists#order_by DataCloudflareZeroTrustLists#order_by}
	OrderBy *string `field:"optional" json:"orderBy" yaml:"orderBy"`
	// Case-insensitive substring match on the list name or description. When combined with `filter`, both must match (logical AND).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/zero_trust_lists#search DataCloudflareZeroTrustLists#search}
	Search *string `field:"optional" json:"search" yaml:"search"`
	// Specify the list type. Available values: "SERIAL", "URL", "DOMAIN", "EMAIL", "IP", "CATEGORY", "LOCATION", "DEVICE", "AAGUID".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/zero_trust_lists#type DataCloudflareZeroTrustLists#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

