// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustgatewaypolicy


type ZeroTrustGatewayPolicyRuleSettingsForward struct {
	// Specify the public Worker ID of the target Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/zero_trust_gateway_policy#worker_id ZeroTrustGatewayPolicy#worker_id}
	WorkerId *string `field:"required" json:"workerId" yaml:"workerId"`
}

