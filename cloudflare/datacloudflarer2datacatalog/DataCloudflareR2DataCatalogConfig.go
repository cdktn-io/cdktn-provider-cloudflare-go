// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflarer2datacatalog

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataCloudflareR2DataCatalogConfig struct {
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
	// Specifies the R2 bucket name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/data-sources/r2_data_catalog#bucket_name DataCloudflareR2DataCatalog#bucket_name}
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Use this to identify the account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/data-sources/r2_data_catalog#account_id DataCloudflareR2DataCatalog#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
}

