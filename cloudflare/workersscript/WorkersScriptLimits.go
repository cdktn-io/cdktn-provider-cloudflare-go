// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workersscript


type WorkersScriptLimits struct {
	// The amount of CPU time this Worker can use in milliseconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/resources/workers_script#cpu_ms WorkersScript#cpu_ms}
	CpuMs *float64 `field:"optional" json:"cpuMs" yaml:"cpuMs"`
	// The number of subrequests this Worker can make per request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/resources/workers_script#subrequests WorkersScript#subrequests}
	Subrequests *float64 `field:"optional" json:"subrequests" yaml:"subrequests"`
}

