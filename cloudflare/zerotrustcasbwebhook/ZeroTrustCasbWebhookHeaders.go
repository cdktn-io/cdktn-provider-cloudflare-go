// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustcasbwebhook


type ZeroTrustCasbWebhookHeaders struct {
	// Header key name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_casb_webhook#key ZeroTrustCasbWebhook#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Header value. Required on Create and Evaluate. On Update, omit or set to null to keep existing value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_casb_webhook#value ZeroTrustCasbWebhook#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

