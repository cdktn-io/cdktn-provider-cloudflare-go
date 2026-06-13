// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package flagshipflag


type FlagshipFlagRules struct {
	// Conditions the context must satisfy for this rule to match. An empty array matches all contexts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/resources/flagship_flag#conditions FlagshipFlag#conditions}
	Conditions interface{} `field:"required" json:"conditions" yaml:"conditions"`
	// Evaluation order; lower numbers are evaluated first. Must be unique across the flag's rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/resources/flagship_flag#priority FlagshipFlag#priority}
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Variation served when this rule matches. Must be a key in `variations`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/resources/flagship_flag#serve_variation FlagshipFlag#serve_variation}
	ServeVariation *string `field:"required" json:"serveVariation" yaml:"serveVariation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/resources/flagship_flag#rollout FlagshipFlag#rollout}.
	Rollout *FlagshipFlagRulesRollout `field:"optional" json:"rollout" yaml:"rollout"`
}

