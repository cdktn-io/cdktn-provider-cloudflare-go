// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package contentscanningexpression


type ContentScanningExpressionBody struct {
	// Defines the custom content extraction expression used to reach content objects in the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/resources/content_scanning_expression#payload ContentScanningExpression#payload}
	Payload *string `field:"required" json:"payload" yaml:"payload"`
}

