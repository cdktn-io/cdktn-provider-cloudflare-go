// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package flagshipflag


type FlagshipFlagRulesRollout struct {
	// Percentage of matching traffic (0–100) served this variation.
	//
	// For multi-way splits, use cumulative upper bounds across rules (e.g. 30, 70, 100).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/resources/flagship_flag#percentage FlagshipFlag#percentage}
	Percentage *float64 `field:"required" json:"percentage" yaml:"percentage"`
	// Context attribute used for sticky bucketing. Defaults to `targetingKey`. If absent at evaluation time, bucketing is random per request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/resources/flagship_flag#attribute FlagshipFlag#attribute}
	Attribute *string `field:"optional" json:"attribute" yaml:"attribute"`
}

