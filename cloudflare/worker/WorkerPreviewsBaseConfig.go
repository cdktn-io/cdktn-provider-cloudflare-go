// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package worker


type WorkerPreviewsBaseConfig struct {
	// Cache options used when creating new Previews.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/resources/worker#cache_options Worker#cache_options}
	CacheOptions *WorkerPreviewsBaseConfigCacheOptions `field:"optional" json:"cacheOptions" yaml:"cacheOptions"`
	// Bindings used when creating new Previews, keyed by binding name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/resources/worker#env Worker#env}
	Env interface{} `field:"optional" json:"env" yaml:"env"`
	// Resource limits enforced at runtime for newly created Previews.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/resources/worker#limits Worker#limits}
	Limits *WorkerPreviewsBaseConfigLimits `field:"optional" json:"limits" yaml:"limits"`
	// Whether logpush is enabled when creating new Previews.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/resources/worker#logpush Worker#logpush}
	Logpush interface{} `field:"optional" json:"logpush" yaml:"logpush"`
	// Observability settings used when creating new Previews.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/resources/worker#observability Worker#observability}
	Observability *WorkerPreviewsBaseConfigObservability `field:"optional" json:"observability" yaml:"observability"`
	// Placement configuration used when creating new Previews.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/resources/worker#placement Worker#placement}
	Placement *WorkerPreviewsBaseConfigPlacement `field:"optional" json:"placement" yaml:"placement"`
	// Other Workers that should consume logs from newly created Previews.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/resources/worker#tail_consumers Worker#tail_consumers}
	TailConsumers interface{} `field:"optional" json:"tailConsumers" yaml:"tailConsumers"`
}

