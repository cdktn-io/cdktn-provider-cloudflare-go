// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccesspolicy


type ZeroTrustAccessPolicyExcludeCloudflareAccountMember struct {
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/resources/zero_trust_access_policy#account_id ZeroTrustAccessPolicy#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
}

