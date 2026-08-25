// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ruleset


type RulesetRulesActionParametersVary struct {
	// Controls how response Vary headers without a per-header override contribute to the cache key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/ruleset#default Ruleset#default}
	Default *RulesetRulesActionParametersVaryDefault `field:"required" json:"default" yaml:"default"`
	// A mapping of lowercase request header names to their vary configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/ruleset#headers Ruleset#headers}
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
}

