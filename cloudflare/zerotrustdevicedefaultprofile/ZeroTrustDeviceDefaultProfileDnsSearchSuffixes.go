// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustdevicedefaultprofile


type ZeroTrustDeviceDefaultProfileDnsSearchSuffixes struct {
	// The DNS search suffix to append when resolving short hostnames.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.0/docs/resources/zero_trust_device_default_profile#suffix ZeroTrustDeviceDefaultProfile#suffix}
	Suffix *string `field:"required" json:"suffix" yaml:"suffix"`
	// A description of the DNS search suffix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.0/docs/resources/zero_trust_device_default_profile#description ZeroTrustDeviceDefaultProfile#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

