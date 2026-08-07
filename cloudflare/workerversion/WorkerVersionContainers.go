// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workerversion


type WorkerVersionContainers struct {
	// Select which Durable Object class should get this container attached.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/resources/worker_version#class_name WorkerVersion#class_name}
	ClassName *string `field:"required" json:"className" yaml:"className"`
}

