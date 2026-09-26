// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package worker


type WorkerPreviewsBaseConfigObservability struct {
	// Whether observability is enabled for the Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/resources/worker#enabled Worker#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The sampling rate for observability. From 0 to 1 (1 = 100%, 0.1 = 10%).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/resources/worker#head_sampling_rate Worker#head_sampling_rate}
	HeadSamplingRate *float64 `field:"optional" json:"headSamplingRate" yaml:"headSamplingRate"`
	// Real-time Issues settings for the Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/resources/worker#issues Worker#issues}
	Issues *WorkerPreviewsBaseConfigObservabilityIssues `field:"optional" json:"issues" yaml:"issues"`
	// Log settings for the Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/resources/worker#logs Worker#logs}
	Logs *WorkerPreviewsBaseConfigObservabilityLogs `field:"optional" json:"logs" yaml:"logs"`
	// Whether query strings are removed from request URLs in logs and traces.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/resources/worker#redact_query_string Worker#redact_query_string}
	RedactQueryString interface{} `field:"optional" json:"redactQueryString" yaml:"redactQueryString"`
	// Trace settings for the Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/resources/worker#traces Worker#traces}
	Traces *WorkerPreviewsBaseConfigObservabilityTraces `field:"optional" json:"traces" yaml:"traces"`
}

