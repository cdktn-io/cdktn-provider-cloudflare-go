// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package shareresource

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ShareResourceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/resources/share_resource#account_id ShareResource#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Resource Metadata.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/resources/share_resource#meta ShareResource#meta}
	Meta *string `field:"required" json:"meta" yaml:"meta"`
	// Account identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/resources/share_resource#resource_account_id ShareResource#resource_account_id}
	ResourceAccountId *string `field:"required" json:"resourceAccountId" yaml:"resourceAccountId"`
	// Share Resource identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/resources/share_resource#resource_id ShareResource#resource_id}
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// Resource Type. Available values: "custom-ruleset", "gateway-policy", "gateway-destination-ip", "gateway-block-page-settings", "gateway-extended-email-matching", "idp-federation-grant".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/resources/share_resource#resource_type ShareResource#resource_type}
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
	// Share identifier tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/resources/share_resource#share_id ShareResource#share_id}
	ShareId *string `field:"required" json:"shareId" yaml:"shareId"`
}

