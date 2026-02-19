// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package emailsecurityblocksender

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EmailSecurityBlockSenderConfig struct {
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
	// Account Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/email_security_block_sender#account_id EmailSecurityBlockSender#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/email_security_block_sender#is_regex EmailSecurityBlockSender#is_regex}.
	IsRegex interface{} `field:"required" json:"isRegex" yaml:"isRegex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/email_security_block_sender#pattern EmailSecurityBlockSender#pattern}.
	Pattern *string `field:"required" json:"pattern" yaml:"pattern"`
	// Available values: "EMAIL", "DOMAIN", "IP", "UNKNOWN".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/email_security_block_sender#pattern_type EmailSecurityBlockSender#pattern_type}
	PatternType *string `field:"required" json:"patternType" yaml:"patternType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/email_security_block_sender#comments EmailSecurityBlockSender#comments}.
	Comments *string `field:"optional" json:"comments" yaml:"comments"`
}

