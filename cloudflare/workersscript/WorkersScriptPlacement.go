// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workersscript


type WorkersScriptPlacement struct {
	// Enables [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	//
	// Targeted placement mode.
	// Available values: "smart", "targeted".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/workers_script#mode WorkersScript#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

