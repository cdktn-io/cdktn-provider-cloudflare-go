// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aisearchinstance


type AiSearchInstanceRetrievalOptions struct {
	// Controls how keyword search terms are matched.
	//
	// exact_match requires all terms to appear (AND); fuzzy_match returns results containing any term (OR). Defaults to exact_match.
	// Available values: "exact_match", "fuzzy_match".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.18.0/docs/resources/ai_search_instance#keyword_match_mode AiSearchInstance#keyword_match_mode}
	KeywordMatchMode *string `field:"optional" json:"keywordMatchMode" yaml:"keywordMatchMode"`
}

