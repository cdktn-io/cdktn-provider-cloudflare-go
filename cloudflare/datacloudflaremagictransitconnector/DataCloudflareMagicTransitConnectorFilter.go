// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflaremagictransitconnector


type DataCloudflareMagicTransitConnectorFilter struct {
	// Filter connectors by device type. Available values: "MANAGED", "LICENSED".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.1/docs/data-sources/magic_transit_connector#device_type DataCloudflareMagicTransitConnector#device_type}
	DeviceType *string `field:"optional" json:"deviceType" yaml:"deviceType"`
}

