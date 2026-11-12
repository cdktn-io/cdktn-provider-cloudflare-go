// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustcwsbasicauthsecret

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ZeroTrustCwsBasicAuthSecretConfig struct {
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
	// Cloudflare account ID (32-character hexadecimal string).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/zero_trust_cws_basic_auth_secret#account_id ZeroTrustCwsBasicAuthSecret#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// ID of the parent credential.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/zero_trust_cws_basic_auth_secret#credential_id ZeroTrustCwsBasicAuthSecret#credential_id}
	CredentialId *string `field:"required" json:"credentialId" yaml:"credentialId"`
	// Optional deactivation deadline (RFC 3339). Omit or null for no expiry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/zero_trust_cws_basic_auth_secret#expires_at ZeroTrustCwsBasicAuthSecret#expires_at}
	ExpiresAt *string `field:"optional" json:"expiresAt" yaml:"expiresAt"`
	// Password stored only as a one-way hash and never returned. Maximum 72 UTF-8 bytes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/zero_trust_cws_basic_auth_secret#password ZeroTrustCwsBasicAuthSecret#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
}

