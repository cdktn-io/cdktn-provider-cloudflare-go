// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustcasbpolicy

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ZeroTrustCasbPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_casb_policy#account_id ZeroTrustCasbPolicy#account_id}.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Actions to execute when this policy is triggered, grouped by action type.
	//
	// A policy must contain at least one action across all groups and may include
	// at most one remediation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_casb_policy#actions ZeroTrustCasbPolicy#actions}
	Actions *ZeroTrustCasbPolicyActions `field:"required" json:"actions" yaml:"actions"`
	// When true, the policy applies to all integrations for the account. When false, integration_ids must be provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_casb_policy#applies_to_all_integrations ZeroTrustCasbPolicy#applies_to_all_integrations}
	AppliesToAllIntegrations interface{} `field:"required" json:"appliesToAllIntegrations" yaml:"appliesToAllIntegrations"`
	// Display name for the policy configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_casb_policy#display_name ZeroTrustCasbPolicy#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Boolean specifying if the policy is enabled or disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_casb_policy#enabled ZeroTrustCasbPolicy#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The finding type this policy is associated with. All remediation actions must match this finding type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_casb_policy#finding_type_id ZeroTrustCasbPolicy#finding_type_id}
	FindingTypeId *string `field:"required" json:"findingTypeId" yaml:"findingTypeId"`
	// Optional description of what this policy does.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_casb_policy#description ZeroTrustCasbPolicy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The integrations this policy applies to. Required when applies_to_all_integrations is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_casb_policy#integration_ids ZeroTrustCasbPolicy#integration_ids}
	IntegrationIds *[]*string `field:"optional" json:"integrationIds" yaml:"integrationIds"`
}

