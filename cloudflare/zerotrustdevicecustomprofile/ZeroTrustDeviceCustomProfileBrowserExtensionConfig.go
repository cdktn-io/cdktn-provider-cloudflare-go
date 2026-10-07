// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustdevicecustomprofile


type ZeroTrustDeviceCustomProfileBrowserExtensionConfig struct {
	// Whether the user may disable the browser extension proxy. Available values: "unlocked", "locked".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_device_custom_profile#proxy_control ZeroTrustDeviceCustomProfile#proxy_control}
	ProxyControl *string `field:"required" json:"proxyControl" yaml:"proxyControl"`
	// Whether the browser extension proxy is active.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_device_custom_profile#proxy_enabled ZeroTrustDeviceCustomProfile#proxy_enabled}
	ProxyEnabled interface{} `field:"required" json:"proxyEnabled" yaml:"proxyEnabled"`
}

