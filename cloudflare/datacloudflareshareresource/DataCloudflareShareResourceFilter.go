// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflareshareresource


type DataCloudflareShareResourceFilter struct {
	// Filter share resources by resource_type. Available values: "custom-ruleset", "gateway-policy", "gateway-destination-ip", "gateway-block-page-settings", "gateway-extended-email-matching", "idp-federation-grant".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/data-sources/share_resource#resource_type DataCloudflareShareResource#resource_type}
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// Filter share resources by status. Available values: "active", "deleting", "deleted".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/data-sources/share_resource#status DataCloudflareShareResource#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}

