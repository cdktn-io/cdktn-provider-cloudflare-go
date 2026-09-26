// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workflow


type WorkflowConcurrency struct {
	// Maximum number of instances of this workflow that can run concurrently.
	//
	// Additional instances are queued and started as running instances complete. Must not exceed the account concurrency limit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/resources/workflow#limit Workflow#limit}
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
}

