// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustcasbpolicy


type ZeroTrustCasbPolicyActionsWebhookConfigs struct {
	// The ID of the webhook configuration to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_casb_policy#webhook_config_id ZeroTrustCasbPolicy#webhook_config_id}
	WebhookConfigId *string `field:"required" json:"webhookConfigId" yaml:"webhookConfigId"`
}

