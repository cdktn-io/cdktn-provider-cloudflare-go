// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustcasbintegration


type ZeroTrustCasbIntegrationBoxBoxServerAuthentication struct {
	// Box Enterprise ID from Admin Console > Accounts & Billing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_casb_integration#enterprise_id ZeroTrustCasbIntegration#enterprise_id}
	EnterpriseId *string `field:"required" json:"enterpriseId" yaml:"enterpriseId"`
}

