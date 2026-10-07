// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fieldextractor


type FieldExtractorRulesFields struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/field_extractor#expression FieldExtractor#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/field_extractor#name FieldExtractor#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

