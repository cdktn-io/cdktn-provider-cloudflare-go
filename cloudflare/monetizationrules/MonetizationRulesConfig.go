// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package monetizationrules

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MonetizationRulesConfig struct {
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
	// Full desired Payment Required ruleset.
	//
	// An empty array clears all payment rules. The ruleset may contain at most 40 unique wallet addresses; address comparison is case-insensitive. Every address must pass wallet screening before deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.28.0/docs/resources/monetization_rules#rules MonetizationRules#rules}
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
	// The unique ID of the zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.28.0/docs/resources/monetization_rules#zone_id MonetizationRules#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
}

