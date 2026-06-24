// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package magictransitcf1site

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MagicTransitCf1SiteConfig struct {
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
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.1/docs/resources/magic_transit_cf1_site#account_id MagicTransitCf1Site#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.1/docs/resources/magic_transit_cf1_site#body MagicTransitCf1Site#body}.
	Body interface{} `field:"required" json:"body" yaml:"body"`
	// A human-provided description of the CF1 Site.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.1/docs/resources/magic_transit_cf1_site#description MagicTransitCf1Site#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.1/docs/resources/magic_transit_cf1_site#location MagicTransitCf1Site#location}.
	Location *MagicTransitCf1SiteLocation `field:"optional" json:"location" yaml:"location"`
	// A human-provided name describing the CF1 Site that should be unique within the account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.1/docs/resources/magic_transit_cf1_site#name MagicTransitCf1Site#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

