// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezerotrustdevicesubnet

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataCloudflareZeroTrustDeviceSubnetConfig struct {
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
	// The UUID of the subnet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/data-sources/zero_trust_device_subnet#subnet_id DataCloudflareZeroTrustDeviceSubnet#subnet_id}
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// Cloudflare account ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/data-sources/zero_trust_device_subnet#account_id DataCloudflareZeroTrustDeviceSubnet#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
}

