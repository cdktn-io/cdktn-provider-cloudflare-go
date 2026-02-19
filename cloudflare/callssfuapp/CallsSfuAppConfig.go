// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package callssfuapp

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CallsSfuAppConfig struct {
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
	// The account identifier tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/calls_sfu_app#account_id CallsSfuApp#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// A Cloudflare-generated unique identifier for a item.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/calls_sfu_app#app_id CallsSfuApp#app_id}
	AppId *string `field:"optional" json:"appId" yaml:"appId"`
	// A short description of Calls app, not shown to end users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/calls_sfu_app#name CallsSfuApp#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

