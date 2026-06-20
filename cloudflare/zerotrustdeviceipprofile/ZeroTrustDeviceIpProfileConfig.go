// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustdeviceipprofile

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ZeroTrustDeviceIpProfileConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.0/docs/resources/zero_trust_device_ip_profile#account_id ZeroTrustDeviceIpProfile#account_id}.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The wirefilter expression to match registrations. Available values: "identity.name", "identity.email", "identity.groups.id", "identity.groups.name", "identity.groups.email", "identity.saml_attributes".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.0/docs/resources/zero_trust_device_ip_profile#match ZeroTrustDeviceIpProfile#match}
	Match *string `field:"required" json:"match" yaml:"match"`
	// A user-friendly name for the Device IP profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.0/docs/resources/zero_trust_device_ip_profile#name ZeroTrustDeviceIpProfile#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The precedence of the Device IP profile.
	//
	// Lower values indicate higher precedence. Device IP profile will be evaluated in ascending order of this field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.0/docs/resources/zero_trust_device_ip_profile#precedence ZeroTrustDeviceIpProfile#precedence}
	Precedence *float64 `field:"required" json:"precedence" yaml:"precedence"`
	// The ID of the Subnet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.0/docs/resources/zero_trust_device_ip_profile#subnet_id ZeroTrustDeviceIpProfile#subnet_id}
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// An optional description of the Device IP profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.0/docs/resources/zero_trust_device_ip_profile#description ZeroTrustDeviceIpProfile#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether the Device IP profile will be applied to matching devices.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.0/docs/resources/zero_trust_device_ip_profile#enabled ZeroTrustDeviceIpProfile#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

