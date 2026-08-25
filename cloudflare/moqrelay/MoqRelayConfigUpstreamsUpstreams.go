// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package moqrelay


type MoqRelayConfigUpstreamsUpstreams struct {
	// Upstream MOQT server publisher URL.
	//
	// Must be an absolute URL with a
	// host and a scheme the relay can dial: moqt:// (raw QUIC) or https://
	// (WebTransport). Validated on update (PUT); rejected with 21013.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/moq_relay#url MoqRelay#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
}

