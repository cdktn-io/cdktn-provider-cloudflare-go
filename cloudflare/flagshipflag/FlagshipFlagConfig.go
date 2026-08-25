// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package flagshipflag

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type FlagshipFlagConfig struct {
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
	// Cloudflare account ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/flagship_flag#account_id FlagshipFlag#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// App identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/flagship_flag#app_id FlagshipFlag#app_id}
	AppId *string `field:"required" json:"appId" yaml:"appId"`
	// Variation served when no rule matches or the flag is disabled. Must be a key in `variations`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/flagship_flag#default_variation FlagshipFlag#default_variation}
	DefaultVariation *string `field:"required" json:"defaultVariation" yaml:"defaultVariation"`
	// When false, the flag bypasses all rules and always serves `default_variation`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/flagship_flag#enabled FlagshipFlag#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Unique identifier for the flag within an app. Used in all evaluation and SDK calls.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/flagship_flag#key FlagshipFlag#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Targeting rules evaluated in ascending `priority`;
	//
	// the first matching rule wins. An empty array means the flag always serves `default_variation`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/flagship_flag#rules FlagshipFlag#rules}
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
	// Map of variation name to value.
	//
	// All values must be the same type (boolean, string, number, or JSON object/array). Each serialized value must be 10KB or smaller.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/flagship_flag#variations FlagshipFlag#variations}
	Variations *map[string]*string `field:"required" json:"variations" yaml:"variations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/flagship_flag#description FlagshipFlag#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Flag key (slug).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/flagship_flag#flag_key FlagshipFlag#flag_key}
	FlagKey *string `field:"optional" json:"flagKey" yaml:"flagKey"`
	// Value type of the flag's variations.
	//
	// Inferred from the variation values on write, so it may be omitted in requests.
	// Available values: "boolean", "string", "number", "json".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/flagship_flag#type FlagshipFlag#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

