// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflareemailsecurityallowpolicy


type DataCloudflareEmailSecurityAllowPolicyFilter struct {
	// The sorting direction. Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/email_security_allow_policy#direction DataCloudflareEmailSecurityAllowPolicy#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Filter to show only policies where messages from the sender are exempted from Spam, Spoof, and Bulk dispositions (not Malicious or Suspicious).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/email_security_allow_policy#is_acceptable_sender DataCloudflareEmailSecurityAllowPolicy#is_acceptable_sender}
	IsAcceptableSender interface{} `field:"optional" json:"isAcceptableSender" yaml:"isAcceptableSender"`
	// Filter to show only policies where messages to the recipient bypass all detections.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/email_security_allow_policy#is_exempt_recipient DataCloudflareEmailSecurityAllowPolicy#is_exempt_recipient}
	IsExemptRecipient interface{} `field:"optional" json:"isExemptRecipient" yaml:"isExemptRecipient"`
	// Filter to show only policies where messages from the sender bypass all detections and link following.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/email_security_allow_policy#is_trusted_sender DataCloudflareEmailSecurityAllowPolicy#is_trusted_sender}
	IsTrustedSender interface{} `field:"optional" json:"isTrustedSender" yaml:"isTrustedSender"`
	// Field to sort by. Available values: "pattern", "created_at".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/email_security_allow_policy#order DataCloudflareEmailSecurityAllowPolicy#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
	// Filter by exact pattern value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/email_security_allow_policy#pattern DataCloudflareEmailSecurityAllowPolicy#pattern}
	Pattern *string `field:"optional" json:"pattern" yaml:"pattern"`
	// Filter by pattern type. Available values: "EMAIL", "DOMAIN", "IP", "UNKNOWN".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/email_security_allow_policy#pattern_type DataCloudflareEmailSecurityAllowPolicy#pattern_type}
	PatternType *string `field:"optional" json:"patternType" yaml:"patternType"`
	// Search term for filtering records. Behavior may change.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/email_security_allow_policy#search DataCloudflareEmailSecurityAllowPolicy#search}
	Search *string `field:"optional" json:"search" yaml:"search"`
	// Filter to show only policies that enforce DMARC, SPF, or DKIM authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/email_security_allow_policy#verify_sender DataCloudflareEmailSecurityAllowPolicy#verify_sender}
	VerifySender interface{} `field:"optional" json:"verifySender" yaml:"verifySender"`
}

