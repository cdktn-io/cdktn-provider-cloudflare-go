// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package snippets

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SnippetsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/snippets#files Snippets#files}
	Files *[]*string `field:"required" json:"files" yaml:"files"`
	// Metadata about the snippet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/snippets#metadata Snippets#metadata}
	Metadata *SnippetsMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// The identifying name of the snippet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/snippets#snippet_name Snippets#snippet_name}
	SnippetName *string `field:"required" json:"snippetName" yaml:"snippetName"`
	// The unique ID of the zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/snippets#zone_id Snippets#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
}

