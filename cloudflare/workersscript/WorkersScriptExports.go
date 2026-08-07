// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workersscript


type WorkersScriptExports struct {
	// The kind of entrypoint. A `type: worker` entry overrides the top-level `cache_options` for this specific entrypoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/resources/workers_script#type WorkersScript#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Per-entrypoint cache override. When present, this overrides the top-level `cache_options` for this specific entrypoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/resources/workers_script#cache WorkersScript#cache}
	Cache *WorkersScriptExportsCache `field:"optional" json:"cache" yaml:"cache"`
}

