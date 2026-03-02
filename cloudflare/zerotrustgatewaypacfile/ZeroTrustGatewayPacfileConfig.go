// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustgatewaypacfile

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ZeroTrustGatewayPacfileConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.18.0/docs/resources/zero_trust_gateway_pacfile#account_id ZeroTrustGatewayPacfile#account_id}.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Actual contents of the PAC file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.18.0/docs/resources/zero_trust_gateway_pacfile#contents ZeroTrustGatewayPacfile#contents}
	Contents *string `field:"required" json:"contents" yaml:"contents"`
	// Name of the PAC file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.18.0/docs/resources/zero_trust_gateway_pacfile#name ZeroTrustGatewayPacfile#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Detailed description of the PAC file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.18.0/docs/resources/zero_trust_gateway_pacfile#description ZeroTrustGatewayPacfile#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// URL-friendly version of the PAC file name. If not provided, it will be auto-generated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.18.0/docs/resources/zero_trust_gateway_pacfile#slug ZeroTrustGatewayPacfile#slug}
	Slug *string `field:"optional" json:"slug" yaml:"slug"`
}

