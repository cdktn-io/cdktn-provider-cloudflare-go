// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflareemailsecurityallowpolicies

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataCloudflareEmailSecurityAllowPoliciesConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/email_security_allow_policies#account_id DataCloudflareEmailSecurityAllowPolicies#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The sorting direction. Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/email_security_allow_policies#direction DataCloudflareEmailSecurityAllowPolicies#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Filter to show only policies where messages from the sender are exempted from Spam, Spoof, and Bulk dispositions (not Malicious or Suspicious).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/email_security_allow_policies#is_acceptable_sender DataCloudflareEmailSecurityAllowPolicies#is_acceptable_sender}
	IsAcceptableSender interface{} `field:"optional" json:"isAcceptableSender" yaml:"isAcceptableSender"`
	// Filter to show only policies where messages to the recipient bypass all detections.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/email_security_allow_policies#is_exempt_recipient DataCloudflareEmailSecurityAllowPolicies#is_exempt_recipient}
	IsExemptRecipient interface{} `field:"optional" json:"isExemptRecipient" yaml:"isExemptRecipient"`
	// Filter to show only policies where messages from the sender bypass all detections and link following.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/email_security_allow_policies#is_trusted_sender DataCloudflareEmailSecurityAllowPolicies#is_trusted_sender}
	IsTrustedSender interface{} `field:"optional" json:"isTrustedSender" yaml:"isTrustedSender"`
	// Max items to fetch, default: 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/email_security_allow_policies#max_items DataCloudflareEmailSecurityAllowPolicies#max_items}
	MaxItems *float64 `field:"optional" json:"maxItems" yaml:"maxItems"`
	// Field to sort by. Available values: "pattern", "created_at".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/email_security_allow_policies#order DataCloudflareEmailSecurityAllowPolicies#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
	// Filter by exact pattern value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/email_security_allow_policies#pattern DataCloudflareEmailSecurityAllowPolicies#pattern}
	Pattern *string `field:"optional" json:"pattern" yaml:"pattern"`
	// Filter by pattern type. Available values: "EMAIL", "DOMAIN", "IP", "UNKNOWN".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/email_security_allow_policies#pattern_type DataCloudflareEmailSecurityAllowPolicies#pattern_type}
	PatternType *string `field:"optional" json:"patternType" yaml:"patternType"`
	// Search term for filtering records. Behavior may change.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/email_security_allow_policies#search DataCloudflareEmailSecurityAllowPolicies#search}
	Search *string `field:"optional" json:"search" yaml:"search"`
	// Filter to show only policies that enforce DMARC, SPF, or DKIM authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/email_security_allow_policies#verify_sender DataCloudflareEmailSecurityAllowPolicies#verify_sender}
	VerifySender interface{} `field:"optional" json:"verifySender" yaml:"verifySender"`
}

