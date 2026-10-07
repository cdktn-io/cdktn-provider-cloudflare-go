// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package worker


type WorkerPreviewsBaseConfigPlacement struct {
	// TCP host and port for targeted placement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/worker#host Worker#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// HTTP hostname for targeted placement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/worker#hostname Worker#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// Enables [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement). Available values: "smart", "targeted".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/worker#mode Worker#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Cloud region for targeted placement in format 'provider:region'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/worker#region Worker#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Array of placement targets (currently limited to single target).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/worker#target Worker#target}
	Target interface{} `field:"optional" json:"target" yaml:"target"`
}

