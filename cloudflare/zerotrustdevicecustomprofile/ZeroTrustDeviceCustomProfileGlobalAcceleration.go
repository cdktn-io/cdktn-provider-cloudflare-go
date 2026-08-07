// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustdevicecustomprofile


type ZeroTrustDeviceCustomProfileGlobalAcceleration struct {
	// IP:port entries for the API endpoints.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/resources/zero_trust_device_custom_profile#api_endpoints ZeroTrustDeviceCustomProfile#api_endpoints}
	ApiEndpoints *[]*string `field:"required" json:"apiEndpoints" yaml:"apiEndpoints"`
	// Global acceleration settings are used only when "enabled".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/resources/zero_trust_device_custom_profile#enabled ZeroTrustDeviceCustomProfile#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// IP:port entries for the MASQUE tunnel endpoints. Either wireguard_endpoints or masque_endpoints must be provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/resources/zero_trust_device_custom_profile#masque_endpoints ZeroTrustDeviceCustomProfile#masque_endpoints}
	MasqueEndpoints *[]*string `field:"required" json:"masqueEndpoints" yaml:"masqueEndpoints"`
	// IP:port entries for the WireGuard tunnel endpoints. Either wireguard_endpoints or masque_endpoints must be provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/resources/zero_trust_device_custom_profile#wireguard_endpoints ZeroTrustDeviceCustomProfile#wireguard_endpoints}
	WireguardEndpoints *[]*string `field:"required" json:"wireguardEndpoints" yaml:"wireguardEndpoints"`
}

