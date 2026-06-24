// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dlsprefixbinding

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DlsPrefixBindingConfig struct {
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
	// Identifier of a Cloudflare account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.1/docs/resources/dls_prefix_binding#account_id DlsPrefixBinding#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// IP prefix in CIDR notation to bind.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.1/docs/resources/dls_prefix_binding#cidr DlsPrefixBinding#cidr}
	Cidr *string `field:"required" json:"cidr" yaml:"cidr"`
	// The ID of the parent IP prefix that contains the CIDR.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.1/docs/resources/dls_prefix_binding#prefix_id DlsPrefixBinding#prefix_id}
	PrefixId *string `field:"required" json:"prefixId" yaml:"prefixId"`
	// Region key from managed regions (e.g., "us", "eu").
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.1/docs/resources/dls_prefix_binding#region_key DlsPrefixBinding#region_key}
	RegionKey *string `field:"required" json:"regionKey" yaml:"regionKey"`
}

