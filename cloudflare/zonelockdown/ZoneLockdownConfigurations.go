// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zonelockdown


type ZoneLockdownConfigurations struct {
	// For `target: ip`: The configuration target.
	//
	// You must set the target to `ip` when specifying an IP address in the Zone Lockdown rule.
	//
	// For `target: ip_range`: The configuration target. You must set the target to `ip_range` when specifying an IP address range in the Zone Lockdown rule.
	// Available values: "ip", "ip_range".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/zone_lockdown#target ZoneLockdown#target}
	Target *string `field:"optional" json:"target" yaml:"target"`
	// For `target: ip`: The IP address to match.
	//
	// This address will be compared to the IP address of incoming requests.
	//
	// For `target: ip_range`: The IP address range to match. You can only use prefix lengths `/16` and `/24`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/zone_lockdown#value ZoneLockdown#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

