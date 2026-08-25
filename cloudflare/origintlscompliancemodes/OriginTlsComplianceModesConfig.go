// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package origintlscompliancemodes

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OriginTlsComplianceModesConfig struct {
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
	// List of TLS compliance modes that constrain the key-exchange algorithms Cloudflare may use when establishing the TLS connection to the zone's origin.
	//
	// Currently supported values are `fips` (FIPS-approved curves) and `pqh` (post-quantum hybrid). Future modes (e.g. `cnsa2`) may be added; clients should treat unknown values as opaque strings. Multiple modes are combined as the intersection of their permitted algorithm lists; selections whose intersection is empty are rejected. An empty list clears the constraint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/origin_tls_compliance_modes#value OriginTlsComplianceModes#value}
	Value *[]*string `field:"required" json:"value" yaml:"value"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/origin_tls_compliance_modes#zone_id OriginTlsComplianceModes#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
}

