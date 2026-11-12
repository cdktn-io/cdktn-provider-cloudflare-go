// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustcwstenantstructure

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ZeroTrustCwsTenantStructureConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/zero_trust_cws_tenant_structure#account_id ZeroTrustCwsTenantStructure#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Platform this structure models. Immutable after creation. Available values: "aws", "cloudflare".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/zero_trust_cws_tenant_structure#application ZeroTrustCwsTenantStructure#application}
	Application *string `field:"required" json:"application" yaml:"application"`
	// Human-readable name; unique per account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/zero_trust_cws_tenant_structure#name ZeroTrustCwsTenantStructure#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Optional unique Rego mount path in the form tenant.<application>.<name>; application must match. Omit to leave unmounted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/zero_trust_cws_tenant_structure#import_path ZeroTrustCwsTenantStructure#import_path}
	ImportPath *string `field:"optional" json:"importPath" yaml:"importPath"`
}

