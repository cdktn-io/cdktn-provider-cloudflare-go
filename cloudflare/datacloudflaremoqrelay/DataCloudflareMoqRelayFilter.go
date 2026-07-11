// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflaremoqrelay


type DataCloudflareMoqRelayFilter struct {
	// Sort order by `created`. When true, results are returned oldest-first (ascending); otherwise newest-first (descending, the default).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/data-sources/moq_relay#asc DataCloudflareMoqRelay#asc}
	Asc interface{} `field:"optional" json:"asc" yaml:"asc"`
	// Cursor for pagination.
	//
	// Returns relays created strictly after this
	// RFC 3339 timestamp (typically the `created` value of the last item
	// on the current page, to fetch the next page).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/data-sources/moq_relay#created_after DataCloudflareMoqRelay#created_after}
	CreatedAfter *string `field:"optional" json:"createdAfter" yaml:"createdAfter"`
	// Cursor for pagination.
	//
	// Returns relays created strictly before this
	// RFC 3339 timestamp (typically the `created` value of the first item
	// on the current page, to fetch the previous page).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/data-sources/moq_relay#created_before DataCloudflareMoqRelay#created_before}
	CreatedBefore *string `field:"optional" json:"createdBefore" yaml:"createdBefore"`
	// Maximum number of relays to return per page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/data-sources/moq_relay#per_page DataCloudflareMoqRelay#per_page}
	PerPage *float64 `field:"optional" json:"perPage" yaml:"perPage"`
}

