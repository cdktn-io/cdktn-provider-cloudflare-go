// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package firewallrule

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type FirewallRuleConfig struct {
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
	// The action to perform when the threshold of matched traffic within the configured period is exceeded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/firewall_rule#action FirewallRule#action}
	Action *FirewallRuleAction `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/firewall_rule#filter FirewallRule#filter}.
	Filter *FirewallRuleFilter `field:"required" json:"filter" yaml:"filter"`
	// Defines an identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/firewall_rule#zone_id FirewallRule#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
}

