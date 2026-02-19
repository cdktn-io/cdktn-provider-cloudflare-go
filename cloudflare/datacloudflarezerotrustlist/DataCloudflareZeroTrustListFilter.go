// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezerotrustlist


type DataCloudflareZeroTrustListFilter struct {
	// Specify the list type. Available values: "SERIAL", "URL", "DOMAIN", "EMAIL", "IP", "CATEGORY", "LOCATION", "DEVICE".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/data-sources/zero_trust_list#type DataCloudflareZeroTrustList#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

