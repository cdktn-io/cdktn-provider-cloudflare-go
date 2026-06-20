// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflareshare


type DataCloudflareShareFilter struct {
	// Direction to sort objects. Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.0/docs/data-sources/share#direction DataCloudflareShare#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Filter shares by kind. Available values: "sent", "received".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.0/docs/data-sources/share#kind DataCloudflareShare#kind}
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
	// Order shares by values in the given field. Available values: "name", "created".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.0/docs/data-sources/share#order DataCloudflareShare#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
	// Filter share resources by resource_types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.0/docs/data-sources/share#resource_types DataCloudflareShare#resource_types}
	ResourceTypes *[]*string `field:"optional" json:"resourceTypes" yaml:"resourceTypes"`
	// Filter shares by status. Available values: "active", "deleting", "deleted".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.0/docs/data-sources/share#status DataCloudflareShare#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Filter shares by tag.
	//
	// Each value is either `key=value` (matches shares whose tags contain that key/value pair) or `key` alone (matches shares that have any value for that key). May be repeated; multiple `tag` parameters are ANDed together. Maximum 20 `tag` parameters per request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.0/docs/data-sources/share#tag DataCloudflareShare#tag}
	Tag *[]*string `field:"optional" json:"tag" yaml:"tag"`
	// Filter shares by target_type. Available values: "account", "organization".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.0/docs/data-sources/share#target_type DataCloudflareShare#target_type}
	TargetType *string `field:"optional" json:"targetType" yaml:"targetType"`
}

