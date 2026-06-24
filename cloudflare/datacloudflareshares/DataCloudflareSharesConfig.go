// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflareshares

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataCloudflareSharesConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Account identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.1/docs/data-sources/shares#account_id DataCloudflareShares#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Direction to sort objects. Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.1/docs/data-sources/shares#direction DataCloudflareShares#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Include recipient counts in the response.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.1/docs/data-sources/shares#include_recipient_counts DataCloudflareShares#include_recipient_counts}
	IncludeRecipientCounts interface{} `field:"optional" json:"includeRecipientCounts" yaml:"includeRecipientCounts"`
	// Include resources in the response.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.1/docs/data-sources/shares#include_resources DataCloudflareShares#include_resources}
	IncludeResources interface{} `field:"optional" json:"includeResources" yaml:"includeResources"`
	// Filter shares by kind. Available values: "sent", "received".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.1/docs/data-sources/shares#kind DataCloudflareShares#kind}
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
	// Max items to fetch, default: 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.1/docs/data-sources/shares#max_items DataCloudflareShares#max_items}
	MaxItems *float64 `field:"optional" json:"maxItems" yaml:"maxItems"`
	// Order shares by values in the given field. Available values: "name", "created".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.1/docs/data-sources/shares#order DataCloudflareShares#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
	// Filter share resources by resource_types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.1/docs/data-sources/shares#resource_types DataCloudflareShares#resource_types}
	ResourceTypes *[]*string `field:"optional" json:"resourceTypes" yaml:"resourceTypes"`
	// Filter shares by status. Available values: "active", "deleting", "deleted".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.1/docs/data-sources/shares#status DataCloudflareShares#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Filter shares by tag.
	//
	// Each value is either `key=value` (matches shares whose tags contain that key/value pair) or `key` alone (matches shares that have any value for that key). May be repeated; multiple `tag` parameters are ANDed together. Maximum 20 `tag` parameters per request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.1/docs/data-sources/shares#tag DataCloudflareShares#tag}
	Tag *[]*string `field:"optional" json:"tag" yaml:"tag"`
	// Filter shares by target_type. Available values: "account", "organization".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.1/docs/data-sources/shares#target_type DataCloudflareShares#target_type}
	TargetType *string `field:"optional" json:"targetType" yaml:"targetType"`
}

