// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package emailsecurityallowpolicy

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EmailSecurityAllowPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.25.0/docs/resources/email_security_allow_policy#account_id EmailSecurityAllowPolicy#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Exempts messages from this sender from Spam, Spoof and Bulk dispositions only; Malicious and Suspicious dispositions still apply.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.25.0/docs/resources/email_security_allow_policy#is_acceptable_sender EmailSecurityAllowPolicy#is_acceptable_sender}
	IsAcceptableSender interface{} `field:"required" json:"isAcceptableSender" yaml:"isAcceptableSender"`
	// Bypasses all detections for messages to this recipient.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.25.0/docs/resources/email_security_allow_policy#is_exempt_recipient EmailSecurityAllowPolicy#is_exempt_recipient}
	IsExemptRecipient interface{} `field:"required" json:"isExemptRecipient" yaml:"isExemptRecipient"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.25.0/docs/resources/email_security_allow_policy#is_regex EmailSecurityAllowPolicy#is_regex}.
	IsRegex interface{} `field:"required" json:"isRegex" yaml:"isRegex"`
	// Bypasses all detections and link following for messages from this sender.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.25.0/docs/resources/email_security_allow_policy#is_trusted_sender EmailSecurityAllowPolicy#is_trusted_sender}
	IsTrustedSender interface{} `field:"required" json:"isTrustedSender" yaml:"isTrustedSender"`
	// The pattern value to match.
	//
	// The format depends on `pattern_type`: a valid email address for EMAIL (e.g. `user@example.com`), a valid domain name for DOMAIN (e.g. `example.com`), or a plain IPv4 or IPv6 address or CIDR block for IP (e.g. `1.2.3.4`, `1.2.3.0/24`, `2606:4700:4700::1111`, or `2606:4700:4700::/48`); the API rejects private or unique-local, loopback, link-local, unspecified, and IPv4 broadcast addresses, including their IPv4-mapped IPv6 equivalents.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.25.0/docs/resources/email_security_allow_policy#pattern EmailSecurityAllowPolicy#pattern}
	Pattern *string `field:"required" json:"pattern" yaml:"pattern"`
	// Type of pattern matching.
	//
	// - EMAIL: matches a full email address (e.g. `user@example.com`)
	// - DOMAIN: matches a domain name (e.g. `example.com`)
	// - IP: matches a plain IPv4 or IPv6 address (e.g. `1.2.3.4` or `2606:4700:4700::1111`) or CIDR block (e.g. `1.2.3.0/24` or `2606:4700:4700::/48`). The API rejects private or unique-local, loopback, link-local, unspecified, and IPv4 broadcast addresses, including their IPv4-mapped IPv6 equivalents.
	// - UNKNOWN: deprecated; you cannot use this when creating or updating policies, but it may appear on existing entries.
	// Available values: "EMAIL", "DOMAIN", "IP", "UNKNOWN".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.25.0/docs/resources/email_security_allow_policy#pattern_type EmailSecurityAllowPolicy#pattern_type}
	PatternType *string `field:"required" json:"patternType" yaml:"patternType"`
	// Enforce DMARC, SPF or DKIM authentication. When on, Email Security only honors policies that pass authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.25.0/docs/resources/email_security_allow_policy#verify_sender EmailSecurityAllowPolicy#verify_sender}
	VerifySender interface{} `field:"required" json:"verifySender" yaml:"verifySender"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.25.0/docs/resources/email_security_allow_policy#comments EmailSecurityAllowPolicy#comments}.
	Comments *string `field:"optional" json:"comments" yaml:"comments"`
	// Deprecated as of July 1, 2025. Use `is_exempt_recipient` instead. End of life: July 1, 2026.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.25.0/docs/resources/email_security_allow_policy#is_recipient EmailSecurityAllowPolicy#is_recipient}
	IsRecipient interface{} `field:"optional" json:"isRecipient" yaml:"isRecipient"`
	// Deprecated as of July 1, 2025. Use `is_trusted_sender` instead. End of life: July 1, 2026.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.25.0/docs/resources/email_security_allow_policy#is_sender EmailSecurityAllowPolicy#is_sender}
	IsSender interface{} `field:"optional" json:"isSender" yaml:"isSender"`
	// Deprecated as of July 1, 2025. Use `is_acceptable_sender` instead. End of life: July 1, 2026.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.25.0/docs/resources/email_security_allow_policy#is_spoof EmailSecurityAllowPolicy#is_spoof}
	IsSpoof interface{} `field:"optional" json:"isSpoof" yaml:"isSpoof"`
}

