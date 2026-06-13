// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workflow


type WorkflowSchedules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/resources/workflow#cron Workflow#cron}.
	Cron *string `field:"required" json:"cron" yaml:"cron"`
}

