// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package moqrelay


type MoqRelayConfigUpstreams struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/resources/moq_relay#enabled MoqRelay#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Ordered list of upstream MOQT server publishers.
	//
	// Each entry is an
	// object (not a bare string) so per-upstream configuration can be
	// added in the future without another breaking change.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/resources/moq_relay#upstreams MoqRelay#upstreams}
	Upstreams interface{} `field:"optional" json:"upstreams" yaml:"upstreams"`
}

