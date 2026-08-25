// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ctalerting

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CtAlertingConfig struct {
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
	// Whether CT alerting is enabled for the zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/ct_alerting#enabled CtAlerting#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/ct_alerting#zone_id CtAlerting#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// Email addresses that receive CT alert notifications for the zone.
	//
	// A maximum of 100 addresses may be configured. Each address must be a valid RFC 5322 email address and must not contain a comma.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/ct_alerting#emails CtAlerting#emails}
	Emails *[]*string `field:"optional" json:"emails" yaml:"emails"`
}

