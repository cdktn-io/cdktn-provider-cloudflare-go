// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package contentscanningexpression

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ContentScanningExpressionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/content_scanning_expression#body ContentScanningExpression#body}.
	Body interface{} `field:"required" json:"body" yaml:"body"`
	// Defines an identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/content_scanning_expression#zone_id ContentScanningExpression#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
}

