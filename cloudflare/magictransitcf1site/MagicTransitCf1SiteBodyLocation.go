// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package magictransitcf1site


type MagicTransitCf1SiteBodyLocation struct {
	// Latitude of the CF1 Site.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/resources/magic_transit_cf1_site#lat MagicTransitCf1Site#lat}
	Lat *float64 `field:"optional" json:"lat" yaml:"lat"`
	// Longitude of the CF1 Site.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/resources/magic_transit_cf1_site#long MagicTransitCf1Site#long}
	Long *float64 `field:"optional" json:"long" yaml:"long"`
	// Name of nearest town, city, or village.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/resources/magic_transit_cf1_site#name MagicTransitCf1Site#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

