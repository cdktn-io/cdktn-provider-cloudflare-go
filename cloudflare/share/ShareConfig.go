// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package share

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ShareConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/share#account_id Share#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The name of the share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/share#name Share#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Accounts to exclude from an organization-targeted share.
	//
	// Only allowed when `recipients` contains no account entries. The share owner cannot be excluded. Duplicate entries are ignored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/share#excluded_recipients Share#excluded_recipients}
	ExcludedRecipients interface{} `field:"optional" json:"excludedRecipients" yaml:"excludedRecipients"`
	// Include recipient counts in the response.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/share#include_recipient_counts Share#include_recipient_counts}
	IncludeRecipientCounts interface{} `field:"optional" json:"includeRecipientCounts" yaml:"includeRecipientCounts"`
	// Include resources in the response.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/share#include_resources Share#include_resources}
	IncludeResources interface{} `field:"optional" json:"includeResources" yaml:"includeResources"`
	// Include share tags in the response.
	//
	// `api` reads tags from the tag service; `mirror` reads a local, eventually consistent replica. If the tag lookup fails, the response omits `tags`.
	// Available values: "api", "mirror".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/share#include_tags_from Share#include_tags_from}
	IncludeTagsFrom *string `field:"optional" json:"includeTagsFrom" yaml:"includeTagsFrom"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/share#recipients Share#recipients}.
	Recipients interface{} `field:"optional" json:"recipients" yaml:"recipients"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/share#resources Share#resources}.
	Resources interface{} `field:"optional" json:"resources" yaml:"resources"`
	// Tags associated with the share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/share#tags Share#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

