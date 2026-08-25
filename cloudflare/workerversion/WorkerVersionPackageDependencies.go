// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workerversion


type WorkerVersionPackageDependencies struct {
	// The exact version that was resolved and installed by the package manager.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/worker_version#installed_version WorkerVersion#installed_version}
	InstalledVersion *string `field:"required" json:"installedVersion" yaml:"installedVersion"`
	// The npm package name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/worker_version#name WorkerVersion#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The version constraint as written in package.json.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/worker_version#package_json_version WorkerVersion#package_json_version}
	PackageJsonVersion *string `field:"required" json:"packageJsonVersion" yaml:"packageJsonVersion"`
}

