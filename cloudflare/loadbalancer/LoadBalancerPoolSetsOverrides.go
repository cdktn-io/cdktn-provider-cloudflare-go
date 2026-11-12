// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package loadbalancer


type LoadBalancerPoolSetsOverrides struct {
	// The pool ID to use when all other pools are detected as unhealthy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/load_balancer#fallback_pool LoadBalancer#fallback_pool}
	FallbackPool *string `field:"optional" json:"fallbackPool" yaml:"fallbackPool"`
	// The default weight for pools not listed in `pool_weights`. The declarative alternative to `random_steering.default_weight`; mutually exclusive with `random_steering`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/load_balancer#pool_default_weight LoadBalancer#pool_default_weight}
	PoolDefaultWeight *float64 `field:"optional" json:"poolDefaultWeight" yaml:"poolDefaultWeight"`
	// A flat, ordered list of pool IDs to route the matched audience to.
	//
	// Replaces the resolved topology with exactly these pools. Mutually exclusive with `fixed_response`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/load_balancer#pools LoadBalancer#pools}
	Pools *[]*string `field:"optional" json:"pools" yaml:"pools"`
	// A mapping of pool IDs to custom weights, relative to the other pools.
	//
	// The declarative alternative to `random_steering.pool_weights`; mutually exclusive with `random_steering`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/load_balancer#pool_weights LoadBalancer#pool_weights}
	PoolWeights *map[string]*float64 `field:"optional" json:"poolWeights" yaml:"poolWeights"`
	// Steering Policy for this load balancer.
	//
	// - `"off"`: Use `default_pools`.
	// - `"geo"`: Use `region_pools`/`country_pools`/`pop_pools`. For non-proxied requests, the country for `country_pools` is determined by `location_strategy`.
	// - `"random"`: Select a pool randomly.
	// - `"dynamic_latency"`: Use round trip time to select the closest pool in default_pools (requires pool health checks).
	// - `"proximity"`: Use the pools' latitude and longitude to select the closest pool using the Cloudflare PoP location for proxied requests or the location determined by `location_strategy` for non-proxied requests.
	// - `"least_outstanding_requests"`: Select a pool by taking into consideration `random_steering` weights, as well as each pool's number of outstanding requests. Pools with more pending requests are weighted proportionately less relative to others.
	// - `"least_connections"`: Select a pool by taking into consideration `random_steering` weights, as well as each pool's number of open connections. Pools with more open connections are weighted proportionately less relative to others. Supported for HTTP/1 and HTTP/2 connections.
	// - `""`: Will map to `"geo"` if you use `region_pools`/`country_pools`/`pop_pools` otherwise `"off"`.
	// Available values: "off", "geo", "random", "dynamic_latency", "proximity", "least_outstanding_requests", "least_connections", "".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/load_balancer#steering_policy LoadBalancer#steering_policy}
	SteeringPolicy *string `field:"optional" json:"steeringPolicy" yaml:"steeringPolicy"`
}

