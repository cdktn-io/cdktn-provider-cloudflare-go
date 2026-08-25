// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aisearchinstance


type AiSearchInstanceSourceParamsWebCrawler struct {
	// Options for parse_type 'discover', where Browser Run discovers URLs by link following and sitemaps. Ignored for 'sitemap'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/ai_search_instance#discover_options AiSearchInstance#discover_options}
	DiscoverOptions *AiSearchInstanceSourceParamsWebCrawlerDiscoverOptions `field:"optional" json:"discoverOptions" yaml:"discoverOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/ai_search_instance#parse_options AiSearchInstance#parse_options}.
	ParseOptions *AiSearchInstanceSourceParamsWebCrawlerParseOptions `field:"optional" json:"parseOptions" yaml:"parseOptions"`
	// How URLs are discovered.
	//
	// 'sitemap' reads XML sitemaps; 'discover' follows links recursively and requires the source to be a Verified zone on this account.
	// Available values: "sitemap", "discover".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/ai_search_instance#parse_type AiSearchInstance#parse_type}
	ParseType *string `field:"optional" json:"parseType" yaml:"parseType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/ai_search_instance#store_options AiSearchInstance#store_options}.
	StoreOptions *AiSearchInstanceSourceParamsWebCrawlerStoreOptions `field:"optional" json:"storeOptions" yaml:"storeOptions"`
}

