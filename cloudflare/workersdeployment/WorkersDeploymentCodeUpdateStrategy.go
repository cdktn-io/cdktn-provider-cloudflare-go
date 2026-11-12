// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workersdeployment


type WorkersDeploymentCodeUpdateStrategy struct {
	// For `mode: immediate`: Apply Durable Object code updates immediately.
	//
	// For `mode: deferred`: Allow active Durable Objects to hibernate before updating their code.
	// Available values: "immediate", "deferred".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/workers_deployment#mode WorkersDeployment#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// For `mode: deferred`: Set the maximum number of seconds to wait before forcing the update.
	//
	// When omitted, the default maximum delay applies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/workers_deployment#max_delay WorkersDeployment#max_delay}
	MaxDelay *float64 `field:"optional" json:"maxDelay" yaml:"maxDelay"`
}

