// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workerversion


type WorkerVersionExportsCache struct {
	// Whether caching is enabled for this entrypoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/resources/worker_version#enabled WorkerVersion#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

