// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package tokenvalidationrules


type TokenValidationRulesSelectorInclude struct {
	// Included hostnames.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.1/docs/resources/token_validation_rules#host TokenValidationRules#host}
	Host *[]*string `field:"optional" json:"host" yaml:"host"`
}

