// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tokenvalidationconfig


type TokenValidationConfigCredentials struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.14.0/docs/resources/token_validation_config#keys TokenValidationConfig#keys}.
	Keys interface{} `field:"required" json:"keys" yaml:"keys"`
}

