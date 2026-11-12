// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustcwspolicyset


type ZeroTrustCwsPolicySetPolicies struct {
	// ID of the member policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/zero_trust_cws_policy_set#policy_id ZeroTrustCwsPolicySet#policy_id}
	PolicyId *string `field:"required" json:"policyId" yaml:"policyId"`
	// Whether this policy should enforce when it matches. Defaults to true; false keeps the policy attached for non-enforcing evaluation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/zero_trust_cws_policy_set#enforce ZeroTrustCwsPolicySet#enforce}
	Enforce interface{} `field:"optional" json:"enforce" yaml:"enforce"`
}

