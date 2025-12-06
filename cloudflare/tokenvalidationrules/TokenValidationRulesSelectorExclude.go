// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tokenvalidationrules


type TokenValidationRulesSelectorExclude struct {
	// Excluded operation IDs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.14.0/docs/resources/token_validation_rules#operation_ids TokenValidationRules#operation_ids}
	OperationIds *[]*string `field:"optional" json:"operationIds" yaml:"operationIds"`
}

