// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package emailroutingrule


type EmailRoutingRuleActions struct {
	// Type of supported action. Available values: "drop", "forward", "worker".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/email_routing_rule#type EmailRoutingRule#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// List of values for the action. Currently limited to a single value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/email_routing_rule#value EmailRoutingRule#value}
	Value *[]*string `field:"optional" json:"value" yaml:"value"`
}

