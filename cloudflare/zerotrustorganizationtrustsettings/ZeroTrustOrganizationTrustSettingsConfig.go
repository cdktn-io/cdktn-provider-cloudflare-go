// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustorganizationtrustsettings

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ZeroTrustOrganizationTrustSettingsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/zero_trust_organization_trust_settings#account_id ZeroTrustOrganizationTrustSettings#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Whether all accounts in the organization are trusted to use this account's cross-account Access resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/zero_trust_organization_trust_settings#grant_trust_to_all_accounts_in_org ZeroTrustOrganizationTrustSettings#grant_trust_to_all_accounts_in_org}
	GrantTrustToAllAccountsInOrg interface{} `field:"required" json:"grantTrustToAllAccountsInOrg" yaml:"grantTrustToAllAccountsInOrg"`
	// Explicit account IDs trusted to use this account's cross-account Access resources.
	//
	// At most 100 unique valid account IDs are accepted. Omitted from responses when empty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/zero_trust_organization_trust_settings#trusting_accounts ZeroTrustOrganizationTrustSettings#trusting_accounts}
	TrustingAccounts *[]*string `field:"optional" json:"trustingAccounts" yaml:"trustingAccounts"`
}

