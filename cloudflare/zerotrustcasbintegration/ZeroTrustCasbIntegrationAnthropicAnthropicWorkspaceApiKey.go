// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustcasbintegration


type ZeroTrustCasbIntegrationAnthropicAnthropicWorkspaceApiKey struct {
	// Anthropic Workspace API key. This value is write-only and is never persisted to Terraform state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_casb_integration#api_key ZeroTrustCasbIntegration#api_key}
	ApiKey *string `field:"required" json:"apiKey" yaml:"apiKey"`
	// Workspace ID, found in the Anthropic Console URL after /workspaces/.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_casb_integration#tenant_id ZeroTrustCasbIntegration#tenant_id}
	TenantId *string `field:"required" json:"tenantId" yaml:"tenantId"`
}

