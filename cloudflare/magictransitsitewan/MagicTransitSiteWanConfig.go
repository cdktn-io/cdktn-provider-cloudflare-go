// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package magictransitsitewan

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MagicTransitSiteWanConfig struct {
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
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/magic_transit_site_wan#account_id MagicTransitSiteWan#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/magic_transit_site_wan#physport MagicTransitSiteWan#physport}.
	Physport *float64 `field:"required" json:"physport" yaml:"physport"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/magic_transit_site_wan#site_id MagicTransitSiteWan#site_id}
	SiteId *string `field:"required" json:"siteId" yaml:"siteId"`
	// Magic WAN health check rate for tunnels created on this link.
	//
	// The default value is `mid`.
	// Available values: "low", "mid", "high".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/magic_transit_site_wan#health_check_rate MagicTransitSiteWan#health_check_rate}
	HealthCheckRate *string `field:"optional" json:"healthCheckRate" yaml:"healthCheckRate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/magic_transit_site_wan#load_balance_inner_flows MagicTransitSiteWan#load_balance_inner_flows}.
	LoadBalanceInnerFlows interface{} `field:"optional" json:"loadBalanceInnerFlows" yaml:"loadBalanceInnerFlows"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/magic_transit_site_wan#name MagicTransitSiteWan#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/magic_transit_site_wan#priority MagicTransitSiteWan#priority}.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// (optional) if omitted, use DHCP. Submit secondary_address when site is in high availability mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/magic_transit_site_wan#static_addressing MagicTransitSiteWan#static_addressing}
	StaticAddressing *MagicTransitSiteWanStaticAddressing `field:"optional" json:"staticAddressing" yaml:"staticAddressing"`
	// VLAN ID. Use zero for untagged.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/magic_transit_site_wan#vlan_tag MagicTransitSiteWan#vlan_tag}
	VlanTag *float64 `field:"optional" json:"vlanTag" yaml:"vlanTag"`
}

