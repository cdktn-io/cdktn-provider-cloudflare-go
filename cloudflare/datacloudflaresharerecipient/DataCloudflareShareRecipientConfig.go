// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflaresharerecipient

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataCloudflareShareRecipientConfig struct {
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
	// Account identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/data-sources/share_recipient#account_id DataCloudflareShareRecipient#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Share Recipient identifier tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/data-sources/share_recipient#recipient_id DataCloudflareShareRecipient#recipient_id}
	RecipientId *string `field:"required" json:"recipientId" yaml:"recipientId"`
	// Share identifier tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/data-sources/share_recipient#share_id DataCloudflareShareRecipient#share_id}
	ShareId *string `field:"required" json:"shareId" yaml:"shareId"`
	// Include resources in the response.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/data-sources/share_recipient#include_resources DataCloudflareShareRecipient#include_resources}
	IncludeResources interface{} `field:"optional" json:"includeResources" yaml:"includeResources"`
}

