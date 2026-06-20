// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aigateway


type AiGatewaySpendLimitsRulesMetadata struct {
	// Available values: "partition", "filter".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.0/docs/resources/ai_gateway#mode AiGateway#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.0/docs/resources/ai_gateway#values AiGateway#values}.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

