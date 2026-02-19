// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustgatewaysettings


type ZeroTrustGatewaySettingsSettingsBodyScanning struct {
	// Specify the inspection mode as either `deep` or `shallow`. Available values: "deep", "shallow".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_settings#inspection_mode ZeroTrustGatewaySettings#inspection_mode}
	InspectionMode *string `field:"optional" json:"inspectionMode" yaml:"inspectionMode"`
}

