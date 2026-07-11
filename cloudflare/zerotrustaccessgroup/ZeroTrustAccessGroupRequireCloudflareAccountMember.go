// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessgroup


type ZeroTrustAccessGroupRequireCloudflareAccountMember struct {
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/resources/zero_trust_access_group#account_id ZeroTrustAccessGroup#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
}

