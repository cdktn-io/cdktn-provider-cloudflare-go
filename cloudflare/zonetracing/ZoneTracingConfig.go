// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zonetracing

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ZoneTracingConfig struct {
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
	// Specify the zone ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zone_tracing#zone_id ZoneTracing#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// Up to 100 OpenTelemetry destination identifiers that receive traces.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zone_tracing#destinations ZoneTracing#destinations}
	Destinations *[]*string `field:"optional" json:"destinations" yaml:"destinations"`
	// Whether Cloudflare Traces is enabled for the zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zone_tracing#enabled ZoneTracing#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Whether trace context is sent externally or across a zone boundary.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zone_tracing#forward_context ZoneTracing#forward_context}
	ForwardContext interface{} `field:"optional" json:"forwardContext" yaml:"forwardContext"`
	// Whether traces are persisted in Cloudflare.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zone_tracing#persist ZoneTracing#persist}
	Persist interface{} `field:"optional" json:"persist" yaml:"persist"`
	// When inbound trace context may be continued. Authenticated propagation is not supported yet. Available values: "accept", "authenticated", "reject".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zone_tracing#propagation_policy ZoneTracing#propagation_policy}
	PropagationPolicy *string `field:"optional" json:"propagationPolicy" yaml:"propagationPolicy"`
	// The ratio of requests sampled for tracing, from 0 to 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/zone_tracing#sampling_ratio ZoneTracing#sampling_ratio}
	SamplingRatio *float64 `field:"optional" json:"samplingRatio" yaml:"samplingRatio"`
}

