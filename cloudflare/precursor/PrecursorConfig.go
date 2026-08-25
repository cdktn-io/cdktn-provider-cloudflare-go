// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package precursor

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PrecursorConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/precursor#zone_id Precursor#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// The zone-level Precursor enforcement mode applied to requests that do not match a more specific enforcement rule.
	//
	// Available values: "off", "min-friction", "max-security".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/precursor#default_mode Precursor#default_mode}
	DefaultMode *string `field:"optional" json:"defaultMode" yaml:"defaultMode"`
	// The ordered list of enforcement rules for the zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/precursor#enforcement_rules Precursor#enforcement_rules}
	EnforcementRules interface{} `field:"optional" json:"enforcementRules" yaml:"enforcementRules"`
}

