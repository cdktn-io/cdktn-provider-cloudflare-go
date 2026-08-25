// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ruleset


type RulesetRulesActionParametersVaryDefault struct {
	// How the header value is treated when building the cache key. Available values: "bypass", "passthrough", "normalize".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/ruleset#action Ruleset#action}
	Action *string `field:"required" json:"action" yaml:"action"`
}

