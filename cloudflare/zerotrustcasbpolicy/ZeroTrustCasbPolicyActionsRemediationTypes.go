// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustcasbpolicy


type ZeroTrustCasbPolicyActionsRemediationTypes struct {
	// The ID of the remediation type to execute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/resources/zero_trust_casb_policy#remediation_type_id ZeroTrustCasbPolicy#remediation_type_id}
	RemediationTypeId *string `field:"required" json:"remediationTypeId" yaml:"remediationTypeId"`
}

