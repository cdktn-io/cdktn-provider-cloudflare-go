// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflarecustomcsr

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataCloudflareCustomCsrConfig struct {
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
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/data-sources/custom_csr#account_id DataCloudflareCustomCsr#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Custom CSR identifier tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/data-sources/custom_csr#custom_csr_id DataCloudflareCustomCsr#custom_csr_id}
	CustomCsrId *string `field:"optional" json:"customCsrId" yaml:"customCsrId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/data-sources/custom_csr#filter DataCloudflareCustomCsr#filter}.
	Filter *DataCloudflareCustomCsrFilter `field:"optional" json:"filter" yaml:"filter"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/data-sources/custom_csr#zone_id DataCloudflareCustomCsr#zone_id}
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}

