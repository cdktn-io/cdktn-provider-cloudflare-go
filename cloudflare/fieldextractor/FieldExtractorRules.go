// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fieldextractor


type FieldExtractorRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/resources/field_extractor#fields FieldExtractor#fields}.
	Fields interface{} `field:"required" json:"fields" yaml:"fields"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/resources/field_extractor#ref FieldExtractor#ref}.
	Ref *string `field:"required" json:"ref" yaml:"ref"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/resources/field_extractor#description FieldExtractor#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

