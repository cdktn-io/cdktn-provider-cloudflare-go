// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workersscript


type WorkersScriptCodeUpdateStrategy struct {
	// Apply Durable Object code updates immediately. Available values: "immediate", "deferred".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/workers_script#mode WorkersScript#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// Set the maximum number of seconds to wait before forcing the update. When omitted, the default maximum delay applies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/workers_script#max_delay WorkersScript#max_delay}
	MaxDelay *float64 `field:"optional" json:"maxDelay" yaml:"maxDelay"`
}

