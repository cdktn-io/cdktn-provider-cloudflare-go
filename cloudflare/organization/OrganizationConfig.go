// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package organization

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OrganizationConfig struct {
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
	// The Organization name, limited to 65 UTF-8 bytes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/organization#name Organization#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/organization#parent Organization#parent}.
	Parent *OrganizationParent `field:"optional" json:"parent" yaml:"parent"`
	// The current parent Organization ID. A different value is rejected; this endpoint cannot move an Organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/organization#parent_id Organization#parent_id}
	ParentId *string `field:"optional" json:"parentId" yaml:"parentId"`
	// Optional profile fields, each limited to 1,000 UTF-8 bytes. An empty profile is ignored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/organization#profile Organization#profile}
	Profile *OrganizationProfile `field:"optional" json:"profile" yaml:"profile"`
}

