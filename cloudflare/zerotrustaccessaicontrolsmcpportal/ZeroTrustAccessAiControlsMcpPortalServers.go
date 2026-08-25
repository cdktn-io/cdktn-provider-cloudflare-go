// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessaicontrolsmcpportal


type ZeroTrustAccessAiControlsMcpPortalServers struct {
	// Unique identifier for the MCP server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/zero_trust_access_ai_controls_mcp_portal#server_id ZeroTrustAccessAiControlsMcpPortal#server_id}
	ServerId *string `field:"required" json:"serverId" yaml:"serverId"`
	// Disable this server by default for clients connecting through the portal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/zero_trust_access_ai_controls_mcp_portal#default_disabled ZeroTrustAccessAiControlsMcpPortal#default_disabled}
	DefaultDisabled interface{} `field:"optional" json:"defaultDisabled" yaml:"defaultDisabled"`
	// Use end-user OAuth credentials when connecting this server to the portal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/zero_trust_access_ai_controls_mcp_portal#on_behalf ZeroTrustAccessAiControlsMcpPortal#on_behalf}
	OnBehalf interface{} `field:"optional" json:"onBehalf" yaml:"onBehalf"`
	// Portal-specific prompt overrides.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/zero_trust_access_ai_controls_mcp_portal#updated_prompts ZeroTrustAccessAiControlsMcpPortal#updated_prompts}
	UpdatedPrompts interface{} `field:"optional" json:"updatedPrompts" yaml:"updatedPrompts"`
	// Portal-specific tool overrides.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/zero_trust_access_ai_controls_mcp_portal#updated_tools ZeroTrustAccessAiControlsMcpPortal#updated_tools}
	UpdatedTools interface{} `field:"optional" json:"updatedTools" yaml:"updatedTools"`
}

