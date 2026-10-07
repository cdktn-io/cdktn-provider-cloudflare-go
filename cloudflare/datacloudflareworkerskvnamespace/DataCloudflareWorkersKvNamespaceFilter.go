// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflareworkerskvnamespace


type DataCloudflareWorkersKvNamespaceFilter struct {
	// Sort namespaces in ascending (`asc`) or descending (`desc`) order. Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/workers_kv_namespace#direction DataCloudflareWorkersKvNamespace#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Namespace field to sort by (`id` or `title`). Available values: "id", "title".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/workers_kv_namespace#order DataCloudflareWorkersKvNamespace#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
}

