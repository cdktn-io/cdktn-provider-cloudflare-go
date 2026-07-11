// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package moqrelay


type MoqRelayConfigA struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/resources/moq_relay#lingering_subscribe MoqRelay#lingering_subscribe}.
	LingeringSubscribe *MoqRelayConfigLingeringSubscribe `field:"optional" json:"lingeringSubscribe" yaml:"lingeringSubscribe"`
	// Upstreams are external MOQT server publishers that a relay falls back to when it has no local publisher for a requested namespace/track.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/resources/moq_relay#upstreams MoqRelay#upstreams}
	Upstreams *MoqRelayConfigUpstreams `field:"optional" json:"upstreams" yaml:"upstreams"`
}

