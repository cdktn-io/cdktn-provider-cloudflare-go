// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessaicontrolsmcpportal

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ZeroTrustAccessAiControlsMcpPortalConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/zero_trust_access_ai_controls_mcp_portal#account_id ZeroTrustAccessAiControlsMcpPortal#account_id}.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Hostname where the MCP portal is available.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/zero_trust_access_ai_controls_mcp_portal#hostname ZeroTrustAccessAiControlsMcpPortal#hostname}
	Hostname *string `field:"required" json:"hostname" yaml:"hostname"`
	// Unique identifier for the MCP portal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/zero_trust_access_ai_controls_mcp_portal#id ZeroTrustAccessAiControlsMcpPortal#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Display name for the MCP portal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/zero_trust_access_ai_controls_mcp_portal#name ZeroTrustAccessAiControlsMcpPortal#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Deprecated: use `code_mode` for new integrations.
	//
	// `true` maps to any non-off Code Mode policy; `false` maps to `code_mode: off`. If both fields are sent, they must be consistent or the request returns a 400.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/zero_trust_access_ai_controls_mcp_portal#allow_code_mode ZeroTrustAccessAiControlsMcpPortal#allow_code_mode}
	AllowCodeMode interface{} `field:"optional" json:"allowCodeMode" yaml:"allowCodeMode"`
	// Code Mode policy for this portal.
	//
	// `off`: Code Mode is unavailable; query parameters are ignored. `opt_in`: Code Mode is off by default; clients turn it on with `?codemode=search_and_execute`. `default_on`: Code Mode is on by default; clients can opt out with `?codemode=off`. `enforced`: Code Mode is always on; query parameters are ignored. Defaults to `opt_in` when omitted on create. If both `code_mode` and `allow_code_mode` are sent, they must be consistent or the request returns a 400.
	// Available values: "off", "opt_in", "default_on", "enforced".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/zero_trust_access_ai_controls_mcp_portal#code_mode ZeroTrustAccessAiControlsMcpPortal#code_mode}
	CodeMode *string `field:"optional" json:"codeMode" yaml:"codeMode"`
	// Optional description of the MCP portal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/zero_trust_access_ai_controls_mcp_portal#description ZeroTrustAccessAiControlsMcpPortal#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Route outbound MCP traffic through Zero Trust Secure Web Gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/zero_trust_access_ai_controls_mcp_portal#secure_web_gateway ZeroTrustAccessAiControlsMcpPortal#secure_web_gateway}
	SecureWebGateway interface{} `field:"optional" json:"secureWebGateway" yaml:"secureWebGateway"`
	// MCP servers attached to the portal and their portal-specific settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/zero_trust_access_ai_controls_mcp_portal#servers ZeroTrustAccessAiControlsMcpPortal#servers}
	Servers interface{} `field:"optional" json:"servers" yaml:"servers"`
}

