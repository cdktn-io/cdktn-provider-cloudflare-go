// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ruleset


type RulesetRulesActionParametersImmutable struct {
	// The operation to perform. Available values: "set", "remove".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.18.0/docs/resources/ruleset#operation Ruleset#operation}
	Operation *string `field:"required" json:"operation" yaml:"operation"`
	// Whether to apply the directive only to Cloudflare's cache.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.18.0/docs/resources/ruleset#cloudflare_only Ruleset#cloudflare_only}
	CloudflareOnly interface{} `field:"optional" json:"cloudflareOnly" yaml:"cloudflareOnly"`
}

