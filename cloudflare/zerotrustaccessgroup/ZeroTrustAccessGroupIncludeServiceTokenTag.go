// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessgroup


type ZeroTrustAccessGroupIncludeServiceTokenTag struct {
	// The tag key to match. Keys are case-sensitive and may contain Unicode letters and digits, underscores, periods, and hyphens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/zero_trust_access_group#key ZeroTrustAccessGroup#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The tag value to match. Values are case-sensitive, preserve whitespace, and must not contain NUL characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/zero_trust_access_group#value ZeroTrustAccessGroup#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

