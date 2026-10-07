// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package worker


type WorkerPreviewsBaseConfigPlacementTarget struct {
	// TCP host:port for targeted placement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/worker#host Worker#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// HTTP hostname for targeted placement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/worker#hostname Worker#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// Cloud region in format 'provider:region'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/worker#region Worker#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

