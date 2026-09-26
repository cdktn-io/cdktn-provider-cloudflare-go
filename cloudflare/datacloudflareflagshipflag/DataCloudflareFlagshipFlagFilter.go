// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflareflagshipflag


type DataCloudflareFlagshipFlagFilter struct {
	// Max items to return (1–200).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/data-sources/flagship_flag#limit DataCloudflareFlagshipFlag#limit}
	Limit *string `field:"optional" json:"limit" yaml:"limit"`
}

