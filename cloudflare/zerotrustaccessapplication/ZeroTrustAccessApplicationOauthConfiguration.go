// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessapplication


type ZeroTrustAccessApplicationOauthConfiguration struct {
	// Settings for OAuth dynamic client registration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.1/docs/resources/zero_trust_access_application#dynamic_client_registration ZeroTrustAccessApplication#dynamic_client_registration}
	DynamicClientRegistration *ZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistration `field:"optional" json:"dynamicClientRegistration" yaml:"dynamicClientRegistration"`
	// Whether the OAuth configuration is enabled for this application.
	//
	// When set to `false`, Access will not handle OAuth for this application. Defaults to `true` if omitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.1/docs/resources/zero_trust_access_application#enabled ZeroTrustAccessApplication#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Settings for OAuth grant behavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.1/docs/resources/zero_trust_access_application#grant ZeroTrustAccessApplication#grant}
	Grant *ZeroTrustAccessApplicationOauthConfigurationGrant `field:"optional" json:"grant" yaml:"grant"`
}

