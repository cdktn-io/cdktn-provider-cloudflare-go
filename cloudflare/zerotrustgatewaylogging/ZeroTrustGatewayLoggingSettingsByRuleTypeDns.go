// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustgatewaylogging


type ZeroTrustGatewayLoggingSettingsByRuleTypeDns struct {
	// Specify whether to log all requests to this service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.13.0/docs/resources/zero_trust_gateway_logging#log_all ZeroTrustGatewayLogging#log_all}
	LogAll interface{} `field:"optional" json:"logAll" yaml:"logAll"`
	// Specify whether to log only blocking requests to this service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.13.0/docs/resources/zero_trust_gateway_logging#log_blocks ZeroTrustGatewayLogging#log_blocks}
	LogBlocks interface{} `field:"optional" json:"logBlocks" yaml:"logBlocks"`
}

