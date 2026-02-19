// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflarestreamcaptionlanguage

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataCloudflareStreamCaptionLanguageConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/data-sources/stream_caption_language#account_id DataCloudflareStreamCaptionLanguage#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// A Cloudflare-generated unique identifier for a media item.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/data-sources/stream_caption_language#identifier DataCloudflareStreamCaptionLanguage#identifier}
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// The language tag in BCP 47 format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/data-sources/stream_caption_language#language DataCloudflareStreamCaptionLanguage#language}
	Language *string `field:"required" json:"language" yaml:"language"`
}

