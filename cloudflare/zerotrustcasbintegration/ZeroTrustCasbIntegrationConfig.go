// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustcasbintegration

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ZeroTrustCasbIntegrationConfig struct {
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
	// Cloudflare account identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_casb_integration#account_id ZeroTrustCasbIntegration#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Name of the integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_casb_integration#name ZeroTrustCasbIntegration#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Whether the integration is paused.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_casb_integration#paused ZeroTrustCasbIntegration#paused}
	Paused interface{} `field:"required" json:"paused" yaml:"paused"`
	// Anthropic integration configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_casb_integration#anthropic ZeroTrustCasbIntegration#anthropic}
	Anthropic *ZeroTrustCasbIntegrationAnthropic `field:"optional" json:"anthropic" yaml:"anthropic"`
	// AWS integration configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_casb_integration#aws ZeroTrustCasbIntegration#aws}
	Aws *ZeroTrustCasbIntegrationAws `field:"optional" json:"aws" yaml:"aws"`
	// Box integration configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_casb_integration#box ZeroTrustCasbIntegration#box}
	Box *ZeroTrustCasbIntegrationBox `field:"optional" json:"box" yaml:"box"`
	// DLP profile IDs to associate with the integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_casb_integration#dlp_profiles ZeroTrustCasbIntegration#dlp_profiles}
	DlpProfiles *[]*string `field:"optional" json:"dlpProfiles" yaml:"dlpProfiles"`
	// Google Cloud Platform integration configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_casb_integration#google_cloud_platform ZeroTrustCasbIntegration#google_cloud_platform}
	GoogleCloudPlatform *ZeroTrustCasbIntegrationGoogleCloudPlatform `field:"optional" json:"googleCloudPlatform" yaml:"googleCloudPlatform"`
	// Google Workspace integration configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_casb_integration#google_workspace ZeroTrustCasbIntegration#google_workspace}
	GoogleWorkspace *ZeroTrustCasbIntegrationGoogleWorkspace `field:"optional" json:"googleWorkspace" yaml:"googleWorkspace"`
	// OpenAI integration configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_casb_integration#openai ZeroTrustCasbIntegration#openai}
	Openai *ZeroTrustCasbIntegrationOpenai `field:"optional" json:"openai" yaml:"openai"`
	// Permission scopes granted to the integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_casb_integration#permissions ZeroTrustCasbIntegration#permissions}
	Permissions *[]*string `field:"optional" json:"permissions" yaml:"permissions"`
	// Use cases to enroll the integration in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_casb_integration#use_cases ZeroTrustCasbIntegration#use_cases}
	UseCases *[]*string `field:"optional" json:"useCases" yaml:"useCases"`
}

