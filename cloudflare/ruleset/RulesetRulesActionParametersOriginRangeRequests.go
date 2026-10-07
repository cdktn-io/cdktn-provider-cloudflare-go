// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ruleset


type RulesetRulesActionParametersOriginRangeRequests struct {
	// Whether to use range requests. `default` is the behaviour the zone gets without this rule. Available values: "on", "off", "default".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/ruleset#mode Ruleset#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
}

