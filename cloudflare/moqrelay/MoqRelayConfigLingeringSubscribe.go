// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package moqrelay


type MoqRelayConfigLingeringSubscribe struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/resources/moq_relay#enabled MoqRelay#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Relay-level ceiling on lingering subscribe timeout (ms). Default 30000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/resources/moq_relay#max_timeout_ms MoqRelay#max_timeout_ms}
	MaxTimeoutMs *float64 `field:"optional" json:"maxTimeoutMs" yaml:"maxTimeoutMs"`
}

