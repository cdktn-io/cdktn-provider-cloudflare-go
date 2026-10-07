// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package worker


type WorkerPreviewsBaseConfigEnv struct {
	// The kind of resource that the binding provides.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/worker#type Worker#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

