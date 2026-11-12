// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezerotrustcwspolicy


type DataCloudflareZeroTrustCwsPolicyFindOneBy struct {
	// Sort direction (default: asc) Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/data-sources/zero_trust_cws_policy#direction DataCloudflareZeroTrustCwsPolicy#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Filter by policy name (case-insensitive substring match).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/data-sources/zero_trust_cws_policy#name DataCloudflareZeroTrustCwsPolicy#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Field to order results by (default: created_at) Available values: "name", "created_at", "updated_at".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/data-sources/zero_trust_cws_policy#order DataCloudflareZeroTrustCwsPolicy#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
	// Filter by package path (case-sensitive literal substring match).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/data-sources/zero_trust_cws_policy#package_path DataCloudflareZeroTrustCwsPolicy#package_path}
	PackagePath *string `field:"optional" json:"packagePath" yaml:"packagePath"`
	// Filter to policies attached to (members of) this policy set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/data-sources/zero_trust_cws_policy#policy_set_id DataCloudflareZeroTrustCwsPolicy#policy_set_id}
	PolicySetId *string `field:"optional" json:"policySetId" yaml:"policySetId"`
}

