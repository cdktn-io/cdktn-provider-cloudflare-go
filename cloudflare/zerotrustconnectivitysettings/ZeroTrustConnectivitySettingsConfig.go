// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustconnectivitysettings

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ZeroTrustConnectivitySettingsConfig struct {
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
	// Cloudflare account ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_connectivity_settings#account_id ZeroTrustConnectivitySettings#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// A flag to enable the ICMP proxy for the account network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_connectivity_settings#icmp_proxy_enabled ZeroTrustConnectivitySettings#icmp_proxy_enabled}
	IcmpProxyEnabled interface{} `field:"optional" json:"icmpProxyEnabled" yaml:"icmpProxyEnabled"`
	// A flag to enable WARP to WARP traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zero_trust_connectivity_settings#offramp_warp_enabled ZeroTrustConnectivitySettings#offramp_warp_enabled}
	OfframpWarpEnabled interface{} `field:"optional" json:"offrampWarpEnabled" yaml:"offrampWarpEnabled"`
}

