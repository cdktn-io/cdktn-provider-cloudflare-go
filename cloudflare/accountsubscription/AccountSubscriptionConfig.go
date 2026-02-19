// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accountsubscription

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AccountSubscriptionConfig struct {
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
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/account_subscription#account_id AccountSubscription#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// How often the subscription is renewed automatically. Available values: "weekly", "monthly", "quarterly", "yearly".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/account_subscription#frequency AccountSubscription#frequency}
	Frequency *string `field:"optional" json:"frequency" yaml:"frequency"`
	// The rate plan applied to the subscription.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/account_subscription#rate_plan AccountSubscription#rate_plan}
	RatePlan *AccountSubscriptionRatePlan `field:"optional" json:"ratePlan" yaml:"ratePlan"`
}

