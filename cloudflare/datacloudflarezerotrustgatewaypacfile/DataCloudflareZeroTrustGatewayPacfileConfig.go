// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezerotrustgatewaypacfile

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataCloudflareZeroTrustGatewayPacfileConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/data-sources/zero_trust_gateway_pacfile#pacfile_id DataCloudflareZeroTrustGatewayPacfile#pacfile_id}.
	PacfileId *string `field:"required" json:"pacfileId" yaml:"pacfileId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/data-sources/zero_trust_gateway_pacfile#account_id DataCloudflareZeroTrustGatewayPacfile#account_id}.
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
}

