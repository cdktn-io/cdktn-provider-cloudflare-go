// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflarecustomhostnames


type DataCloudflareCustomHostnamesHostname struct {
	// Filters hostnames by a substring match on the hostname value. This parameter cannot be used with the 'id' parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.0/docs/data-sources/custom_hostnames#contain DataCloudflareCustomHostnames#contain}
	Contain *string `field:"optional" json:"contain" yaml:"contain"`
}

