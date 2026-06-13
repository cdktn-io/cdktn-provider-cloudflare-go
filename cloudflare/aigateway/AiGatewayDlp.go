// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aigateway


type AiGatewayDlp struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/resources/ai_gateway#enabled AiGateway#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Available values: "BLOCK", "FLAG".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/resources/ai_gateway#action AiGateway#action}
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/resources/ai_gateway#policies AiGateway#policies}.
	Policies interface{} `field:"optional" json:"policies" yaml:"policies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/resources/ai_gateway#profiles AiGateway#profiles}.
	Profiles *[]*string `field:"optional" json:"profiles" yaml:"profiles"`
}

