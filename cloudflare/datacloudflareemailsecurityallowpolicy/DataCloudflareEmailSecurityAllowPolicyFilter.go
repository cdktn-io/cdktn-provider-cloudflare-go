// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflareemailsecurityallowpolicy


type DataCloudflareEmailSecurityAllowPolicyFilter struct {
	// The sorting direction. Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.25.0/docs/data-sources/email_security_allow_policy#direction DataCloudflareEmailSecurityAllowPolicy#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Filter to show only policies where messages from the sender are exempted from Spam, Spoof, and Bulk dispositions (not Malicious or Suspicious).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.25.0/docs/data-sources/email_security_allow_policy#is_acceptable_sender DataCloudflareEmailSecurityAllowPolicy#is_acceptable_sender}
	IsAcceptableSender interface{} `field:"optional" json:"isAcceptableSender" yaml:"isAcceptableSender"`
	// Filter to show only policies where messages to the recipient bypass all detections.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.25.0/docs/data-sources/email_security_allow_policy#is_exempt_recipient DataCloudflareEmailSecurityAllowPolicy#is_exempt_recipient}
	IsExemptRecipient interface{} `field:"optional" json:"isExemptRecipient" yaml:"isExemptRecipient"`
	// Filter to show only policies where messages from the sender bypass all detections and link following.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.25.0/docs/data-sources/email_security_allow_policy#is_trusted_sender DataCloudflareEmailSecurityAllowPolicy#is_trusted_sender}
	IsTrustedSender interface{} `field:"optional" json:"isTrustedSender" yaml:"isTrustedSender"`
	// Field to sort by. Available values: "pattern", "created_at".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.25.0/docs/data-sources/email_security_allow_policy#order DataCloudflareEmailSecurityAllowPolicy#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.25.0/docs/data-sources/email_security_allow_policy#pattern DataCloudflareEmailSecurityAllowPolicy#pattern}.
	Pattern *string `field:"optional" json:"pattern" yaml:"pattern"`
	// Type of pattern matching.
	//
	// - EMAIL: matches a full email address (e.g. `user@example.com`)
	// - DOMAIN: matches a domain name (e.g. `example.com`)
	// - IP: matches a plain IPv4 or IPv6 address (e.g. `1.2.3.4` or `2606:4700:4700::1111`) or CIDR block (e.g. `1.2.3.0/24` or `2606:4700:4700::/48`). The API rejects private or unique-local, loopback, link-local, unspecified, and IPv4 broadcast addresses, including their IPv4-mapped IPv6 equivalents.
	// - UNKNOWN: deprecated; you cannot use this when creating or updating policies, but it may appear on existing entries.
	// Available values: "EMAIL", "DOMAIN", "IP", "UNKNOWN".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.25.0/docs/data-sources/email_security_allow_policy#pattern_type DataCloudflareEmailSecurityAllowPolicy#pattern_type}
	PatternType *string `field:"optional" json:"patternType" yaml:"patternType"`
	// Search term for filtering records. Behavior may change.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.25.0/docs/data-sources/email_security_allow_policy#search DataCloudflareEmailSecurityAllowPolicy#search}
	Search *string `field:"optional" json:"search" yaml:"search"`
	// Filter to show only policies that enforce DMARC, SPF, or DKIM authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.25.0/docs/data-sources/email_security_allow_policy#verify_sender DataCloudflareEmailSecurityAllowPolicy#verify_sender}
	VerifySender interface{} `field:"optional" json:"verifySender" yaml:"verifySender"`
}

