// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezerotrustcasbfindingtype


type DataCloudflareZeroTrustCasbFindingTypeFilter struct {
	// Sort direction for the primary field.
	//
	// Defaults to 'desc'.
	// Ties are resolved by case-insensitive name, then ID, both ascending regardless of direction.
	// Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/data-sources/zero_trust_casb_finding_type#direction DataCloudflareZeroTrustCasbFindingType#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Return only finding types updated at or before this RFC 3339 timestamp (inclusive).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/data-sources/zero_trust_casb_finding_type#max_updated_at DataCloudflareZeroTrustCasbFindingType#max_updated_at}
	MaxUpdatedAt *string `field:"optional" json:"maxUpdatedAt" yaml:"maxUpdatedAt"`
	// Return only finding types updated at or after this RFC 3339 timestamp (inclusive).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/data-sources/zero_trust_casb_finding_type#min_updated_at DataCloudflareZeroTrustCasbFindingType#min_updated_at}
	MinUpdatedAt *string `field:"optional" json:"minUpdatedAt" yaml:"minUpdatedAt"`
	// Field to sort by (case-insensitive).
	//
	// Defaults to 'created'.
	// Name and vendor sorting is case-insensitive; vendor uses friendly provider names.
	// Ascending order: type Content then Posture; origin custom then system;
	// severity Low, Medium, High, Critical.
	// Available values: "created", "updated", "match_count", "name", "vendor", "type", "origin", "severity".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/data-sources/zero_trust_casb_finding_type#order DataCloudflareZeroTrustCasbFindingType#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
	// Filter finding types by origin. Supports multiple comma-separated, case-insensitive values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/data-sources/zero_trust_casb_finding_type#origins DataCloudflareZeroTrustCasbFindingType#origins}
	Origins *[]*string `field:"optional" json:"origins" yaml:"origins"`
	// When true, return only quarantined finding types; when false, return only non-quarantined finding types. Omit to return both.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/data-sources/zero_trust_casb_finding_type#quarantined DataCloudflareZeroTrustCasbFindingType#quarantined}
	Quarantined interface{} `field:"optional" json:"quarantined" yaml:"quarantined"`
	// Filter finding types by name or ID (case-insensitive substring match).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/data-sources/zero_trust_casb_finding_type#search DataCloudflareZeroTrustCasbFindingType#search}
	Search *string `field:"optional" json:"search" yaml:"search"`
	// Filter finding types by severity. Supports multiple comma-separated, case-insensitive values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/data-sources/zero_trust_casb_finding_type#severities DataCloudflareZeroTrustCasbFindingType#severities}
	Severities *[]*string `field:"optional" json:"severities" yaml:"severities"`
	// Filter finding types by category type. Supports multiple comma-separated, case-insensitive values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/data-sources/zero_trust_casb_finding_type#types DataCloudflareZeroTrustCasbFindingType#types}
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
	// Filter finding types by vendor.
	//
	// Supports multiple comma-separated values. Each value may be either the vendor enum (GOOGLE_WORKSPACE) or the friendly name returned by the read routes (Google Workspace).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/data-sources/zero_trust_casb_finding_type#vendors DataCloudflareZeroTrustCasbFindingType#vendors}
	Vendors *[]*string `field:"optional" json:"vendors" yaml:"vendors"`
}

