// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package regionalhostname

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RegionalHostnameConfig struct {
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
	// DNS hostname to be regionalized, must be a subdomain of the zone.
	//
	// Wildcards are supported for one level, e.g `*.example.com`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/regional_hostname#hostname RegionalHostname#hostname}
	Hostname *string `field:"required" json:"hostname" yaml:"hostname"`
	// Identifying key for the region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/regional_hostname#region_key RegionalHostname#region_key}
	RegionKey *string `field:"required" json:"regionKey" yaml:"regionKey"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/regional_hostname#zone_id RegionalHostname#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// Configure which routing method to use for the regional hostname.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/regional_hostname#routing RegionalHostname#routing}
	Routing *string `field:"optional" json:"routing" yaml:"routing"`
}

