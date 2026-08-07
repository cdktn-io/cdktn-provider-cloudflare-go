// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aisearchinstance


type AiSearchInstancePublicEndpointParams struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/resources/ai_search_instance#authorized_hosts AiSearchInstance#authorized_hosts}.
	AuthorizedHosts *[]*string `field:"optional" json:"authorizedHosts" yaml:"authorizedHosts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/resources/ai_search_instance#chat_completions_endpoint AiSearchInstance#chat_completions_endpoint}.
	ChatCompletionsEndpoint *AiSearchInstancePublicEndpointParamsChatCompletionsEndpoint `field:"optional" json:"chatCompletionsEndpoint" yaml:"chatCompletionsEndpoint"`
	// Custom domain hostnames that alias this public endpoint.
	//
	// GET and create responses return the current set; on update (PUT) this field is only echoed back when supplied in the request body, otherwise it is null (omit it to leave domains unchanged).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/resources/ai_search_instance#custom_domains AiSearchInstance#custom_domains}
	CustomDomains *[]*string `field:"optional" json:"customDomains" yaml:"customDomains"`
	// When false, the instance is reachable only via a registered custom domain and the default <public_endpoint_id>.search.ai.cloudflare.com host returns 404. Requires at least one custom domain. Defaults to true. public_endpoint_params is replaced wholesale on update, so resend default_domain_enabled on every update to keep the default host off — omitting it resets to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/resources/ai_search_instance#default_domain_enabled AiSearchInstance#default_domain_enabled}
	DefaultDomainEnabled interface{} `field:"optional" json:"defaultDomainEnabled" yaml:"defaultDomainEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/resources/ai_search_instance#enabled AiSearchInstance#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/resources/ai_search_instance#mcp AiSearchInstance#mcp}.
	Mcp *AiSearchInstancePublicEndpointParamsMcp `field:"optional" json:"mcp" yaml:"mcp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/resources/ai_search_instance#rate_limit AiSearchInstance#rate_limit}.
	RateLimit *AiSearchInstancePublicEndpointParamsRateLimit `field:"optional" json:"rateLimit" yaml:"rateLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/resources/ai_search_instance#search_endpoint AiSearchInstance#search_endpoint}.
	SearchEndpoint *AiSearchInstancePublicEndpointParamsSearchEndpoint `field:"optional" json:"searchEndpoint" yaml:"searchEndpoint"`
}

