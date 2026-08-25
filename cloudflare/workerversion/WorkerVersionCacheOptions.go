// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workerversion


type WorkerVersionCacheOptions struct {
	// Whether cached responses are shared across Worker version uploads.
	//
	// This is independent of `enabled`. It can stay true
	// while caching is off, so the preference survives turning
	// caching off and back on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/worker_version#cross_version_cache WorkerVersion#cross_version_cache}
	CrossVersionCache interface{} `field:"optional" json:"crossVersionCache" yaml:"crossVersionCache"`
	// Whether caching is enabled for this Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/worker_version#enabled WorkerVersion#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

