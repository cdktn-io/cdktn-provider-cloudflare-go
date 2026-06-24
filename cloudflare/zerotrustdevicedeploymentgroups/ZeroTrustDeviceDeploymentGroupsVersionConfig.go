// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustdevicedeploymentgroups


type ZeroTrustDeviceDeploymentGroupsVersionConfig struct {
	// The target environment for the client version (e.g., windows, macos).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.1/docs/resources/zero_trust_device_deployment_groups#target_environment ZeroTrustDeviceDeploymentGroups#target_environment}
	TargetEnvironment *string `field:"required" json:"targetEnvironment" yaml:"targetEnvironment"`
	// The specific client version to deploy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.1/docs/resources/zero_trust_device_deployment_groups#version ZeroTrustDeviceDeploymentGroups#version}
	Version *string `field:"required" json:"version" yaml:"version"`
}

