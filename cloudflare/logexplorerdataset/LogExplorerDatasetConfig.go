// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package logexplorerdataset

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LogExplorerDatasetConfig struct {
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
	// Dataset type name to create (e.g. `http_requests`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/log_explorer_dataset#dataset LogExplorerDataset#dataset}
	Dataset *string `field:"required" json:"dataset" yaml:"dataset"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/log_explorer_dataset#account_id LogExplorerDataset#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Set to `false` to allow deletion of this dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/log_explorer_dataset#deletion_protection LogExplorerDataset#deletion_protection}
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Whether to enable or disable log ingest for this dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/log_explorer_dataset#enabled LogExplorerDataset#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Controls which fields the API ingests.
	//
	// Defaults to all available
	// fields when absent.
	// Each dataset requires its own timestamp field to remain enabled (the
	// field name varies per dataset type); omitting it from an explicit
	// selection is rejected by the API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/log_explorer_dataset#fields LogExplorerDataset#fields}
	Fields interface{} `field:"optional" json:"fields" yaml:"fields"`
	// Optional Logpush filter predicate to restrict which events are ingested.
	//
	// Omitting this attribute means no filter. An empty string also means no filter.
	// Removing a configured filter sends an empty string to clear it in the API.
	// The provider stores missing or null API filters as empty strings in state.
	// See [Logpush filters](https://developers.cloudflare.com/logs/reference/filters/)
	// for syntax and examples.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/log_explorer_dataset#filter LogExplorerDataset#filter}
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
	// Whether to filter attack traffic from the Logpush job.
	//
	// The API defaults to true on creation for http_requests, firewall_events, and network_analytics_logs. If omitted on update, the existing setting is kept. The API omits this setting for datasets that do not support it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/log_explorer_dataset#filter_attack_traffic LogExplorerDataset#filter_attack_traffic}
	FilterAttackTraffic interface{} `field:"optional" json:"filterAttackTraffic" yaml:"filterAttackTraffic"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/log_explorer_dataset#zone_id LogExplorerDataset#zone_id}
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}

