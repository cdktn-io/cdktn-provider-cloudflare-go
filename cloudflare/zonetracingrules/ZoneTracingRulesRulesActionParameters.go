// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zonetracingrules


type ZoneTracingRulesRulesActionParameters struct {
	// The ratio of requests sampled for tracing, from 0 to 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zone_tracing_rules#sampling_ratio ZoneTracingRules#sampling_ratio}
	SamplingRatio *float64 `field:"required" json:"samplingRatio" yaml:"samplingRatio"`
}

