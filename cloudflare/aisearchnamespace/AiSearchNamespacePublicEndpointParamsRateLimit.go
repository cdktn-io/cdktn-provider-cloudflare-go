// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aisearchnamespace


type AiSearchNamespacePublicEndpointParamsRateLimit struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/ai_search_namespace#period_ms AiSearchNamespace#period_ms}.
	PeriodMs *float64 `field:"optional" json:"periodMs" yaml:"periodMs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/ai_search_namespace#requests AiSearchNamespace#requests}.
	Requests *float64 `field:"optional" json:"requests" yaml:"requests"`
	// Available values: "fixed", "sliding".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/ai_search_namespace#technique AiSearchNamespace#technique}
	Technique *string `field:"optional" json:"technique" yaml:"technique"`
}

