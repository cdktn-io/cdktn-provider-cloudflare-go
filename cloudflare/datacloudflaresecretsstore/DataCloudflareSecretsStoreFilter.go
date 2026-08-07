// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflaresecretsstore


type DataCloudflareSecretsStoreFilter struct {
	// Direction to sort objects. Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/data-sources/secrets_store#direction DataCloudflareSecretsStore#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Order stores by values in the given field. Available values: "name", "created", "modified".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/data-sources/secrets_store#order DataCloudflareSecretsStore#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
}

