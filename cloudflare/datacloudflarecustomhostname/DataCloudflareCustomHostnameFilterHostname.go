// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflarecustomhostname


type DataCloudflareCustomHostnameFilterHostname struct {
	// Filters hostnames by a substring match on the hostname value.
	//
	// This parameter cannot be used with the 'id', 'hostname', 'hostname.exact', or 'hostname.startsWith' parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/data-sources/custom_hostname#contain DataCloudflareCustomHostname#contain}
	Contain *string `field:"optional" json:"contain" yaml:"contain"`
	// Fully qualified domain name to match against.
	//
	// This parameter cannot be used with the 'id', 'hostname', 'hostname.contain', or 'hostname.startsWith' parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/data-sources/custom_hostname#exact DataCloudflareCustomHostname#exact}
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
	// Filters hostnames by a prefix match on the hostname value.
	//
	// This parameter cannot be used with the 'id', 'hostname', 'hostname.exact', or 'hostname.contain' parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/data-sources/custom_hostname#starts_with DataCloudflareCustomHostname#starts_with}
	StartsWith *string `field:"optional" json:"startsWith" yaml:"startsWith"`
}

