// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aisearchnamespace


type AiSearchNamespacePublicEndpointParams struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/ai_search_namespace#authorized_hosts AiSearchNamespace#authorized_hosts}.
	AuthorizedHosts *[]*string `field:"optional" json:"authorizedHosts" yaml:"authorizedHosts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/ai_search_namespace#chat_completions_endpoint AiSearchNamespace#chat_completions_endpoint}.
	ChatCompletionsEndpoint *AiSearchNamespacePublicEndpointParamsChatCompletionsEndpoint `field:"optional" json:"chatCompletionsEndpoint" yaml:"chatCompletionsEndpoint"`
	// Custom domain hostnames that alias this public endpoint.
	//
	// GET and create responses return the current set; on update (PUT) this field is only echoed back when supplied in the request body, otherwise it is null (omit it to leave domains unchanged).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/ai_search_namespace#custom_domains AiSearchNamespace#custom_domains}
	CustomDomains *[]*string `field:"optional" json:"customDomains" yaml:"customDomains"`
	// When false, the instance is reachable only via a registered custom domain and the default <public_endpoint_id>.search.ai.cloudflare.com host returns 404. Requires at least one custom domain. Defaults to true. public_endpoint_params is replaced wholesale on update, so resend default_domain_enabled on every update to keep the default host off — omitting it resets to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/ai_search_namespace#default_domain_enabled AiSearchNamespace#default_domain_enabled}
	DefaultDomainEnabled interface{} `field:"optional" json:"defaultDomainEnabled" yaml:"defaultDomainEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/ai_search_namespace#enabled AiSearchNamespace#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Instance IDs exposed through the namespace public endpoint.
	//
	// Empty means nothing is searchable. Every ID must be an existing instance in this namespace, and the list cannot exceed the account's multi-instance search limit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/ai_search_namespace#instances_allowed AiSearchNamespace#instances_allowed}
	InstancesAllowed *[]*string `field:"optional" json:"instancesAllowed" yaml:"instancesAllowed"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/ai_search_namespace#mcp AiSearchNamespace#mcp}.
	Mcp *AiSearchNamespacePublicEndpointParamsMcp `field:"optional" json:"mcp" yaml:"mcp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/ai_search_namespace#rate_limit AiSearchNamespace#rate_limit}.
	RateLimit *AiSearchNamespacePublicEndpointParamsRateLimit `field:"optional" json:"rateLimit" yaml:"rateLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/ai_search_namespace#search_endpoint AiSearchNamespace#search_endpoint}.
	SearchEndpoint *AiSearchNamespacePublicEndpointParamsSearchEndpoint `field:"optional" json:"searchEndpoint" yaml:"searchEndpoint"`
}

