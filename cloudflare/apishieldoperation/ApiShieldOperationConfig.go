// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apishieldoperation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ApiShieldOperationConfig struct {
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
	// The endpoint which can contain path parameter templates in curly braces, each will be replaced from left to right with {varN}, starting with {var1}, during insertion.
	//
	// This will further be Cloudflare-normalized upon insertion. See: https://developers.cloudflare.com/rules/normalization/how-it-works/.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/resources/api_shield_operation#endpoint ApiShieldOperation#endpoint}
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// RFC3986-compliant host.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/resources/api_shield_operation#host ApiShieldOperation#host}
	Host *string `field:"required" json:"host" yaml:"host"`
	// The HTTP method used to access the endpoint. Available values: "GET", "POST", "HEAD", "OPTIONS", "PUT", "DELETE", "CONNECT", "PATCH", "TRACE".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/resources/api_shield_operation#method ApiShieldOperation#method}
	Method *string `field:"required" json:"method" yaml:"method"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/resources/api_shield_operation#zone_id ApiShieldOperation#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// Add feature(s) to the results.
	//
	// The feature name that is given here corresponds to the resulting feature object. Have a look at the top-level object description for more details on the specific meaning.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/resources/api_shield_operation#feature ApiShieldOperation#feature}
	Feature *[]*string `field:"optional" json:"feature" yaml:"feature"`
	// When true, includes OpenAPI schemas (both uploaded and learned) for the operation in the response.
	//
	// Due to the conversion overhead, this parameter is only supported on single-operation retrieval.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/resources/api_shield_operation#with_schemas ApiShieldOperation#with_schemas}
	WithSchemas interface{} `field:"optional" json:"withSchemas" yaml:"withSchemas"`
}

