// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflarepipelinestream


type DataCloudflarePipelineStreamFilter struct {
	// Filters streams by name (case-insensitive substring).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.1/docs/data-sources/pipeline_stream#name DataCloudflarePipelineStream#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specifies the public ID of the pipeline.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.1/docs/data-sources/pipeline_stream#pipeline_id DataCloudflarePipelineStream#pipeline_id}
	PipelineId *string `field:"optional" json:"pipelineId" yaml:"pipelineId"`
}

