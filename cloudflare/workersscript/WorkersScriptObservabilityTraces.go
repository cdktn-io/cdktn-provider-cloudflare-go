// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workersscript


type WorkersScriptObservabilityTraces struct {
	// A list of destinations where traces will be exported to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.0/docs/resources/workers_script#destinations WorkersScript#destinations}
	Destinations *[]*string `field:"optional" json:"destinations" yaml:"destinations"`
	// Whether traces are enabled for the Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.0/docs/resources/workers_script#enabled WorkersScript#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The sampling rate for traces. From 0 to 1 (1 = 100%, 0.1 = 10%). Default is 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.0/docs/resources/workers_script#head_sampling_rate WorkersScript#head_sampling_rate}
	HeadSamplingRate *float64 `field:"optional" json:"headSamplingRate" yaml:"headSamplingRate"`
	// Whether trace persistence is enabled for the Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.0/docs/resources/workers_script#persist WorkersScript#persist}
	Persist interface{} `field:"optional" json:"persist" yaml:"persist"`
	// Controls how inbound trace context (traceparent/tracestate) headers on incoming requests are handled.
	//
	// "authenticated" (default) honors inbound trace context only when accompanied by a valid trace auth token. "accept" unconditionally accepts inbound trace context. Requires the trace propagation feature to be enabled.
	// Available values: "authenticated", "accept".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.0/docs/resources/workers_script#propagation_policy WorkersScript#propagation_policy}
	PropagationPolicy *string `field:"optional" json:"propagationPolicy" yaml:"propagationPolicy"`
}

