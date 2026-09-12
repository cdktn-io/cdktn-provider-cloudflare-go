// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package nelsetting


type NelSettingValue struct {
	// Whether Network Error Logging is enabled for the zone. When enabled, browsers report network errors to Cloudflare's NEL endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.25.0/docs/resources/nel_setting#enabled NelSetting#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

