// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package moqrelay


type MoqRelayConfigUpstreamsUpstreams struct {
	// Upstream MOQT server publisher URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/resources/moq_relay#url MoqRelay#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
}

