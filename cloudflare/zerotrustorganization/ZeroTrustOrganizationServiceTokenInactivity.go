// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustorganization


type ZeroTrustOrganizationServiceTokenInactivity struct {
	// The action applied to an inactive service token. Available values: "disable", "delete".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_organization#action ZeroTrustOrganization#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// Whether automatic enforcement for inactive service tokens is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_organization#enabled ZeroTrustOrganization#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The number of days a service token must be inactive before the configured action is applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_organization#inactivity_threshold_days ZeroTrustOrganization#inactivity_threshold_days}
	InactivityThresholdDays *float64 `field:"required" json:"inactivityThresholdDays" yaml:"inactivityThresholdDays"`
}

