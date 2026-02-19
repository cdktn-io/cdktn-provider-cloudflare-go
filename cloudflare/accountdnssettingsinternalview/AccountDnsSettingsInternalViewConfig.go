// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accountdnssettingsinternalview

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AccountDnsSettingsInternalViewConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/account_dns_settings_internal_view#account_id AccountDnsSettingsInternalView#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The name of the view.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/account_dns_settings_internal_view#name AccountDnsSettingsInternalView#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The list of zones linked to this view.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/account_dns_settings_internal_view#zones AccountDnsSettingsInternalView#zones}
	Zones *[]*string `field:"required" json:"zones" yaml:"zones"`
}

