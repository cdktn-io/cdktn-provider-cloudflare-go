// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrusttunnelwarpconnectorconfig


type ZeroTrustTunnelWarpConnectorConfigConfigVips struct {
	// Virtual IP address (IPv4 or IPv6).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.1/docs/resources/zero_trust_tunnel_warp_connector_config#address ZeroTrustTunnelWarpConnectorConfigA#address}
	Address *string `field:"required" json:"address" yaml:"address"`
}

