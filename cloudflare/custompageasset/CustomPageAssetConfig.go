// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package custompageasset

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CustomPageAssetConfig struct {
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
	// A short description of the custom asset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/custom_page_asset#description CustomPageAsset#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// The unique name of the custom asset. Can only contain letters (A-Z, a-z), numbers (0-9), and underscores (_).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/custom_page_asset#name CustomPageAsset#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The URL where the asset content is fetched from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/custom_page_asset#url CustomPageAsset#url}
	Url *string `field:"required" json:"url" yaml:"url"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/custom_page_asset#account_id CustomPageAsset#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/custom_page_asset#zone_id CustomPageAsset#zone_id}
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}

