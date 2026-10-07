// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustcasbintegration


type ZeroTrustCasbIntegrationGoogleWorkspace struct {
	// Authenticate with a service account granted domain-wide delegation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_casb_integration#google_domain_wide_delegation_service_account ZeroTrustCasbIntegration#google_domain_wide_delegation_service_account}
	GoogleDomainWideDelegationServiceAccount *ZeroTrustCasbIntegrationGoogleWorkspaceGoogleDomainWideDelegationServiceAccount `field:"optional" json:"googleDomainWideDelegationServiceAccount" yaml:"googleDomainWideDelegationServiceAccount"`
}

