// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccesspolicy


type ZeroTrustAccessPolicyConnectionRules struct {
	// The RDP-specific rules that define clipboard behavior for RDP connections.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/resources/zero_trust_access_policy#rdp ZeroTrustAccessPolicy#rdp}
	Rdp *ZeroTrustAccessPolicyConnectionRulesRdp `field:"optional" json:"rdp" yaml:"rdp"`
}

