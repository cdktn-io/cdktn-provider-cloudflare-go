// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sharerecipient

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ShareRecipientConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/share_recipient#account_id ShareRecipient#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Share identifier tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/share_recipient#share_id ShareRecipient#share_id}
	ShareId *string `field:"required" json:"shareId" yaml:"shareId"`
	// Organization identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/share_recipient#organization_id ShareRecipient#organization_id}
	OrganizationId *string `field:"optional" json:"organizationId" yaml:"organizationId"`
	// The account that will receive the share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/share_recipient#recipient_account_id ShareRecipient#recipient_account_id}
	RecipientAccountId *string `field:"optional" json:"recipientAccountId" yaml:"recipientAccountId"`
}

