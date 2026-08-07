// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ruleset


type RulesetRulesActionParametersVaryHeaders struct {
	// How the header value is treated when building the cache key. Available values: "bypass", "passthrough", "normalize".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/resources/ruleset#action Ruleset#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// The set of languages to normalize against. Only valid for the `accept-language` header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/resources/ruleset#languages Ruleset#languages}
	Languages *[]*string `field:"optional" json:"languages" yaml:"languages"`
	// The set of media types to normalize against. Only valid for the `accept` header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/resources/ruleset#media_types Ruleset#media_types}
	MediaTypes *[]*string `field:"optional" json:"mediaTypes" yaml:"mediaTypes"`
}

