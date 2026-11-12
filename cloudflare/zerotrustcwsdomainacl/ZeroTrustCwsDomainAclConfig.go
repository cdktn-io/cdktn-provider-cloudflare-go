// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustcwsdomainacl

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ZeroTrustCwsDomainAclConfig struct {
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
	// Cloudflare account ID (32-character hexadecimal string).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/zero_trust_cws_domain_acl#account_id ZeroTrustCwsDomainAcl#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Default action when no entry matches Available values: "ALLOW", "BLOCK".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/zero_trust_cws_domain_acl#default_action ZeroTrustCwsDomainAcl#default_action}
	DefaultAction *string `field:"required" json:"defaultAction" yaml:"defaultAction"`
	// Domain entries for this ACL. Each domain must appear at most once.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/zero_trust_cws_domain_acl#entries ZeroTrustCwsDomainAcl#entries}
	Entries interface{} `field:"required" json:"entries" yaml:"entries"`
	// Name of the domain ACL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/zero_trust_cws_domain_acl#name ZeroTrustCwsDomainAcl#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

