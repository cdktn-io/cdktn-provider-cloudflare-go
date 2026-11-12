// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package magictransitsitelan


type MagicTransitSiteLanRemoteSubnets struct {
	// A valid CIDR notation representing an IP range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/magic_transit_site_lan#local_alias_prefix MagicTransitSiteLan#local_alias_prefix}
	LocalAliasPrefix *string `field:"required" json:"localAliasPrefix" yaml:"localAliasPrefix"`
	// A valid CIDR notation representing an IP range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/magic_transit_site_lan#remote_prefix MagicTransitSiteLan#remote_prefix}
	RemotePrefix *string `field:"required" json:"remotePrefix" yaml:"remotePrefix"`
}

