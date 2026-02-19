// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package urlnormalizationsettings

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type UrlNormalizationSettingsConfig struct {
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
	// The scope of the URL normalization. Available values: "incoming", "both", "none".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/url_normalization_settings#scope UrlNormalizationSettings#scope}
	Scope *string `field:"required" json:"scope" yaml:"scope"`
	// The type of URL normalization performed by Cloudflare. Available values: "cloudflare", "rfc3986".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/url_normalization_settings#type UrlNormalizationSettings#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The unique ID of the zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/url_normalization_settings#zone_id UrlNormalizationSettings#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
}

