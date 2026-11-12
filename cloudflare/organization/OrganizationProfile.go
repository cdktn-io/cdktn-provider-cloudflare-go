// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package organization


type OrganizationProfile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/organization#business_address Organization#business_address}.
	BusinessAddress *string `field:"optional" json:"businessAddress" yaml:"businessAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/organization#business_email Organization#business_email}.
	BusinessEmail *string `field:"optional" json:"businessEmail" yaml:"businessEmail"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/organization#business_name Organization#business_name}.
	BusinessName *string `field:"optional" json:"businessName" yaml:"businessName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/organization#business_phone Organization#business_phone}.
	BusinessPhone *string `field:"optional" json:"businessPhone" yaml:"businessPhone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/organization#external_metadata Organization#external_metadata}.
	ExternalMetadata *string `field:"optional" json:"externalMetadata" yaml:"externalMetadata"`
}

