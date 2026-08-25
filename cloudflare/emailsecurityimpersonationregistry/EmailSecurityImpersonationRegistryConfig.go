// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package emailsecurityimpersonationregistry

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EmailSecurityImpersonationRegistryConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/email_security_impersonation_registry#account_id EmailSecurityImpersonationRegistry#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/email_security_impersonation_registry#email EmailSecurityImpersonationRegistry#email}.
	Email *string `field:"required" json:"email" yaml:"email"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/email_security_impersonation_registry#is_email_regex EmailSecurityImpersonationRegistry#is_email_regex}.
	IsEmailRegex interface{} `field:"required" json:"isEmailRegex" yaml:"isEmailRegex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/email_security_impersonation_registry#name EmailSecurityImpersonationRegistry#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/email_security_impersonation_registry#comments EmailSecurityImpersonationRegistry#comments}.
	Comments *string `field:"optional" json:"comments" yaml:"comments"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/email_security_impersonation_registry#directory_id EmailSecurityImpersonationRegistry#directory_id}.
	DirectoryId *float64 `field:"optional" json:"directoryId" yaml:"directoryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/email_security_impersonation_registry#directory_node_id EmailSecurityImpersonationRegistry#directory_node_id}.
	DirectoryNodeId *float64 `field:"optional" json:"directoryNodeId" yaml:"directoryNodeId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/email_security_impersonation_registry#external_directory_node_id EmailSecurityImpersonationRegistry#external_directory_node_id}.
	ExternalDirectoryNodeId *string `field:"optional" json:"externalDirectoryNodeId" yaml:"externalDirectoryNodeId"`
	// Available values: "A1S_INTERNAL", "SNOOPY-CASB_OFFICE_365", "SNOOPY-OFFICE_365", "SNOOPY-GOOGLE_DIRECTORY".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/email_security_impersonation_registry#provenance EmailSecurityImpersonationRegistry#provenance}
	Provenance *string `field:"optional" json:"provenance" yaml:"provenance"`
}

