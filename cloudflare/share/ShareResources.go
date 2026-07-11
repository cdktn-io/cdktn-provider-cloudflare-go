// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package share


type ShareResources struct {
	// Resource Metadata.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/resources/share#meta Share#meta}
	Meta *string `field:"required" json:"meta" yaml:"meta"`
	// Account identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/resources/share#resource_account_id Share#resource_account_id}
	ResourceAccountId *string `field:"required" json:"resourceAccountId" yaml:"resourceAccountId"`
	// Share Resource identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/resources/share#resource_id Share#resource_id}
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// Resource Type. Available values: "custom-ruleset", "gateway-policy", "gateway-destination-ip", "gateway-block-page-settings", "gateway-extended-email-matching", "idp-federation-grant".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/resources/share#resource_type Share#resource_type}
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
}

