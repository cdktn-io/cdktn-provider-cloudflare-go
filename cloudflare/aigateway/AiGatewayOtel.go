// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aigateway


type AiGatewayOtel struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/resources/ai_gateway#headers AiGateway#headers}.
	Headers *map[string]*string `field:"required" json:"headers" yaml:"headers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/resources/ai_gateway#url AiGateway#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/resources/ai_gateway#authorization AiGateway#authorization}.
	Authorization *string `field:"optional" json:"authorization" yaml:"authorization"`
	// Available values: "json", "protobuf".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/resources/ai_gateway#content_type AiGateway#content_type}
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
}

