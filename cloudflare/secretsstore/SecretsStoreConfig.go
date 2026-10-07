// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package secretsstore

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SecretsStoreConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/secrets_store#account_id SecretsStore#account_id}.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The name of the store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/secrets_store#name SecretsStore#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// When true, cascade-deletes all secrets in the store before deleting the store itself.
	//
	// Required when deleting a non-empty store. Without this parameter, attempting to
	// delete a non-empty store returns 409.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/secrets_store#force SecretsStore#force}
	Force interface{} `field:"optional" json:"force" yaml:"force"`
}

