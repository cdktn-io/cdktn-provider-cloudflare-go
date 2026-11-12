// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package loadbalancer


type LoadBalancerPoolSetsMatch struct {
	// When true, matches every request. Cannot be combined with `topology`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/load_balancer#default LoadBalancer#default}
	Default interface{} `field:"optional" json:"default" yaml:"default"`
	// Matches requests by location.
	//
	// Set any combination of `pops`, `countries`, and `regions` (at least one is required); a request matches when its value appears in any populated list (e.g. `regions: ["WNAM"]` with `countries: ["US"]` matches a request in either WNAM or the US).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/load_balancer#topology LoadBalancer#topology}
	Topology *LoadBalancerPoolSetsMatchTopology `field:"optional" json:"topology" yaml:"topology"`
}

