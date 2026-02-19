// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apishieldschema

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ApiShieldSchemaConfig struct {
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
	// Schema file bytes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/api_shield_schema#file ApiShieldSchema#file}
	File *string `field:"required" json:"file" yaml:"file"`
	// Kind of schema Available values: "openapi_v3".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/api_shield_schema#kind ApiShieldSchema#kind}
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/api_shield_schema#zone_id ApiShieldSchema#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// Name of the schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/api_shield_schema#name ApiShieldSchema#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/api_shield_schema#schema_id ApiShieldSchema#schema_id}.
	SchemaId *string `field:"optional" json:"schemaId" yaml:"schemaId"`
	// Flag whether schema is enabled for validation. Available values: "true", "false".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/api_shield_schema#validation_enabled ApiShieldSchema#validation_enabled}
	ValidationEnabled *string `field:"optional" json:"validationEnabled" yaml:"validationEnabled"`
}

