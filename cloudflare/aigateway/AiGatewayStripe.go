// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aigateway


type AiGatewayStripe struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/resources/ai_gateway#authorization AiGateway#authorization}.
	Authorization *string `field:"required" json:"authorization" yaml:"authorization"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/resources/ai_gateway#usage_events AiGateway#usage_events}.
	UsageEvents interface{} `field:"required" json:"usageEvents" yaml:"usageEvents"`
}

