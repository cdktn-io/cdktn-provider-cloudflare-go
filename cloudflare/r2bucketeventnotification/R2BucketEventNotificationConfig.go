// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package r2bucketeventnotification

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type R2BucketEventNotificationConfig struct {
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
	// Cloudflare account ID that owns the R2 resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/r2_bucket_event_notification#account_id R2BucketEventNotification#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Name of the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/r2_bucket_event_notification#bucket_name R2BucketEventNotification#bucket_name}
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// ID of the Cloudflare Queue that receives notifications for matching R2 object events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/r2_bucket_event_notification#queue_id R2BucketEventNotification#queue_id}
	QueueId *string `field:"required" json:"queueId" yaml:"queueId"`
	// Array of rules to drive notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/r2_bucket_event_notification#rules R2BucketEventNotification#rules}
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
	// Jurisdiction of the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/r2_bucket_event_notification#jurisdiction R2BucketEventNotification#jurisdiction}
	Jurisdiction *string `field:"optional" json:"jurisdiction" yaml:"jurisdiction"`
}

