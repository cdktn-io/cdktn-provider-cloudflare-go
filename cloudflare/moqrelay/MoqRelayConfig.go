// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package moqrelay

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MoqRelayConfig struct {
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
	// Cloudflare account identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/resources/moq_relay#account_id MoqRelay#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Human-readable name for the relay.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/resources/moq_relay#name MoqRelay#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// upstreams and lingering_subscribe are mutually exclusive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/resources/moq_relay#config MoqRelay#config}
	Config *MoqRelayConfigA `field:"optional" json:"config" yaml:"config"`
}

