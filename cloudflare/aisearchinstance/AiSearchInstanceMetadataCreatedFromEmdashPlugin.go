// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aisearchinstance


type AiSearchInstanceMetadataCreatedFromEmdashPlugin struct {
	// Available values: "native", "rest".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.28.1/docs/resources/ai_search_instance#type AiSearchInstance#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.28.1/docs/resources/ai_search_instance#version AiSearchInstance#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
}

