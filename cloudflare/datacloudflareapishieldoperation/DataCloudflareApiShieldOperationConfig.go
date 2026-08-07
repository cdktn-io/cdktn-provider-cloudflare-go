// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflareapishieldoperation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataCloudflareApiShieldOperationConfig struct {
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
	// Add feature(s) to the results.
	//
	// The feature name that is given here corresponds to the resulting feature object. Have a look at the top-level object description for more details on the specific meaning.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/data-sources/api_shield_operation#feature DataCloudflareApiShieldOperation#feature}
	Feature *[]*string `field:"optional" json:"feature" yaml:"feature"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/data-sources/api_shield_operation#filter DataCloudflareApiShieldOperation#filter}.
	Filter *DataCloudflareApiShieldOperationFilter `field:"optional" json:"filter" yaml:"filter"`
	// UUID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/data-sources/api_shield_operation#operation_id DataCloudflareApiShieldOperation#operation_id}
	OperationId *string `field:"optional" json:"operationId" yaml:"operationId"`
	// When true, includes OpenAPI schemas (both uploaded and learned) for the operation in the response.
	//
	// Due to the conversion overhead, this parameter is only supported on single-operation retrieval.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/data-sources/api_shield_operation#with_schemas DataCloudflareApiShieldOperation#with_schemas}
	WithSchemas interface{} `field:"optional" json:"withSchemas" yaml:"withSchemas"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/data-sources/api_shield_operation#zone_id DataCloudflareApiShieldOperation#zone_id}
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}

