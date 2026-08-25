// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessaicontrolsmcpportal


type ZeroTrustAccessAiControlsMcpPortalServersUpdatedTools struct {
	// Name of the tool or prompt capability to override.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/zero_trust_access_ai_controls_mcp_portal#name ZeroTrustAccessAiControlsMcpPortal#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Custom name exposed for the capability.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/zero_trust_access_ai_controls_mcp_portal#alias ZeroTrustAccessAiControlsMcpPortal#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Custom description exposed for the capability.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/zero_trust_access_ai_controls_mcp_portal#description ZeroTrustAccessAiControlsMcpPortal#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether the capability is available through the MCP server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/zero_trust_access_ai_controls_mcp_portal#enabled ZeroTrustAccessAiControlsMcpPortal#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

