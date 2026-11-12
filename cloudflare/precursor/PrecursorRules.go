// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package precursor


type PrecursorRules struct {
	// The filter expression that determines which requests the rule matches.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/precursor#expression Precursor#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Whether Precursor shows an interstitial when it takes action on a request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/precursor#show_interstitial Precursor#show_interstitial}
	ShowInterstitial interface{} `field:"required" json:"showInterstitial" yaml:"showInterstitial"`
	// An informative description of the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/precursor#description Precursor#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether the rule is active.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/precursor#enabled Precursor#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

