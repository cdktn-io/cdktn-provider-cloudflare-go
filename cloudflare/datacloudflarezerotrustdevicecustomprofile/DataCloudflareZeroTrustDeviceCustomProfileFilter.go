// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezerotrustdevicecustomprofile


type DataCloudflareZeroTrustDeviceCustomProfileFilter struct {
	// Filter profiles by client type. When omitted, only WARP profiles are returned. Available values: "warp", "browser_extension".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/zero_trust_device_custom_profile#profile_type DataCloudflareZeroTrustDeviceCustomProfile#profile_type}
	ProfileType *string `field:"optional" json:"profileType" yaml:"profileType"`
}

