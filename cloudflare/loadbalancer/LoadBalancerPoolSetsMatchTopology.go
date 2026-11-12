// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package loadbalancer


type LoadBalancerPoolSetsMatchTopology struct {
	// A list of ISO 3166-1 alpha-2 country codes. Matches when the request's country is in this list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/load_balancer#countries LoadBalancer#countries}
	Countries *[]*string `field:"optional" json:"countries" yaml:"countries"`
	// A list of Cloudflare PoP codes. Matches when the request's PoP is in this list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/load_balancer#pops LoadBalancer#pops}
	Pops *[]*string `field:"optional" json:"pops" yaml:"pops"`
	// A list of Cloudflare region codes (e.g. `WNAM`, `ENAM`, `WEU`). Matches when the request's region is in this list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/load_balancer#regions LoadBalancer#regions}
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
}

