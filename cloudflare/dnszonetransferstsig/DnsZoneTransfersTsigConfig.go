// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dnszonetransferstsig

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DnsZoneTransfersTsigConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/dns_zone_transfers_tsig#account_id DnsZoneTransfersTsig#account_id}.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// TSIG algorithm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/dns_zone_transfers_tsig#algo DnsZoneTransfersTsig#algo}
	Algo *string `field:"required" json:"algo" yaml:"algo"`
	// TSIG key name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/dns_zone_transfers_tsig#name DnsZoneTransfersTsig#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// TSIG secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/dns_zone_transfers_tsig#secret DnsZoneTransfersTsig#secret}
	Secret *string `field:"required" json:"secret" yaml:"secret"`
}

