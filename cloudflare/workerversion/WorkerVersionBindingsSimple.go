// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workerversion


type WorkerVersionBindingsSimple struct {
	// The limit (requests per period).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/worker_version#limit WorkerVersion#limit}
	Limit *float64 `field:"required" json:"limit" yaml:"limit"`
	// The period in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/worker_version#period WorkerVersion#period}
	Period *float64 `field:"required" json:"period" yaml:"period"`
	// Duration in seconds to apply the mitigation action after the rate limit is exceeded.
	//
	// Valid values are 0 (disabled), 10, or multiples of 60 up to 86400. Must be greater than or equal to the period when non-zero.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/worker_version#mitigation_timeout WorkerVersion#mitigation_timeout}
	MitigationTimeout *float64 `field:"optional" json:"mitigationTimeout" yaml:"mitigationTimeout"`
}

