// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustcasbintegration


type ZeroTrustCasbIntegrationOpenaiChatgptStandardApiKey struct {
	// OpenAI Admin API key with api.management.read access. This value is write-only and is never persisted to Terraform state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_casb_integration#admin_api_key ZeroTrustCasbIntegration#admin_api_key}
	AdminApiKey *string `field:"required" json:"adminApiKey" yaml:"adminApiKey"`
	// OpenAI Organization ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_casb_integration#organization_id ZeroTrustCasbIntegration#organization_id}
	OrganizationId *string `field:"required" json:"organizationId" yaml:"organizationId"`
	// OpenAI Project API key, used for DLP. This value is write-only and is never persisted to Terraform state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_casb_integration#project_api_key ZeroTrustCasbIntegration#project_api_key}
	ProjectApiKey *string `field:"optional" json:"projectApiKey" yaml:"projectApiKey"`
	// OpenAI Project ID, used for DLP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_casb_integration#project_id ZeroTrustCasbIntegration#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
}

