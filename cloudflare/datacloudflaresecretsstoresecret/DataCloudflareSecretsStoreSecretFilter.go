// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflaresecretsstoresecret


type DataCloudflareSecretsStoreSecretFilter struct {
	// Direction to sort objects. Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/data-sources/secrets_store_secret#direction DataCloudflareSecretsStoreSecret#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Order secrets by values in the given field. Available values: "name", "comment", "created", "modified", "status".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/data-sources/secrets_store_secret#order DataCloudflareSecretsStoreSecret#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
	// Only secrets with the given scopes will be returned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/data-sources/secrets_store_secret#scopes DataCloudflareSecretsStoreSecret#scopes}
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
	// Search secrets using a filter string, filtering across name and comment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/data-sources/secrets_store_secret#search DataCloudflareSecretsStoreSecret#search}
	Search *string `field:"optional" json:"search" yaml:"search"`
}

