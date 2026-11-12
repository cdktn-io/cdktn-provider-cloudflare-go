// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package logexplorerdataset


type LogExplorerDatasetFields struct {
	// Whether the API includes this field in log ingest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/log_explorer_dataset#enabled LogExplorerDataset#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Field name. Names are matched without regard to case.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/log_explorer_dataset#name LogExplorerDataset#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

