// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessapplication


type ZeroTrustAccessApplicationMfaConfig struct {
	// The authenticators allowed for MFA. Available values: "totp", "biometrics", "security_key".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/resources/zero_trust_access_application#allowed_authenticators ZeroTrustAccessApplication#allowed_authenticators}
	AllowedAuthenticators *[]*string `field:"optional" json:"allowedAuthenticators" yaml:"allowedAuthenticators"`
	// Whether MFA is disabled for this application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/resources/zero_trust_access_application#mfa_disabled ZeroTrustAccessApplication#mfa_disabled}
	MfaDisabled interface{} `field:"optional" json:"mfaDisabled" yaml:"mfaDisabled"`
	// How often a user will be forced to re-authenticate with MFA.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/resources/zero_trust_access_application#session_duration ZeroTrustAccessApplication#session_duration}
	SessionDuration *string `field:"optional" json:"sessionDuration" yaml:"sessionDuration"`
}

