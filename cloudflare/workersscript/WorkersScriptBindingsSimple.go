// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workersscript


type WorkersScriptBindingsSimple struct {
	// The rate limit value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.1/docs/resources/workers_script#limit WorkersScript#limit}
	Limit *float64 `field:"required" json:"limit" yaml:"limit"`
	// The rate limit period in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.1/docs/resources/workers_script#period WorkersScript#period}
	Period *float64 `field:"required" json:"period" yaml:"period"`
	// Duration in seconds to apply the mitigation action after the rate limit is exceeded.
	//
	// Valid values are 0 (disabled), 10, or multiples of 60 up to 86400. Must be greater than or equal to the period when non-zero.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.1/docs/resources/workers_script#mitigation_timeout WorkersScript#mitigation_timeout}
	MitigationTimeout *float64 `field:"optional" json:"mitigationTimeout" yaml:"mitigationTimeout"`
}

