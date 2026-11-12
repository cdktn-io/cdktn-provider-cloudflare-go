// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezerotrustcwsbasicauthcredential


type DataCloudflareZeroTrustCwsBasicAuthCredentialFilter struct {
	// Filter to credentials attached to this deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/data-sources/zero_trust_cws_basic_auth_credential#deployment_id DataCloudflareZeroTrustCwsBasicAuthCredential#deployment_id}
	DeploymentId *string `field:"optional" json:"deploymentId" yaml:"deploymentId"`
	// Sort direction (default: asc) Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/data-sources/zero_trust_cws_basic_auth_credential#direction DataCloudflareZeroTrustCwsBasicAuthCredential#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Field to order results by (default: created_at) Available values: "username", "created_at", "updated_at".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/data-sources/zero_trust_cws_basic_auth_credential#order DataCloudflareZeroTrustCwsBasicAuthCredential#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
	// Filter by username (case-insensitive substring match).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/data-sources/zero_trust_cws_basic_auth_credential#username DataCloudflareZeroTrustCwsBasicAuthCredential#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

