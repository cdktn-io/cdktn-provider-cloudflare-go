// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workersscript


type WorkersScriptExportsCache struct {
	// Whether caching is enabled for this entrypoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/resources/workers_script#enabled WorkersScript#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

