// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflarepipelinesink


type DataCloudflarePipelineSinkFilter struct {
	// Filters sinks by name (case-insensitive substring).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/data-sources/pipeline_sink#name DataCloudflarePipelineSink#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/data-sources/pipeline_sink#pipeline_id DataCloudflarePipelineSink#pipeline_id}.
	PipelineId *string `field:"optional" json:"pipelineId" yaml:"pipelineId"`
}

