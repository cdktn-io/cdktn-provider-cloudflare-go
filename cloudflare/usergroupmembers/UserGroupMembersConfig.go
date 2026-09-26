// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package usergroupmembers

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type UserGroupMembersConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/resources/user_group_members#account_id UserGroupMembers#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/resources/user_group_members#members UserGroupMembers#members}.
	Members interface{} `field:"required" json:"members" yaml:"members"`
	// User Group identifier tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/resources/user_group_members#user_group_id UserGroupMembers#user_group_id}
	UserGroupId *string `field:"required" json:"userGroupId" yaml:"userGroupId"`
	// The sort order of returned user group members by email. Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/resources/user_group_members#direction UserGroupMembers#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// A string used for filtering members by partial email match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/resources/user_group_members#fuzzy_email UserGroupMembers#fuzzy_email}
	FuzzyEmail *string `field:"optional" json:"fuzzyEmail" yaml:"fuzzyEmail"`
	// Page number of paginated results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/resources/user_group_members#page UserGroupMembers#page}
	Page *float64 `field:"optional" json:"page" yaml:"page"`
	// Maximum number of results per page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/resources/user_group_members#per_page UserGroupMembers#per_page}
	PerPage *float64 `field:"optional" json:"perPage" yaml:"perPage"`
}

