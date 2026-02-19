// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustdevicemanagednetworks

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ZeroTrustDeviceManagedNetworksConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_device_managed_networks#account_id ZeroTrustDeviceManagedNetworks#account_id}.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The configuration object containing information for the WARP client to detect the managed network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_device_managed_networks#config ZeroTrustDeviceManagedNetworks#config}
	Config *ZeroTrustDeviceManagedNetworksConfigA `field:"required" json:"config" yaml:"config"`
	// The name of the device managed network. This name must be unique.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_device_managed_networks#name ZeroTrustDeviceManagedNetworks#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of device managed network. Available values: "tls".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_device_managed_networks#type ZeroTrustDeviceManagedNetworks#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

