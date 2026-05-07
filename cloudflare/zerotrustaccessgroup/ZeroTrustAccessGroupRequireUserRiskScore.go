// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessgroup


type ZeroTrustAccessGroupRequireUserRiskScore struct {
	// A list of risk score levels to match. Values can be low, medium, high, or unscored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.1/docs/resources/zero_trust_access_group#user_risk_score ZeroTrustAccessGroup#user_risk_score}
	UserRiskScore *[]*string `field:"required" json:"userRiskScore" yaml:"userRiskScore"`
}

