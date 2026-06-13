// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aigateway


type AiGatewayGuardrails struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/resources/ai_gateway#prompt AiGateway#prompt}.
	Prompt *AiGatewayGuardrailsPrompt `field:"required" json:"prompt" yaml:"prompt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/resources/ai_gateway#response AiGateway#response}.
	Response *AiGatewayGuardrailsResponse `field:"required" json:"response" yaml:"response"`
}

