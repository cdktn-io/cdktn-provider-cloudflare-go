// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package worker


type WorkerPreviewsBaseConfigObservabilityIssues struct {
	// Whether real-time Issues are enabled for the Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/worker#enabled Worker#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

