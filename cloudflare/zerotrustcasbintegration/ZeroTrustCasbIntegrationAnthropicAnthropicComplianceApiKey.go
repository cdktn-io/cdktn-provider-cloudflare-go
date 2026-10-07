// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustcasbintegration


type ZeroTrustCasbIntegrationAnthropicAnthropicComplianceApiKey struct {
	// Anthropic Compliance API key. This value is write-only and is never persisted to Terraform state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_casb_integration#compliance_api_key ZeroTrustCasbIntegration#compliance_api_key}
	ComplianceApiKey *string `field:"required" json:"complianceApiKey" yaml:"complianceApiKey"`
	// Organization ID. Auto-extracted from the key if not provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_casb_integration#tenant_id ZeroTrustCasbIntegration#tenant_id}
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
}

