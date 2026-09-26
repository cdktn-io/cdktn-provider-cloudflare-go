// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workersscript


type WorkersScriptObservabilityIssues struct {
	// Whether real-time Issues are enabled for the Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/resources/workers_script#enabled WorkersScript#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

