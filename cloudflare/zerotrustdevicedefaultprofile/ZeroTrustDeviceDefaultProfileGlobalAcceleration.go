// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustdevicedefaultprofile


type ZeroTrustDeviceDefaultProfileGlobalAcceleration struct {
	// IP:port entries for the API endpoints.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/zero_trust_device_default_profile#api_endpoints ZeroTrustDeviceDefaultProfile#api_endpoints}
	ApiEndpoints *[]*string `field:"required" json:"apiEndpoints" yaml:"apiEndpoints"`
	// Global acceleration settings are used only when "enabled".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/zero_trust_device_default_profile#enabled ZeroTrustDeviceDefaultProfile#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// IP:port entries for the MASQUE tunnel endpoints. Either wireguard_endpoints or masque_endpoints must be provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/zero_trust_device_default_profile#masque_endpoints ZeroTrustDeviceDefaultProfile#masque_endpoints}
	MasqueEndpoints *[]*string `field:"required" json:"masqueEndpoints" yaml:"masqueEndpoints"`
	// IP:port entries for the WireGuard tunnel endpoints. Either wireguard_endpoints or masque_endpoints must be provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/zero_trust_device_default_profile#wireguard_endpoints ZeroTrustDeviceDefaultProfile#wireguard_endpoints}
	WireguardEndpoints *[]*string `field:"required" json:"wireguardEndpoints" yaml:"wireguardEndpoints"`
}

