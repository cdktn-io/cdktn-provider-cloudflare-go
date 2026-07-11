// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustdnslocation


type ZeroTrustDnsLocationMaxTtl struct {
	// `inherit` uses the account `max_ttl_secs`. `override` uses this location's `ttl_secs`. `disabled` leaves returned TTLs unchanged. Available values: "inherit", "override", "disabled".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/resources/zero_trust_dns_location#mode ZeroTrustDnsLocation#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// Location-specific cap on DNS response TTLs, in seconds.
	//
	// Required when `mode` is `override`. Must be omitted when `mode` is `inherit` or `disabled`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/resources/zero_trust_dns_location#ttl_secs ZeroTrustDnsLocation#ttl_secs}
	TtlSecs *float64 `field:"optional" json:"ttlSecs" yaml:"ttlSecs"`
}

