// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustcasbpolicy


type ZeroTrustCasbPolicyActions struct {
	// Remediation actions to execute (at most one).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_casb_policy#remediation_types ZeroTrustCasbPolicy#remediation_types}
	RemediationTypes interface{} `field:"optional" json:"remediationTypes" yaml:"remediationTypes"`
	// Webhook actions to execute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_casb_policy#webhook_configs ZeroTrustCasbPolicy#webhook_configs}
	WebhookConfigs interface{} `field:"optional" json:"webhookConfigs" yaml:"webhookConfigs"`
}

