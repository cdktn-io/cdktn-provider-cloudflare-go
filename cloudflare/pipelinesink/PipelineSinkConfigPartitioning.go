// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipelinesink


type PipelineSinkConfigPartitioning struct {
	// The pattern of the date string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.0/docs/resources/pipeline_sink#time_pattern PipelineSink#time_pattern}
	TimePattern *string `field:"optional" json:"timePattern" yaml:"timePattern"`
}

