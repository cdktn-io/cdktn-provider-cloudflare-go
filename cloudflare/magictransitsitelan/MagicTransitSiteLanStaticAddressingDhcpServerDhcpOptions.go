// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package magictransitsitelan


type MagicTransitSiteLanStaticAddressingDhcpServerDhcpOptions struct {
	// DHCP option number (1-254).
	//
	// Options 0 and 255 are reserved by RFC 2132. Options 3, 6, and 51 are not allowed because they conflict with connector-managed configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/resources/magic_transit_site_lan#code MagicTransitSiteLan#code}
	Code *float64 `field:"required" json:"code" yaml:"code"`
	// The type of the option value.
	//
	// text: a string (max 255 bytes). hex: colon-separated hex bytes (e.g. "01:04:aa:bb:cc", max 255 bytes). ip: an IPv4 address (e.g. "10.20.30.40"). byte: an unsigned integer 0-255 (1 byte). short: an unsigned integer 0-65535 (2 bytes). integer: an unsigned integer 0-4294967295 (4 bytes).
	// Available values: "text", "hex", "ip", "byte", "short", "integer".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/resources/magic_transit_site_lan#type MagicTransitSiteLan#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The option value, interpreted according to the type field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/resources/magic_transit_site_lan#value MagicTransitSiteLan#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

