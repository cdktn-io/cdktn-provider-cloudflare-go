// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package listitem


type ListItemHostname struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/list_item#url_hostname ListItem#url_hostname}.
	UrlHostname *string `field:"required" json:"urlHostname" yaml:"urlHostname"`
	// Only applies to wildcard hostnames (e.g., *.example.com). When true (default), the rule blocks only subdomains. When false, the rule blocks both the root domain and subdomains.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/list_item#exclude_exact_hostname ListItem#exclude_exact_hostname}
	ExcludeExactHostname interface{} `field:"optional" json:"excludeExactHostname" yaml:"excludeExactHostname"`
}

