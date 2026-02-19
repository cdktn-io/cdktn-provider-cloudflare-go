// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zone

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ZoneConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zone#account Zone#account}.
	Account *ZoneAccount `field:"required" json:"account" yaml:"account"`
	// The domain name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zone#name Zone#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Indicates whether the zone is only using Cloudflare DNS services.
	//
	// A
	// true value means the zone will not receive security or performance
	// benefits.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zone#paused Zone#paused}
	Paused interface{} `field:"optional" json:"paused" yaml:"paused"`
	// A full zone implies that DNS is hosted with Cloudflare.
	//
	// A partial zone is
	// typically a partner-hosted zone or a CNAME setup.
	// Available values: "full", "partial", "secondary", "internal".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zone#type Zone#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// An array of domains used for custom name servers. This is only available for Business and Enterprise plans.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zone#vanity_name_servers Zone#vanity_name_servers}
	VanityNameServers *[]*string `field:"optional" json:"vanityNameServers" yaml:"vanityNameServers"`
}

