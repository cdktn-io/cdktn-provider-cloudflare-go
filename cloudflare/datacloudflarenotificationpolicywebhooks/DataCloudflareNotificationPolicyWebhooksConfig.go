// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflarenotificationpolicywebhooks

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataCloudflareNotificationPolicyWebhooksConfig struct {
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
	// The account id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/data-sources/notification_policy_webhooks#account_id DataCloudflareNotificationPolicyWebhooks#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The unique identifier of a webhook.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/data-sources/notification_policy_webhooks#webhook_id DataCloudflareNotificationPolicyWebhooks#webhook_id}
	WebhookId *string `field:"required" json:"webhookId" yaml:"webhookId"`
}

