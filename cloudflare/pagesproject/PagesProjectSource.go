// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pagesproject


type PagesProjectSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.12.0/docs/resources/pages_project#config PagesProject#config}.
	Config *PagesProjectSourceConfig `field:"optional" json:"config" yaml:"config"`
	// The source control management provider. Available values: "github", "gitlab".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.12.0/docs/resources/pages_project#type PagesProject#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

