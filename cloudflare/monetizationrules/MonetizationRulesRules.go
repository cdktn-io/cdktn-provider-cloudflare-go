// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package monetizationrules


type MonetizationRulesRules struct {
	// 0x-prefixed 20-byte hexadecimal Ethereum address.
	//
	// Mixed-case addresses must carry a valid EIP-55 checksum; all-lowercase or all-uppercase addresses are also accepted. The address must pass wallet screening whenever its rule is deployed or patched.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.28.1/docs/resources/monetization_rules#address MonetizationRules#address}
	Address *string `field:"required" json:"address" yaml:"address"`
	// Wirefilter expression identifying the requests that require payment. Forwarded verbatim to the Rulesets API, which validates its syntax.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.28.1/docs/resources/monetization_rules#expression MonetizationRules#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// X402 payment scheme.
	//
	// "exact" requires the specified payment amount; "upto" permits a payment up to the specified amount. Both fixed-price schemes require the price field.
	// Available values: "exact", "upto", "origin_controlled".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.28.1/docs/resources/monetization_rules#scheme MonetizationRules#scheme}
	Scheme *string `field:"required" json:"scheme" yaml:"scheme"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.28.1/docs/resources/monetization_rules#description MonetizationRules#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.28.1/docs/resources/monetization_rules#enabled MonetizationRules#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Optional on input.
	//
	// Include the ID of an existing payment rule to update it in place, preserving its stable identity across the full-ruleset replacement; omit it to create a new rule. An ID that does not match an existing payment rule in the zone is rejected. Rules omitted from the request are deleted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.28.1/docs/resources/monetization_rules#id MonetizationRules#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Price in the smallest indivisible unit of the configured payment token, encoded as a decimal string.
	//
	// Must be a canonical decimal integer in [1000, 100000000]: no leading zeros, and a bare JSON number is rejected. The price is required for the fixed-price schemes ("exact" and "upto") and must be at least 1000, the smallest amount the payment facilitator can settle ($0.001 for a 6-decimal token such as USDC), and at most 100000000 ($100 for a 6-decimal token). When the scheme is "origin_controlled" the origin server sets pricing dynamically and the rule carries no price: the field must be omitted — any provided value, including "0", is rejected because it would not be enforced. Responses always encode this as a string and omit it for "origin_controlled" rules; the pattern matches exactly the set of values the server accepts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.28.1/docs/resources/monetization_rules#price MonetizationRules#price}
	Price *string `field:"optional" json:"price" yaml:"price"`
}

