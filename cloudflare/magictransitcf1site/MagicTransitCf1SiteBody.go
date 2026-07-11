// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package magictransitcf1site


type MagicTransitCf1SiteBody struct {
	// A human-provided name describing the CF1 Site that should be unique within the account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/resources/magic_transit_cf1_site#name MagicTransitCf1Site#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A human-provided description of the CF1 Site.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/resources/magic_transit_cf1_site#description MagicTransitCf1Site#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/resources/magic_transit_cf1_site#location MagicTransitCf1Site#location}.
	Location *MagicTransitCf1SiteBodyLocation `field:"optional" json:"location" yaml:"location"`
}

