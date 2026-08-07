// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workflow


type WorkflowDefaultRetention struct {
	// Specifies the duration in milliseconds or as a string like '5 minutes'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/resources/workflow#error_retention Workflow#error_retention}
	ErrorRetention *map[string]interface{} `field:"optional" json:"errorRetention" yaml:"errorRetention"`
	// Specifies the duration in milliseconds or as a string like '5 minutes'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/resources/workflow#success_retention Workflow#success_retention}
	SuccessRetention *map[string]interface{} `field:"optional" json:"successRetention" yaml:"successRetention"`
}

