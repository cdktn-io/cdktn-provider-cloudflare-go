// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessapplication


type ZeroTrustAccessApplicationTargetCriteria struct {
	// The port that the targets use for the chosen communication protocol. A port cannot be assigned to multiple protocols.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.28.1/docs/resources/zero_trust_access_application#port ZeroTrustAccessApplication#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The communication protocol your application secures. Available values: "SSH", "RDP".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.28.1/docs/resources/zero_trust_access_application#protocol ZeroTrustAccessApplication#protocol}
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Target is excluded when any selector in this rule matches.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.28.1/docs/resources/zero_trust_access_application#exclude ZeroTrustAccessApplication#exclude}
	Exclude *ZeroTrustAccessApplicationTargetCriteriaExclude `field:"optional" json:"exclude" yaml:"exclude"`
	// Target matches when any selector in this rule matches.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.28.1/docs/resources/zero_trust_access_application#include ZeroTrustAccessApplication#include}
	Include *ZeroTrustAccessApplicationTargetCriteriaInclude `field:"optional" json:"include" yaml:"include"`
	// Target matches only when every selector in this rule matches.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.28.1/docs/resources/zero_trust_access_application#require ZeroTrustAccessApplication#require}
	Require *ZeroTrustAccessApplicationTargetCriteriaRequire `field:"optional" json:"require" yaml:"require"`
	// Deprecated: use include.target_attributes for hostname selectors and rule tags for tag-based matching. New or changed maps may contain only hostname. The API accepts unchanged legacy maps on update, including non-hostname keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.28.1/docs/resources/zero_trust_access_application#target_attributes ZeroTrustAccessApplication#target_attributes}
	TargetAttributes interface{} `field:"optional" json:"targetAttributes" yaml:"targetAttributes"`
}

