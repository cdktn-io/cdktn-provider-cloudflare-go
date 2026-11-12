// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accessrule


type AccessRuleConfiguration struct {
	// For `target: ip`: The configuration target.
	//
	// You must set the target to `ip` when specifying an IP address in the rule.
	//
	// For `target: ip6`: The configuration target. You must set the target to `ip6` when specifying an IPv6 address in the rule.
	//
	// For `target: ip_range`: The configuration target. You must set the target to `ip_range` when specifying an IP address range in the rule.
	//
	// For `target: asn`: The configuration target. You must set the target to `asn` when specifying an Autonomous System Number (ASN) in the rule.
	//
	// For `target: country`: The configuration target. You must set the target to `country` when specifying a country code in the rule.
	// Available values: "ip", "ip6", "ip_range", "asn", "country".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/access_rule#target AccessRule#target}
	Target *string `field:"optional" json:"target" yaml:"target"`
	// For `target: ip`: The IP address to match.
	//
	// This address will be compared to the IP address of incoming requests.
	//
	// For `target: ip6`: The IPv6 address to match.
	//
	// For `target: ip_range`: The IP address range to match. You can only use prefix lengths `/16` and `/24` for IPv4 ranges, and prefix lengths `/32`, `/48`, and `/64` for IPv6 ranges.
	//
	// For `target: asn`: The AS number to match.
	//
	// For `target: country`: The two-letter ISO-3166-1 alpha-2 code to match. For more information, refer to [IP Access rules: Parameters](https://developers.cloudflare.com/waf/tools/ip-access-rules/parameters/#country).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/access_rule#value AccessRule#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

