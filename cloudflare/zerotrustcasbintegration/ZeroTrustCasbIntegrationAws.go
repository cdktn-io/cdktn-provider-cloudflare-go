// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustcasbintegration


type ZeroTrustCasbIntegrationAws struct {
	// Authenticate by delegating to a cross-account IAM role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_casb_integration#aws_iam_role ZeroTrustCasbIntegration#aws_iam_role}
	AwsIamRole *ZeroTrustCasbIntegrationAwsAwsIamRole `field:"optional" json:"awsIamRole" yaml:"awsIamRole"`
}

