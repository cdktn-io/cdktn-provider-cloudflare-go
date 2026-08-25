// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aigateway


type AiGatewaySpendLimitsRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/ai_gateway#limit AiGateway#limit}.
	Limit *float64 `field:"required" json:"limit" yaml:"limit"`
	// Available values: "cost".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/ai_gateway#limit_type AiGateway#limit_type}
	LimitType *string `field:"required" json:"limitType" yaml:"limitType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/ai_gateway#window AiGateway#window}.
	Window *float64 `field:"required" json:"window" yaml:"window"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/ai_gateway#ai_gateway_provider AiGateway#ai_gateway_provider}.
	AiGatewayProvider *AiGatewaySpendLimitsRulesAiGatewayProvider `field:"optional" json:"aiGatewayProvider" yaml:"aiGatewayProvider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/ai_gateway#enabled AiGateway#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/ai_gateway#id AiGateway#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/ai_gateway#metadata AiGateway#metadata}.
	Metadata interface{} `field:"optional" json:"metadata" yaml:"metadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/ai_gateway#model AiGateway#model}.
	Model *AiGatewaySpendLimitsRulesModel `field:"optional" json:"model" yaml:"model"`
	// Available values: "fixed", "sliding".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/ai_gateway#technique AiGateway#technique}
	Technique *string `field:"optional" json:"technique" yaml:"technique"`
}

