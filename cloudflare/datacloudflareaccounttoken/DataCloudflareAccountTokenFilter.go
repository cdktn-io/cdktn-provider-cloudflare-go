// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflareaccounttoken


type DataCloudflareAccountTokenFilter struct {
	// Direction to order results. Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/data-sources/account_token#direction DataCloudflareAccountToken#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// When true, includes recently-expired tokens in the response.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/data-sources/account_token#include_expired DataCloudflareAccountToken#include_expired}
	IncludeExpired interface{} `field:"optional" json:"includeExpired" yaml:"includeExpired"`
}

