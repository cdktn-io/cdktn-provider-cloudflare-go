// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package flagshipflag


type FlagshipFlagRulesConditions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/resources/flagship_flag#attribute FlagshipFlag#attribute}.
	Attribute *string `field:"optional" json:"attribute" yaml:"attribute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/resources/flagship_flag#clauses FlagshipFlag#clauses}.
	Clauses interface{} `field:"optional" json:"clauses" yaml:"clauses"`
	// Available values: "AND", "OR".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/resources/flagship_flag#logical_operator FlagshipFlag#logical_operator}
	LogicalOperator *string `field:"optional" json:"logicalOperator" yaml:"logicalOperator"`
	// Available values: "equals", "not_equals", "greater_than", "less_than", "greater_than_or_equals", "less_than_or_equals", "contains", "starts_with", "ends_with", "in", "not_in".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/resources/flagship_flag#operator FlagshipFlag#operator}
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// Value to compare against the context attribute.
	//
	// Must be an array for `in` and `not_in`; numeric and ISO-8601 datetime strings are accepted by the ordering operators.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/resources/flagship_flag#value FlagshipFlag#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

