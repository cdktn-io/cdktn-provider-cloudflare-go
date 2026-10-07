// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustcasbintegration


type ZeroTrustCasbIntegrationGoogleCloudPlatform struct {
	// Authenticate with a service account key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_casb_integration#google_cloud_platform_service_account ZeroTrustCasbIntegration#google_cloud_platform_service_account}
	GoogleCloudPlatformServiceAccount *ZeroTrustCasbIntegrationGoogleCloudPlatformGoogleCloudPlatformServiceAccount `field:"optional" json:"googleCloudPlatformServiceAccount" yaml:"googleCloudPlatformServiceAccount"`
}

