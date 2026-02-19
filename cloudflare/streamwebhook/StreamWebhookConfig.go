// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package streamwebhook

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type StreamWebhookConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/stream_webhook#account_id StreamWebhook#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The URL where webhooks will be sent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/stream_webhook#notification_url StreamWebhook#notification_url}
	NotificationUrl *string `field:"required" json:"notificationUrl" yaml:"notificationUrl"`
}

