// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workerskv

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WorkersKvConfig struct {
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
	// ID of the Cloudflare account that owns the Workers KV namespaces.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/workers_kv#account_id WorkersKv#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// A key's name.
	//
	// The name may be at most 512 bytes. All printable, non-whitespace characters are valid. Use percent-encoding to define key names as part of a URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/workers_kv#key_name WorkersKv#key_name}
	KeyName *string `field:"required" json:"keyName" yaml:"keyName"`
	// ID of the Workers KV namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/workers_kv#namespace_id WorkersKv#namespace_id}
	NamespaceId *string `field:"required" json:"namespaceId" yaml:"namespaceId"`
	// A byte sequence to be stored, up to 25 MiB in length.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/workers_kv#value WorkersKv#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// Expires the key at a certain time, measured in number of seconds since the UNIX epoch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/workers_kv#expiration WorkersKv#expiration}
	Expiration *float64 `field:"optional" json:"expiration" yaml:"expiration"`
	// Expires the key after a number of seconds. Must be at least 60.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/workers_kv#expiration_ttl WorkersKv#expiration_ttl}
	ExpirationTtl *float64 `field:"optional" json:"expirationTtl" yaml:"expirationTtl"`
	// Associates arbitrary JSON data with a key/value pair.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/workers_kv#metadata WorkersKv#metadata}
	Metadata *string `field:"optional" json:"metadata" yaml:"metadata"`
}

