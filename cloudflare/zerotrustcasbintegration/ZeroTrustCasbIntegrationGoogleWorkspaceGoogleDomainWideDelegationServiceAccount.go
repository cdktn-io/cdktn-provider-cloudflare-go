// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustcasbintegration


type ZeroTrustCasbIntegrationGoogleWorkspaceGoogleDomainWideDelegationServiceAccount struct {
	// A Google Workspace super administrator email address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_casb_integration#administrator_email ZeroTrustCasbIntegration#administrator_email}
	AdministratorEmail *string `field:"required" json:"administratorEmail" yaml:"administratorEmail"`
	// Contents of a Google service account JSON key file.
	//
	// This value is write-only and is never persisted to Terraform state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_casb_integration#service_account_key_json ZeroTrustCasbIntegration#service_account_key_json}
	ServiceAccountKeyJson *string `field:"required" json:"serviceAccountKeyJson" yaml:"serviceAccountKeyJson"`
}

