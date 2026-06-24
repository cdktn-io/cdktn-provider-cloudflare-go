// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustdevicecustomprofile


type ZeroTrustDeviceCustomProfileVirtualNetworks struct {
	// List of virtual network IDs the device is allowed to access.
	//
	// When virtual_networks is set, at least one entry is required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.1/docs/resources/zero_trust_device_custom_profile#allowed ZeroTrustDeviceCustomProfile#allowed}
	Allowed *[]*string `field:"required" json:"allowed" yaml:"allowed"`
	// The default virtual network ID. Must be included in the `allowed` list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.1/docs/resources/zero_trust_device_custom_profile#default ZeroTrustDeviceCustomProfile#default}
	Default *string `field:"required" json:"default" yaml:"default"`
}

