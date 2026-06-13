// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrusttunnelwarpconnector

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ZeroTrustTunnelWarpConnectorConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// A user-friendly name for a tunnel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/resources/zero_trust_tunnel_warp_connector#name ZeroTrustTunnelWarpConnector#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Cloudflare account ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/resources/zero_trust_tunnel_warp_connector#account_id ZeroTrustTunnelWarpConnector#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Indicates that the tunnel will be created to be highly available. If omitted, defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/resources/zero_trust_tunnel_warp_connector#ha ZeroTrustTunnelWarpConnector#ha}
	Ha interface{} `field:"optional" json:"ha" yaml:"ha"`
	// Sets the password required to run a locally-managed tunnel.
	//
	// Must be at least 32 bytes and encoded as a base64 string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/resources/zero_trust_tunnel_warp_connector#tunnel_secret ZeroTrustTunnelWarpConnector#tunnel_secret}
	TunnelSecret *string `field:"optional" json:"tunnelSecret" yaml:"tunnelSecret"`
}

