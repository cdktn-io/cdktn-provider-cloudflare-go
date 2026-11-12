// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package r2datacatalog


type R2DataCatalogCredential struct {
	// Provides the Cloudflare API token for accessing R2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/r2_data_catalog#token R2DataCatalog#token}
	Token *string `field:"required" json:"token" yaml:"token"`
}

