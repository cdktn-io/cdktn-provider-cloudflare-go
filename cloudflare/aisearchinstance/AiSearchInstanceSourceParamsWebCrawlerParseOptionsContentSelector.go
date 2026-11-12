// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aisearchinstance


type AiSearchInstanceSourceParamsWebCrawlerParseOptionsContentSelector struct {
	// Glob pattern to match against the page URL path.
	//
	// Uses standard glob syntax: * matches within a segment, ** crosses directories.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/ai_search_instance#path AiSearchInstance#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// CSS selector to extract content from pages matching the path pattern.
	//
	// Disallowed characters: ;, `, $, {, }, \. Must target a single element; if multiple elements match, the crawler ignores the selector and uses the full page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/ai_search_instance#selector AiSearchInstance#selector}
	Selector *string `field:"required" json:"selector" yaml:"selector"`
}

