// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package loadbalancer


type LoadBalancerPoolSets struct {
	// Disable this specific pool set. It will no longer be evaluated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/load_balancer#disabled LoadBalancer#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// A collection of fields used to directly respond to the client instead of routing to a pool.
	//
	// When supplied on a rule, that rule stops further rule evaluation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/load_balancer#fixed_response LoadBalancer#fixed_response}
	FixedResponse *LoadBalancerPoolSetsFixedResponse `field:"optional" json:"fixedResponse" yaml:"fixedResponse"`
	// Determines which requests a pool set applies to.
	//
	// Set `topology` to match by location or `default: true` to match all requests; the two are mutually exclusive. A pool set with no `match` matches all requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/load_balancer#match LoadBalancer#match}
	Match *LoadBalancerPoolSetsMatch `field:"optional" json:"match" yaml:"match"`
	// A human-readable name for this pool set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/load_balancer#name LoadBalancer#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The behavior a pool set applies when its `match` succeeds.
	//
	// A strict subset of a rule's `overrides`: a pool set replaces the topology wholesale with a flat pool list (`pools`), so only the declarative pool-routing fields plus `fallback_pool` and `steering_policy` are settable. All fields are optional.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/load_balancer#overrides LoadBalancer#overrides}
	Overrides *LoadBalancerPoolSetsOverrides `field:"optional" json:"overrides" yaml:"overrides"`
}

