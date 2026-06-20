// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package loadbalancermonitorgroup


type LoadBalancerMonitorGroupMembers struct {
	// Whether this monitor is enabled in the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.0/docs/resources/load_balancer_monitor_group#enabled LoadBalancerMonitorGroup#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The ID of the Monitor to use for checking the health of origins within this pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.0/docs/resources/load_balancer_monitor_group#monitor_id LoadBalancerMonitorGroup#monitor_id}
	MonitorId *string `field:"required" json:"monitorId" yaml:"monitorId"`
	// Whether this monitor is used for monitoring only (does not affect pool health).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.0/docs/resources/load_balancer_monitor_group#monitoring_only LoadBalancerMonitorGroup#monitoring_only}
	MonitoringOnly interface{} `field:"required" json:"monitoringOnly" yaml:"monitoringOnly"`
	// Whether this monitor must be healthy for the pool to be considered healthy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.0/docs/resources/load_balancer_monitor_group#must_be_healthy LoadBalancerMonitorGroup#must_be_healthy}
	MustBeHealthy interface{} `field:"required" json:"mustBeHealthy" yaml:"mustBeHealthy"`
}

