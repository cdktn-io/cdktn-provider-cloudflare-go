// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package schemavalidationsettings

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SchemaValidationSettingsConfig struct {
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
	// The default mitigation action used Mitigation actions are as follows:.
	//
	// - `"log"` - log request when request does not conform to schema
	//   - `"block"` - deny access to the site when request does not conform to schema
	//   - `"none"` - skip running schema validation
	// Available values: "none", "log", "block".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/schema_validation_settings#validation_default_mitigation_action SchemaValidationSettings#validation_default_mitigation_action}
	ValidationDefaultMitigationAction *string `field:"required" json:"validationDefaultMitigationAction" yaml:"validationDefaultMitigationAction"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/schema_validation_settings#zone_id SchemaValidationSettings#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// When set, this overrides both zone level and operation level mitigation actions.
	//
	// - `"none"` - skip running schema validation entirely for the request
	//   - `null` - clears any existing override
	// Available values: "none".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/schema_validation_settings#validation_override_mitigation_action SchemaValidationSettings#validation_override_mitigation_action}
	ValidationOverrideMitigationAction *string `field:"optional" json:"validationOverrideMitigationAction" yaml:"validationOverrideMitigationAction"`
}

