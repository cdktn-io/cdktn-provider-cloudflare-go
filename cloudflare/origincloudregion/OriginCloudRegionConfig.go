// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package origincloudregion

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OriginCloudRegionConfig struct {
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
	// Origin IP address (IPv4 or IPv6).
	//
	// For the single PUT endpoint (`PUT /origin/cloud_regions/{origin_ip}`), this field must match the path parameter or the request will be rejected with a 400 error. For the batch PUT endpoint, this field identifies which mapping to upsert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/origin_cloud_region#origin_ip OriginCloudRegion#origin_ip}
	OriginIp *string `field:"required" json:"originIp" yaml:"originIp"`
	// Cloud vendor region identifier. Must be a valid region for the specified vendor as returned by the supported_regions endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/origin_cloud_region#region OriginCloudRegion#region}
	Region *string `field:"required" json:"region" yaml:"region"`
	// Cloud vendor hosting the origin. Must be one of the supported vendors. Available values: "aws", "azure", "gcp", "oci".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/origin_cloud_region#vendor OriginCloudRegion#vendor}
	Vendor *string `field:"required" json:"vendor" yaml:"vendor"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/origin_cloud_region#zone_id OriginCloudRegion#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
}

