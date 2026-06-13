// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package customcsr

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CustomCsrConfig struct {
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
	// The common name (domain) for the CSR. Must be at most 64 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/resources/custom_csr#common_name CustomCsr#common_name}
	CommonName *string `field:"required" json:"commonName" yaml:"commonName"`
	// Two-letter ISO 3166-1 alpha-2 country code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/resources/custom_csr#country CustomCsr#country}
	Country *string `field:"required" json:"country" yaml:"country"`
	// City or locality name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/resources/custom_csr#locality CustomCsr#locality}
	Locality *string `field:"required" json:"locality" yaml:"locality"`
	// Organization name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/resources/custom_csr#organization CustomCsr#organization}
	Organization *string `field:"required" json:"organization" yaml:"organization"`
	// Subject Alternative Names for the CSR. At least one SAN is required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/resources/custom_csr#sans CustomCsr#sans}
	Sans *[]*string `field:"required" json:"sans" yaml:"sans"`
	// State or province name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/resources/custom_csr#state CustomCsr#state}
	State *string `field:"required" json:"state" yaml:"state"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/resources/custom_csr#account_id CustomCsr#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Optional description for the CSR.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/resources/custom_csr#description CustomCsr#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Key algorithm to use for the CSR. Defaults to rsa2048 if not specified. Available values: "rsa2048", "p256v1".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/resources/custom_csr#key_type CustomCsr#key_type}
	KeyType *string `field:"optional" json:"keyType" yaml:"keyType"`
	// Human-readable name for the CSR.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/resources/custom_csr#name CustomCsr#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Organizational unit name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/resources/custom_csr#organizational_unit CustomCsr#organizational_unit}
	OrganizationalUnit *string `field:"optional" json:"organizationalUnit" yaml:"organizationalUnit"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/resources/custom_csr#zone_id CustomCsr#zone_id}
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}

