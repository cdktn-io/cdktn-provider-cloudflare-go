// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustcasbintegration


type ZeroTrustCasbIntegrationAnthropic struct {
	// Authenticate with an Anthropic Admin API key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_casb_integration#anthropic_admin_api_key ZeroTrustCasbIntegration#anthropic_admin_api_key}
	AnthropicAdminApiKey *ZeroTrustCasbIntegrationAnthropicAnthropicAdminApiKey `field:"optional" json:"anthropicAdminApiKey" yaml:"anthropicAdminApiKey"`
	// Authenticate with an Anthropic Compliance API key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_casb_integration#anthropic_compliance_api_key ZeroTrustCasbIntegration#anthropic_compliance_api_key}
	AnthropicComplianceApiKey *ZeroTrustCasbIntegrationAnthropicAnthropicComplianceApiKey `field:"optional" json:"anthropicComplianceApiKey" yaml:"anthropicComplianceApiKey"`
	// Authenticate with an Anthropic Workspace API key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_casb_integration#anthropic_workspace_api_key ZeroTrustCasbIntegration#anthropic_workspace_api_key}
	AnthropicWorkspaceApiKey *ZeroTrustCasbIntegrationAnthropicAnthropicWorkspaceApiKey `field:"optional" json:"anthropicWorkspaceApiKey" yaml:"anthropicWorkspaceApiKey"`
}

