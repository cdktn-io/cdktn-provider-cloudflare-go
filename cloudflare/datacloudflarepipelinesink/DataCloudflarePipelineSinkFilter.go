// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflarepipelinesink


type DataCloudflarePipelineSinkFilter struct {
	// Sort direction. Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/data-sources/pipeline_sink#direction DataCloudflarePipelineSink#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Filters sinks by name (case-insensitive substring).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/data-sources/pipeline_sink#name DataCloudflarePipelineSink#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Sorts sinks by this field. Defaults to `created_at`. Available values: "name", "created_at".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/data-sources/pipeline_sink#order_by DataCloudflarePipelineSink#order_by}
	OrderBy *string `field:"optional" json:"orderBy" yaml:"orderBy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/data-sources/pipeline_sink#pipeline_id DataCloudflarePipelineSink#pipeline_id}.
	PipelineId *string `field:"optional" json:"pipelineId" yaml:"pipelineId"`
}

