// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflareaisearchinstance


type DataCloudflareAiSearchInstanceFilter struct {
	// Return only instances with this exact Search for Agents hostname (case-insensitive).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/data-sources/ai_search_instance#hostname DataCloudflareAiSearchInstance#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// Field to order results by. Available values: "created_at".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/data-sources/ai_search_instance#order_by DataCloudflareAiSearchInstance#order_by}
	OrderBy *string `field:"optional" json:"orderBy" yaml:"orderBy"`
	// Order direction. Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/data-sources/ai_search_instance#order_by_direction DataCloudflareAiSearchInstance#order_by_direction}
	OrderByDirection *string `field:"optional" json:"orderByDirection" yaml:"orderByDirection"`
	// Filter instances whose id contains this string (case-insensitive).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/data-sources/ai_search_instance#search DataCloudflareAiSearchInstance#search}
	Search *string `field:"optional" json:"search" yaml:"search"`
}

