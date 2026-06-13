// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package secretsstoresecret

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SecretsStoreSecretConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/resources/secrets_store_secret#account_id SecretsStoreSecret#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The name of the secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/resources/secrets_store_secret#name SecretsStoreSecret#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The list of services that can use this secret.
	//
	// Valid values are `workers`, `ai_gateway`, `dex`, and `access`. Must be listed in alphabetical order.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/resources/secrets_store_secret#scopes SecretsStoreSecret#scopes}
	Scopes *[]*string `field:"required" json:"scopes" yaml:"scopes"`
	// Store Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/resources/secrets_store_secret#store_id SecretsStoreSecret#store_id}
	StoreId *string `field:"required" json:"storeId" yaml:"storeId"`
	// The value of the secret.
	//
	// Maximum 64 KiB (65,536 bytes). Note that this is 'write only' - no API response will provide this value, it is only used to create/modify secrets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/resources/secrets_store_secret#value SecretsStoreSecret#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// Freeform text describing the secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/resources/secrets_store_secret#comment SecretsStoreSecret#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
}

