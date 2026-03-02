// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package clientcertificate

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ClientCertificateConfig struct {
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
	// The Certificate Signing Request (CSR). Must be newline-encoded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.18.0/docs/resources/client_certificate#csr ClientCertificate#csr}
	Csr *string `field:"required" json:"csr" yaml:"csr"`
	// The number of days the Client Certificate will be valid after the issued_on date.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.18.0/docs/resources/client_certificate#validity_days ClientCertificate#validity_days}
	ValidityDays *float64 `field:"required" json:"validityDays" yaml:"validityDays"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.18.0/docs/resources/client_certificate#zone_id ClientCertificate#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.18.0/docs/resources/client_certificate#reactivate ClientCertificate#reactivate}.
	Reactivate interface{} `field:"optional" json:"reactivate" yaml:"reactivate"`
}

