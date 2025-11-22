// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tokenvalidationrules


type TokenValidationRulesSelectorInclude struct {
	// Included hostnames.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.13.0/docs/resources/token_validation_rules#host TokenValidationRules#host}
	Host *[]*string `field:"optional" json:"host" yaml:"host"`
}

