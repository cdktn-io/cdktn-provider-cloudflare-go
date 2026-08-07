// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrusttunnelwarpconnectorconfig


type ZeroTrustTunnelWarpConnectorConfigConfig struct {
	// Floating Network Resource ID — the secondary ENI that is moved between nodes on failover.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/resources/zero_trust_tunnel_warp_connector_config#fnr_id ZeroTrustTunnelWarpConnectorConfigA#fnr_id}
	FnrId *string `field:"optional" json:"fnrId" yaml:"fnrId"`
	// VIPs to assign on the CloudflareWARP interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/resources/zero_trust_tunnel_warp_connector_config#vips ZeroTrustTunnelWarpConnectorConfigA#vips}
	Vips interface{} `field:"optional" json:"vips" yaml:"vips"`
	// VIPs to clean up on demotion or version drift.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/resources/zero_trust_tunnel_warp_connector_config#vips_previous ZeroTrustTunnelWarpConnectorConfigA#vips_previous}
	VipsPrevious interface{} `field:"optional" json:"vipsPrevious" yaml:"vipsPrevious"`
}

