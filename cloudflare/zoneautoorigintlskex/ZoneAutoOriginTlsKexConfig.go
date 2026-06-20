// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zoneautoorigintlskex

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ZoneAutoOriginTlsKexConfig struct {
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
	// Controls enablement of Auto-Origin TLS KEX selection for the zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.0/docs/resources/zone_auto_origin_tls_kex#enabled ZoneAutoOriginTlsKex#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.0/docs/resources/zone_auto_origin_tls_kex#zone_id ZoneAutoOriginTlsKex#zone_id}.
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
}

