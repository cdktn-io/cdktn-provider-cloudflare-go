// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflareemailsecurityimpersonationregistries

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataCloudflareEmailSecurityImpersonationRegistriesConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/data-sources/email_security_impersonation_registries#account_id DataCloudflareEmailSecurityImpersonationRegistries#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// The sorting direction. Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/data-sources/email_security_impersonation_registries#direction DataCloudflareEmailSecurityImpersonationRegistries#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Max items to fetch, default: 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/data-sources/email_security_impersonation_registries#max_items DataCloudflareEmailSecurityImpersonationRegistries#max_items}
	MaxItems *float64 `field:"optional" json:"maxItems" yaml:"maxItems"`
	// Field to sort by. Available values: "name", "email", "created_at".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/data-sources/email_security_impersonation_registries#order DataCloudflareEmailSecurityImpersonationRegistries#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
	// Available values: "A1S_INTERNAL", "SNOOPY-CASB_OFFICE_365", "SNOOPY-OFFICE_365", "SNOOPY-GOOGLE_DIRECTORY".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/data-sources/email_security_impersonation_registries#provenance DataCloudflareEmailSecurityImpersonationRegistries#provenance}
	Provenance *string `field:"optional" json:"provenance" yaml:"provenance"`
	// Search term for filtering records. Behavior may change.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/data-sources/email_security_impersonation_registries#search DataCloudflareEmailSecurityImpersonationRegistries#search}
	Search *string `field:"optional" json:"search" yaml:"search"`
}

