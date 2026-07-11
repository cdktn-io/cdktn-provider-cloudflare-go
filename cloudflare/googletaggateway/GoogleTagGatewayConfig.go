// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googletaggateway

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleTagGatewayConfig struct {
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
	// Enables or disables Google Tag Gateway for this zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/resources/google_tag_gateway#enabled GoogleTagGateway#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Specifies the endpoint path for proxying Google Tag Manager requests.
	//
	// Use an absolute path starting with '/', with no nested paths and alphanumeric characters only (e.g. /metrics).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/resources/google_tag_gateway#endpoint GoogleTagGateway#endpoint}
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// Hides the original client IP address from Google when enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/resources/google_tag_gateway#hide_original_ip GoogleTagGateway#hide_original_ip}
	HideOriginalIp interface{} `field:"required" json:"hideOriginalIp" yaml:"hideOriginalIp"`
	// Specify the Google Tag Manager container or measurement ID (e.g. GTM-XXXXXXX or G-XXXXXXXXXX).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/resources/google_tag_gateway#measurement_id GoogleTagGateway#measurement_id}
	MeasurementId *string `field:"required" json:"measurementId" yaml:"measurementId"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/resources/google_tag_gateway#zone_id GoogleTagGateway#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// Set up the associated Google Tag on the zone automatically when enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/resources/google_tag_gateway#set_up_tag GoogleTagGateway#set_up_tag}
	SetUpTag interface{} `field:"optional" json:"setUpTag" yaml:"setUpTag"`
}

