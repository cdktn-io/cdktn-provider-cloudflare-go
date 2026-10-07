// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package worker

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WorkerConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/worker#account_id Worker#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Name of the Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/worker#name Worker#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// If true, delete the Worker even when other Workers still reference it.
	//
	// Service bindings in those Workers may be left broken. Durable Object namespaces implemented by the deleted Worker are deleted even if other Workers reference them.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/worker#force Worker#force}
	Force interface{} `field:"optional" json:"force" yaml:"force"`
	// Whether logpush is enabled for the Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/worker#logpush Worker#logpush}
	Logpush interface{} `field:"optional" json:"logpush" yaml:"logpush"`
	// Observability settings for the Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/worker#observability Worker#observability}
	Observability *WorkerObservability `field:"optional" json:"observability" yaml:"observability"`
	// Template configuration used when creating new Previews for this Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/worker#previews_base_config Worker#previews_base_config}
	PreviewsBaseConfig *WorkerPreviewsBaseConfig `field:"optional" json:"previewsBaseConfig" yaml:"previewsBaseConfig"`
	// Subdomain settings for the Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/worker#subdomain Worker#subdomain}
	Subdomain *WorkerSubdomain `field:"optional" json:"subdomain" yaml:"subdomain"`
	// Tags associated with the Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/worker#tags Worker#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// Other Workers that should consume logs from the Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/worker#tail_consumers Worker#tail_consumers}
	TailConsumers interface{} `field:"optional" json:"tailConsumers" yaml:"tailConsumers"`
}

