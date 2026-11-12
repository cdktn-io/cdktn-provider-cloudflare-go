// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflarepipelinestream


type DataCloudflarePipelineStreamFilter struct {
	// Sort direction. Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/data-sources/pipeline_stream#direction DataCloudflarePipelineStream#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Filters streams by name (case-insensitive substring).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/data-sources/pipeline_stream#name DataCloudflarePipelineStream#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Sorts streams by this field. Defaults to `created_at`. Available values: "name", "created_at".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/data-sources/pipeline_stream#order_by DataCloudflarePipelineStream#order_by}
	OrderBy *string `field:"optional" json:"orderBy" yaml:"orderBy"`
	// Specifies the public ID of the pipeline.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/data-sources/pipeline_stream#pipeline_id DataCloudflarePipelineStream#pipeline_id}
	PipelineId *string `field:"optional" json:"pipelineId" yaml:"pipelineId"`
}

