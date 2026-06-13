// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustdlpcustomprofile


type ZeroTrustDlpCustomProfileContextAwareness struct {
	// If true, scan the context of predefined entries to only return matches surrounded by keywords.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/resources/zero_trust_dlp_custom_profile#enabled ZeroTrustDlpCustomProfile#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Content types to exclude from context analysis and return all matches.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/resources/zero_trust_dlp_custom_profile#skip ZeroTrustDlpCustomProfile#skip}
	Skip *ZeroTrustDlpCustomProfileContextAwarenessSkip `field:"optional" json:"skip" yaml:"skip"`
}

