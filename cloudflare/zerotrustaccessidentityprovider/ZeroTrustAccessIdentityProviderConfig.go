// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessidentityprovider

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ZeroTrustAccessIdentityProviderConfig struct {
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
	// The configuration parameters for the identity provider.
	//
	// To view the required parameters for a specific provider, refer to our [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/zero_trust_access_identity_provider#config ZeroTrustAccessIdentityProvider#config}
	Config *ZeroTrustAccessIdentityProviderConfigA `field:"required" json:"config" yaml:"config"`
	// The name of the identity provider, shown to users on the login page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/zero_trust_access_identity_provider#name ZeroTrustAccessIdentityProvider#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of identity provider.
	//
	// To determine the value for a specific provider, refer to our [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	// Available values: "onetimepin", "azureAD", "saml", "centrify", "facebook", "github", "google-apps", "google", "linkedin", "oidc", "okta", "onelogin", "pingone", "yandex", "cloudflare".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/zero_trust_access_identity_provider#type ZeroTrustAccessIdentityProvider#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/zero_trust_access_identity_provider#account_id ZeroTrustAccessIdentityProvider#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Indicates that the identity provider is immutable and cannot be updated or deleted via the API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/zero_trust_access_identity_provider#read_only ZeroTrustAccessIdentityProvider#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// The UID of the SAML encryption certificate set assigned to this Identity Provider.
	//
	// Only present for SAML identity providers with encryption configured.
	// Create a certificate set via POST to `/identity_providers/{id}/saml_certificate`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/zero_trust_access_identity_provider#saml_certificate_set_id ZeroTrustAccessIdentityProvider#saml_certificate_set_id}
	SamlCertificateSetId *string `field:"optional" json:"samlCertificateSetId" yaml:"samlCertificateSetId"`
	// The configuration settings for enabling a System for Cross-Domain Identity Management (SCIM) with the identity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/zero_trust_access_identity_provider#scim_config ZeroTrustAccessIdentityProvider#scim_config}
	ScimConfig *ZeroTrustAccessIdentityProviderScimConfig `field:"optional" json:"scimConfig" yaml:"scimConfig"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/zero_trust_access_identity_provider#zone_id ZeroTrustAccessIdentityProvider#zone_id}
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}

