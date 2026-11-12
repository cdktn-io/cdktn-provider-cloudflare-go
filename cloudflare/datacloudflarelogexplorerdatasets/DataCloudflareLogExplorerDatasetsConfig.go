// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflarelogexplorerdatasets

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataCloudflareLogExplorerDatasetsConfig struct {
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
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/data-sources/log_explorer_datasets#account_id DataCloudflareLogExplorerDatasets#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Set to true to include zone-scoped datasets belonging to this account. Only supported with account_id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/data-sources/log_explorer_datasets#include_zones DataCloudflareLogExplorerDatasets#include_zones}
	IncludeZones interface{} `field:"optional" json:"includeZones" yaml:"includeZones"`
	// Max items to fetch, default: 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/data-sources/log_explorer_datasets#max_items DataCloudflareLogExplorerDatasets#max_items}
	MaxItems *float64 `field:"optional" json:"maxItems" yaml:"maxItems"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/data-sources/log_explorer_datasets#zone_id DataCloudflareLogExplorerDatasets#zone_id}
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}

