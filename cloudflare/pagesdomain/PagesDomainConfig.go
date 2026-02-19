// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pagesdomain

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PagesDomainConfig struct {
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
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/pages_domain#account_id PagesDomain#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The domain name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/pages_domain#name PagesDomain#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Name of the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/pages_domain#project_name PagesDomain#project_name}
	ProjectName *string `field:"required" json:"projectName" yaml:"projectName"`
}

