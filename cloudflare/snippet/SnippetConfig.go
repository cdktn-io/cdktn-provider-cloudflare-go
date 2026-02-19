// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package snippet

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SnippetConfig struct {
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
	// The list of files belonging to the snippet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/snippet#files Snippet#files}
	Files interface{} `field:"required" json:"files" yaml:"files"`
	// Metadata about the snippet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/snippet#metadata Snippet#metadata}
	Metadata *SnippetMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// The identifying name of the snippet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/snippet#snippet_name Snippet#snippet_name}
	SnippetName *string `field:"required" json:"snippetName" yaml:"snippetName"`
	// The unique ID of the zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/snippet#zone_id Snippet#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
}

