// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package precursor


type PrecursorEnforcementRules struct {
	// The filter expression that determines which requests the rule matches.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/precursor#expression Precursor#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// The override mode Precursor applies to requests matching an enforcement rule. Unlike `default_mode`, this cannot be `off`. Available values: "min-friction", "max-security".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/precursor#mode Precursor#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// An informative description of the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/precursor#description Precursor#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether the rule is active.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/precursor#enabled Precursor#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

