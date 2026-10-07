// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zonetracingrules


type ZoneTracingRulesRules struct {
	// Available values: "set_trace_settings".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zone_tracing_rules#action ZoneTracingRules#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zone_tracing_rules#action_parameters ZoneTracingRules#action_parameters}.
	ActionParameters *ZoneTracingRulesRulesActionParameters `field:"required" json:"actionParameters" yaml:"actionParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zone_tracing_rules#description ZoneTracingRules#description}.
	Description *string `field:"required" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zone_tracing_rules#enabled ZoneTracingRules#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// A Rules language expression that selects requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zone_tracing_rules#expression ZoneTracingRules#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
}

