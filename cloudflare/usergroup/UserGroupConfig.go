// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package usergroup

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type UserGroupConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.0/docs/resources/user_group#account_id UserGroup#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Name of the User group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.0/docs/resources/user_group#name UserGroup#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Policies attached to the User group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.0/docs/resources/user_group#policies UserGroup#policies}
	Policies interface{} `field:"optional" json:"policies" yaml:"policies"`
}

