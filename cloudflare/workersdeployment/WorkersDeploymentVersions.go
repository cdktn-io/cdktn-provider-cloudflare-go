// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workersdeployment


type WorkersDeploymentVersions struct {
	// Percentage of traffic served by this version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/resources/workers_deployment#percentage WorkersDeployment#percentage}
	Percentage *float64 `field:"required" json:"percentage" yaml:"percentage"`
	// Identifier of the Worker Version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/resources/workers_deployment#version_id WorkersDeployment#version_id}
	VersionId *string `field:"required" json:"versionId" yaml:"versionId"`
}

