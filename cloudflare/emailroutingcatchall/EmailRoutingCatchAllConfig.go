// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package emailroutingcatchall

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EmailRoutingCatchAllConfig struct {
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
	// List actions for the catch-all routing rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/resources/email_routing_catch_all#actions EmailRoutingCatchAll#actions}
	Actions interface{} `field:"required" json:"actions" yaml:"actions"`
	// List of matchers for the catch-all routing rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/resources/email_routing_catch_all#matchers EmailRoutingCatchAll#matchers}
	Matchers interface{} `field:"required" json:"matchers" yaml:"matchers"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/resources/email_routing_catch_all#zone_id EmailRoutingCatchAll#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// Routing rule status.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/resources/email_routing_catch_all#enabled EmailRoutingCatchAll#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Routing rule name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/resources/email_routing_catch_all#name EmailRoutingCatchAll#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Public tag (script_tag) of the Worker that owns this rule. Required when `source` is `wrangler`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/resources/email_routing_catch_all#owner_worker_tag EmailRoutingCatchAll#owner_worker_tag}
	OwnerWorkerTag *string `field:"optional" json:"ownerWorkerTag" yaml:"ownerWorkerTag"`
	// Who manages the rule.
	//
	// `api` covers dashboard, generic API, and Terraform;
	// `wrangler` means the rule is managed by a Worker's wrangler.jsonc. Defaults
	// to `api` when omitted on write.
	// Available values: "api", "wrangler".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.22.0/docs/resources/email_routing_catch_all#source EmailRoutingCatchAll#source}
	Source *string `field:"optional" json:"source" yaml:"source"`
}

