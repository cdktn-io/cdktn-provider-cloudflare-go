// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssoconnector

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SsoConnectorConfig struct {
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
	// Account identifier tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/sso_connector#account_id SsoConnector#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Email domain of the new SSO connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/sso_connector#email_domain SsoConnector#email_domain}
	EmailDomain *string `field:"required" json:"emailDomain" yaml:"emailDomain"`
	// Begin the verification process after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/sso_connector#begin_verification SsoConnector#begin_verification}
	BeginVerification interface{} `field:"optional" json:"beginVerification" yaml:"beginVerification"`
	// SSO Connector enabled state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/sso_connector#enabled SsoConnector#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Controls the display of FedRAMP language to the user during SSO login.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/sso_connector#use_fedramp_language SsoConnector#use_fedramp_language}
	UseFedrampLanguage interface{} `field:"optional" json:"useFedrampLanguage" yaml:"useFedrampLanguage"`
}

