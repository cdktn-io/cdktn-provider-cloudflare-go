// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zonednssettings


type ZoneDnsSettingsNameservers struct {
	// Nameserver type. Available values: "cloudflare.standard", "cloudflare.advanced", "custom.account", "custom.tenant", "custom.zone", "custom".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zone_dns_settings#type ZoneDnsSettings#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Identifier of the account-owned Custom Nameserver Set to use for this zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zone_dns_settings#nameserver_set_id ZoneDnsSettings#nameserver_set_id}
	NameserverSetId *string `field:"optional" json:"nameserverSetId" yaml:"nameserverSetId"`
	// Configured nameserver set number to use for this zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zone_dns_settings#ns_set ZoneDnsSettings#ns_set}
	NsSet *float64 `field:"optional" json:"nsSet" yaml:"nsSet"`
}

