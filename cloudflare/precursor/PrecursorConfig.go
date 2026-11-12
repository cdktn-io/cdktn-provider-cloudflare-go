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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/precursor#zone_id Precursor#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// Controls Precursor detection for the zone.
	//
	// `off` disables detection,
	// `monitor` records detections without affecting bot scores, and `arm`
	// allows detections to affect bot scores.
	// Available values: "off", "monitor", "arm".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/precursor#mode Precursor#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// The ordered list of Precursor rules for the zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/precursor#rules Precursor#rules}
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
	// Whether Precursor shows an interstitial when it takes action on a request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/precursor#show_interstitial Precursor#show_interstitial}
	ShowInterstitial interface{} `field:"optional" json:"showInterstitial" yaml:"showInterstitial"`
}

