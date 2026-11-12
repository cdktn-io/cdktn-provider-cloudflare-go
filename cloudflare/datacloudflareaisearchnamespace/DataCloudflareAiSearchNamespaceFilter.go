// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflareaisearchnamespace


type DataCloudflareAiSearchNamespaceFilter struct {
	// Filter namespaces whose name or description contains this string (case-insensitive).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/data-sources/ai_search_namespace#search DataCloudflareAiSearchNamespace#search}
	Search *string `field:"optional" json:"search" yaml:"search"`
}

