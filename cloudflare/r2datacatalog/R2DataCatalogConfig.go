// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package r2datacatalog

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type R2DataCatalogConfig struct {
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
	// Use this to identify the account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.0/docs/resources/r2_data_catalog#account_id R2DataCatalog#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Specifies the R2 bucket name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.0/docs/resources/r2_data_catalog#bucket_name R2DataCatalog#bucket_name}
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
}

