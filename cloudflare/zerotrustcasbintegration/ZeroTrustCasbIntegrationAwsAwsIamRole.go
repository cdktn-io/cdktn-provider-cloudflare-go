// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustcasbintegration


type ZeroTrustCasbIntegrationAwsAwsIamRole struct {
	// External ID required when assuming the IAM role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_casb_integration#external_id ZeroTrustCasbIntegration#external_id}
	ExternalId *string `field:"required" json:"externalId" yaml:"externalId"`
	// ARN of the cross-account IAM role Cloudflare will assume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_casb_integration#role_arn ZeroTrustCasbIntegration#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

