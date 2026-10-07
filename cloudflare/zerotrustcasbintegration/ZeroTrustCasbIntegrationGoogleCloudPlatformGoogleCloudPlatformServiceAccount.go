// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustcasbintegration


type ZeroTrustCasbIntegrationGoogleCloudPlatformGoogleCloudPlatformServiceAccount struct {
	// Contents of a Google service account JSON key file.
	//
	// This value is write-only and is never persisted to Terraform state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_casb_integration#service_account_key_json ZeroTrustCasbIntegration#service_account_key_json}
	ServiceAccountKeyJson *string `field:"required" json:"serviceAccountKeyJson" yaml:"serviceAccountKeyJson"`
}

