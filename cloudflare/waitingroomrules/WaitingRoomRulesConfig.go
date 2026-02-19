// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package waitingroomrules

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WaitingRoomRulesConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/waiting_room_rules#rules WaitingRoomRules#rules}.
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/waiting_room_rules#waiting_room_id WaitingRoomRules#waiting_room_id}.
	WaitingRoomId *string `field:"required" json:"waitingRoomId" yaml:"waitingRoomId"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/waiting_room_rules#zone_id WaitingRoomRules#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
}

