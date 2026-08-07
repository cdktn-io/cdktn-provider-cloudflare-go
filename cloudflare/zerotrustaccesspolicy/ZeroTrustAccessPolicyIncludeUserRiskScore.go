// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccesspolicy


type ZeroTrustAccessPolicyIncludeUserRiskScore struct {
	// A list of risk score levels to match. Values can be low, medium, high, or unscored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/resources/zero_trust_access_policy#user_risk_score ZeroTrustAccessPolicy#user_risk_score}
	UserRiskScore *[]*string `field:"required" json:"userRiskScore" yaml:"userRiskScore"`
}

