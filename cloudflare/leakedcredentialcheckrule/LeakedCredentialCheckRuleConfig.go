// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package leakedcredentialcheckrule

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LeakedCredentialCheckRuleConfig struct {
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
	// Defines an identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/leaked_credential_check_rule#zone_id LeakedCredentialCheckRule#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// Defines ehe ruleset expression to use in matching the password in a request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/leaked_credential_check_rule#password LeakedCredentialCheckRule#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Defines the ruleset expression to use in matching the username in a request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/leaked_credential_check_rule#username LeakedCredentialCheckRule#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

