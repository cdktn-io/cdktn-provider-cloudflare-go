// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustcasbintegration


type ZeroTrustCasbIntegrationOpenai struct {
	// Authenticate with an OpenAI Compliance API key. Requires an Enterprise plan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_casb_integration#chatgpt_compliance_api_key ZeroTrustCasbIntegration#chatgpt_compliance_api_key}
	ChatgptComplianceApiKey *ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKey `field:"optional" json:"chatgptComplianceApiKey" yaml:"chatgptComplianceApiKey"`
	// Authenticate with an OpenAI Admin API key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_casb_integration#chatgpt_standard_api_key ZeroTrustCasbIntegration#chatgpt_standard_api_key}
	ChatgptStandardApiKey *ZeroTrustCasbIntegrationOpenaiChatgptStandardApiKey `field:"optional" json:"chatgptStandardApiKey" yaml:"chatgptStandardApiKey"`
}

