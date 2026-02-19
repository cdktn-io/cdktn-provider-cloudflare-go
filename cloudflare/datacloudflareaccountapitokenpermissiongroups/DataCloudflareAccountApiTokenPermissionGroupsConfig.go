// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflareaccountapitokenpermissiongroups

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataCloudflareAccountApiTokenPermissionGroupsConfig struct {
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
	// Account identifier tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/data-sources/account_api_token_permission_groups#account_id DataCloudflareAccountApiTokenPermissionGroups#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Filter by the name of the permission group. The value must be URL-encoded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/data-sources/account_api_token_permission_groups#name DataCloudflareAccountApiTokenPermissionGroups#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Filter by the scope of the permission group. The value must be URL-encoded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/data-sources/account_api_token_permission_groups#scope DataCloudflareAccountApiTokenPermissionGroups#scope}
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
}

