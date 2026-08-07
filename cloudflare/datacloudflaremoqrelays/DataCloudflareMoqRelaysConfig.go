// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflaremoqrelays

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataCloudflareMoqRelaysConfig struct {
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
	// Cloudflare account identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/data-sources/moq_relays#account_id DataCloudflareMoqRelays#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Sort order by `created`. When true, results are returned oldest-first (ascending); otherwise newest-first (descending, the default).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/data-sources/moq_relays#asc DataCloudflareMoqRelays#asc}
	Asc interface{} `field:"optional" json:"asc" yaml:"asc"`
	// Cursor for pagination.
	//
	// Returns relays created strictly after this
	// RFC 3339 timestamp (typically the `created` value of the last item
	// on the current page, to fetch the next page).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/data-sources/moq_relays#created_after DataCloudflareMoqRelays#created_after}
	CreatedAfter *string `field:"optional" json:"createdAfter" yaml:"createdAfter"`
	// Cursor for pagination.
	//
	// Returns relays created strictly before this
	// RFC 3339 timestamp (typically the `created` value of the first item
	// on the current page, to fetch the previous page).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/data-sources/moq_relays#created_before DataCloudflareMoqRelays#created_before}
	CreatedBefore *string `field:"optional" json:"createdBefore" yaml:"createdBefore"`
	// Max items to fetch, default: 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/data-sources/moq_relays#max_items DataCloudflareMoqRelays#max_items}
	MaxItems *float64 `field:"optional" json:"maxItems" yaml:"maxItems"`
	// Maximum number of relays to return per page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/data-sources/moq_relays#per_page DataCloudflareMoqRelays#per_page}
	PerPage *float64 `field:"optional" json:"perPage" yaml:"perPage"`
}

