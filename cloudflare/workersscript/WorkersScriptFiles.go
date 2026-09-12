// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workersscript


type WorkersScriptFiles struct {
	// Content type of the file, such as `application/wasm`, `text/plain`, or `application/octet-stream`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.25.0/docs/resources/workers_script#content_type WorkersScript#content_type}
	ContentType *string `field:"required" json:"contentType" yaml:"contentType"`
	// Base64-encoded file content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.25.0/docs/resources/workers_script#content_base64 WorkersScript#content_base64}
	ContentBase64 *string `field:"optional" json:"contentBase64" yaml:"contentBase64"`
	// Path to the file content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.25.0/docs/resources/workers_script#content_file WorkersScript#content_file}
	ContentFile *string `field:"optional" json:"contentFile" yaml:"contentFile"`
}

