// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustdevicepostureintegration


type ZeroTrustDevicePostureIntegrationConfigA struct {
	// If present, this id will be passed in the `CF-Access-Client-ID` header when hitting the `api_url`.
	//
	// This id will be passed in the `CF-Access-Client-ID` header when hitting the `api_url`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/zero_trust_device_posture_integration#access_client_id ZeroTrustDevicePostureIntegration#access_client_id}
	AccessClientId *string `field:"optional" json:"accessClientId" yaml:"accessClientId"`
	// If present, this secret will be passed in the `CF-Access-Client-Secret` header when hitting the `api_url`.
	//
	// This secret will be passed in the `CF-Access-Client-Secret` header when hitting the `api_url`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/zero_trust_device_posture_integration#access_client_secret ZeroTrustDevicePostureIntegration#access_client_secret}
	AccessClientSecret *string `field:"optional" json:"accessClientSecret" yaml:"accessClientSecret"`
	// The Workspace One API URL provided in the Workspace One Admin Dashboard.
	//
	// The Crowdstrike API URL.
	//
	// The Uptycs API URL.
	//
	// The Tanium API URL.
	//
	// The SentinelOne S2S API URL.
	//
	// The Custom Device Posture Integration  API URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/zero_trust_device_posture_integration#api_url ZeroTrustDevicePostureIntegration#api_url}
	ApiUrl *string `field:"optional" json:"apiUrl" yaml:"apiUrl"`
	// The Workspace One Authorization URL depending on your region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/zero_trust_device_posture_integration#auth_url ZeroTrustDevicePostureIntegration#auth_url}
	AuthUrl *string `field:"optional" json:"authUrl" yaml:"authUrl"`
	// The Workspace One client ID provided in the Workspace One Admin Dashboard.
	//
	// The Crowdstrike client ID.
	//
	// The Intune client ID.
	//
	// The Kolide client ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/zero_trust_device_posture_integration#client_id ZeroTrustDevicePostureIntegration#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// The Uptycs client secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/zero_trust_device_posture_integration#client_key ZeroTrustDevicePostureIntegration#client_key}
	ClientKey *string `field:"optional" json:"clientKey" yaml:"clientKey"`
	// The Workspace One client secret provided in the Workspace One Admin Dashboard.
	//
	// The Crowdstrike client secret.
	//
	// The Uptycs client secret.
	//
	// The Intune client secret.
	//
	// The Kolide client secret.
	//
	// The Tanium client secret.
	//
	// The SentinelOne S2S client secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/zero_trust_device_posture_integration#client_secret ZeroTrustDevicePostureIntegration#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// The Crowdstrike customer ID.
	//
	// The Uptycs customer ID.
	//
	// The Intune customer ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/zero_trust_device_posture_integration#customer_id ZeroTrustDevicePostureIntegration#customer_id}
	CustomerId *string `field:"optional" json:"customerId" yaml:"customerId"`
}

