// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package tokenvalidationconfig

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type TokenValidationConfigConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/token_validation_config#credentials TokenValidationConfig#credentials}.
	Credentials *TokenValidationConfigCredentials `field:"required" json:"credentials" yaml:"credentials"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/token_validation_config#description TokenValidationConfig#description}.
	Description *string `field:"required" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/token_validation_config#title TokenValidationConfig#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/token_validation_config#token_sources TokenValidationConfig#token_sources}.
	TokenSources *[]*string `field:"required" json:"tokenSources" yaml:"tokenSources"`
	// Available values: "JWT".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/token_validation_config#token_type TokenValidationConfig#token_type}
	TokenType *string `field:"required" json:"tokenType" yaml:"tokenType"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/token_validation_config#zone_id TokenValidationConfig#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
}

