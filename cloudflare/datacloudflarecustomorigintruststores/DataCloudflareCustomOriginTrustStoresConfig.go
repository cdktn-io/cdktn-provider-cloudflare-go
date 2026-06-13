// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflarecustomorigintruststores

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataCloudflareCustomOriginTrustStoresConfig struct {
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
	// Limit to the number of records returned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/data-sources/custom_origin_trust_stores#limit DataCloudflareCustomOriginTrustStores#limit}
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// Max items to fetch, default: 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/data-sources/custom_origin_trust_stores#max_items DataCloudflareCustomOriginTrustStores#max_items}
	MaxItems *float64 `field:"optional" json:"maxItems" yaml:"maxItems"`
	// Offset the results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/data-sources/custom_origin_trust_stores#offset DataCloudflareCustomOriginTrustStores#offset}
	Offset *float64 `field:"optional" json:"offset" yaml:"offset"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/data-sources/custom_origin_trust_stores#zone_id DataCloudflareCustomOriginTrustStores#zone_id}
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}

