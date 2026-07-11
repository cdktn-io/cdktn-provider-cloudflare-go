// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflaresecretsstoresecrets

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataCloudflareSecretsStoreSecretsConfig struct {
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
	// Account Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/data-sources/secrets_store_secrets#account_id DataCloudflareSecretsStoreSecrets#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Store Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/data-sources/secrets_store_secrets#store_id DataCloudflareSecretsStoreSecrets#store_id}
	StoreId *string `field:"required" json:"storeId" yaml:"storeId"`
	// Direction to sort objects Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/data-sources/secrets_store_secrets#direction DataCloudflareSecretsStoreSecrets#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Max items to fetch, default: 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/data-sources/secrets_store_secrets#max_items DataCloudflareSecretsStoreSecrets#max_items}
	MaxItems *float64 `field:"optional" json:"maxItems" yaml:"maxItems"`
	// Order secrets by values in the given field Available values: "name", "comment", "created", "modified", "status".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/data-sources/secrets_store_secrets#order DataCloudflareSecretsStoreSecrets#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
	// Only secrets with the given scopes will be returned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/data-sources/secrets_store_secrets#scopes DataCloudflareSecretsStoreSecrets#scopes}
	Scopes interface{} `field:"optional" json:"scopes" yaml:"scopes"`
	// Search secrets using a filter string, filtering across name and comment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/data-sources/secrets_store_secrets#search DataCloudflareSecretsStoreSecrets#search}
	Search *string `field:"optional" json:"search" yaml:"search"`
}

