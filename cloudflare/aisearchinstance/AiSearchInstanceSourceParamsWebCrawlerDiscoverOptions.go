// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aisearchinstance


type AiSearchInstanceSourceParamsWebCrawlerDiscoverOptions struct {
	// Maximum link-follow depth from the seed URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/ai_search_instance#depth AiSearchInstance#depth}
	Depth *float64 `field:"optional" json:"depth" yaml:"depth"`
	// Follow links that point outside the source domain.
	//
	// Must stay `false` — discover crawls are restricted to the zone you own.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/ai_search_instance#include_external_links AiSearchInstance#include_external_links}
	IncludeExternalLinks interface{} `field:"optional" json:"includeExternalLinks" yaml:"includeExternalLinks"`
	// Follow links to subdomains of the source host.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/ai_search_instance#include_subdomains AiSearchInstance#include_subdomains}
	IncludeSubdomains interface{} `field:"optional" json:"includeSubdomains" yaml:"includeSubdomains"`
	// Maximum number of pages to crawl (1-100000).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/ai_search_instance#limit AiSearchInstance#limit}
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// Maximum content age in seconds to accept (0–604800).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/ai_search_instance#max_age AiSearchInstance#max_age}
	MaxAge *float64 `field:"optional" json:"maxAge" yaml:"maxAge"`
	// Where the crawler looks for URLs: 'sitemaps' reads sitemap XML only, 'links' follows page links only, 'all' does both.
	//
	// Available values: "all", "sitemaps", "links".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/ai_search_instance#source AiSearchInstance#source}
	Source *string `field:"optional" json:"source" yaml:"source"`
}

