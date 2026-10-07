// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflareflagshipflag

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataCloudflareFlagshipFlagConfig struct {
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
	// Cloudflare account ID that owns the Flagship app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/flagship_flag#account_id DataCloudflareFlagshipFlag#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Flagship app ID returned when the app was created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/flagship_flag#app_id DataCloudflareFlagshipFlag#app_id}
	AppId *string `field:"required" json:"appId" yaml:"appId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/flagship_flag#filter DataCloudflareFlagshipFlag#filter}.
	Filter *DataCloudflareFlagshipFlagFilter `field:"optional" json:"filter" yaml:"filter"`
	// Case-sensitive key identifying the flag within the app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/flagship_flag#flag_key DataCloudflareFlagshipFlag#flag_key}
	FlagKey *string `field:"optional" json:"flagKey" yaml:"flagKey"`
}

