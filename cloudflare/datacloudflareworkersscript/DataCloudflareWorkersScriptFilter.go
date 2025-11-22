// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflareworkersscript


type DataCloudflareWorkersScriptFilter struct {
	// Filter scripts by tags. Format: comma-separated list of tag:allowed pairs where allowed is 'yes' or 'no'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.13.0/docs/data-sources/workers_script#tags DataCloudflareWorkersScript#tags}
	Tags *string `field:"optional" json:"tags" yaml:"tags"`
}

