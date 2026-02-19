// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudforceonerequestpriority

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CloudforceOneRequestPriorityConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/cloudforce_one_request_priority#account_id CloudforceOneRequestPriority#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// List of labels.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/cloudforce_one_request_priority#labels CloudforceOneRequestPriority#labels}
	Labels *[]*string `field:"required" json:"labels" yaml:"labels"`
	// Priority.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/cloudforce_one_request_priority#priority CloudforceOneRequestPriority#priority}
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Requirement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/cloudforce_one_request_priority#requirement CloudforceOneRequestPriority#requirement}
	Requirement *string `field:"required" json:"requirement" yaml:"requirement"`
	// The CISA defined Traffic Light Protocol (TLP). Available values: "clear", "amber", "amber-strict", "green", "red".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/cloudforce_one_request_priority#tlp CloudforceOneRequestPriority#tlp}
	Tlp *string `field:"required" json:"tlp" yaml:"tlp"`
}

