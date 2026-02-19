// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zonesetting

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ZoneSettingConfig struct {
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
	// Setting name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zone_setting#setting_id ZoneSetting#setting_id}
	SettingId *string `field:"required" json:"settingId" yaml:"settingId"`
	// Current value of the zone setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zone_setting#value ZoneSetting#value}
	Value *map[string]interface{} `field:"required" json:"value" yaml:"value"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zone_setting#zone_id ZoneSetting#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// ssl-recommender enrollment setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zone_setting#enabled ZoneSetting#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

