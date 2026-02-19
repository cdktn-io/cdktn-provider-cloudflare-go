// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudconnectorrules


type CloudConnectorRulesRulesParameters struct {
	// Host to perform Cloud Connection to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/cloud_connector_rules#host CloudConnectorRules#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
}

