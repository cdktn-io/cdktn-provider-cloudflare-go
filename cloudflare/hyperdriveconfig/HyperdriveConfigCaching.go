// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package hyperdriveconfig


type HyperdriveConfigCaching struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/resources/hyperdrive_config#disabled HyperdriveConfig#disabled}.
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/resources/hyperdrive_config#max_age HyperdriveConfig#max_age}.
	MaxAge *float64 `field:"optional" json:"maxAge" yaml:"maxAge"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/resources/hyperdrive_config#stale_while_revalidate HyperdriveConfig#stale_while_revalidate}.
	StaleWhileRevalidate *float64 `field:"optional" json:"staleWhileRevalidate" yaml:"staleWhileRevalidate"`
}

