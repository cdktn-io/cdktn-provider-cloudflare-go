// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package account

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AccountConfig struct {
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
	// Account name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/account#name Account#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Parent container details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/account#managed_by Account#managed_by}
	ManagedBy *AccountManagedBy `field:"optional" json:"managedBy" yaml:"managedBy"`
	// Account settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/account#settings Account#settings}
	Settings *AccountSettings `field:"optional" json:"settings" yaml:"settings"`
	// Set to `true` and omit `unit` to create a standalone Free Account. If provided, this field must be `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/account#standalone Account#standalone}
	Standalone interface{} `field:"optional" json:"standalone" yaml:"standalone"`
	// Available values: "standard", "enterprise".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/account#type Account#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Information related to the tenant unit.
	//
	// Provide its ID and omit `standalone` to create the Account within an Organization. See https://developers.cloudflare.com/tenant/how-to/manage-accounts/.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/account#unit Account#unit}
	Unit *AccountUnit `field:"optional" json:"unit" yaml:"unit"`
}

