// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipelinestream


type PipelineStreamHttpCors struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/resources/pipeline_stream#origins PipelineStream#origins}.
	Origins *[]*string `field:"optional" json:"origins" yaml:"origins"`
}

