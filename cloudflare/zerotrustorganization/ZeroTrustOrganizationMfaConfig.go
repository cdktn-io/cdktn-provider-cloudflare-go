// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustorganization


type ZeroTrustOrganizationMfaConfig struct {
	// Lists the MFA methods that users can authenticate with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.18.0/docs/resources/zero_trust_organization#allowed_authenticators ZeroTrustOrganization#allowed_authenticators}
	AllowedAuthenticators *[]*string `field:"optional" json:"allowedAuthenticators" yaml:"allowedAuthenticators"`
	// Defines the duration of an MFA session.
	//
	// Must be in minutes (m) or hours (h). Minimum: 0m. Maximum: 720h (30 days). Examples:`5m` or `24h`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.18.0/docs/resources/zero_trust_organization#session_duration ZeroTrustOrganization#session_duration}
	SessionDuration *string `field:"optional" json:"sessionDuration" yaml:"sessionDuration"`
}

