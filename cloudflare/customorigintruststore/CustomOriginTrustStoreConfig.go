// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package customorigintruststore

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CustomOriginTrustStoreConfig struct {
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
	// The root CA certificate in PEM format.
	//
	// Only root CA certificates are accepted; intermediate and leaf certificates are not supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.0/docs/resources/custom_origin_trust_store#certificate CustomOriginTrustStore#certificate}
	Certificate *string `field:"required" json:"certificate" yaml:"certificate"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.0/docs/resources/custom_origin_trust_store#zone_id CustomOriginTrustStore#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
}

