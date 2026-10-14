// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aisearchinstance


type AiSearchInstanceMetadata struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.28.0/docs/resources/ai_search_instance#created_from_aisearch_wizard AiSearchInstance#created_from_aisearch_wizard}.
	CreatedFromAisearchWizard interface{} `field:"optional" json:"createdFromAisearchWizard" yaml:"createdFromAisearchWizard"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.28.0/docs/resources/ai_search_instance#created_from_emdash_plugin AiSearchInstance#created_from_emdash_plugin}.
	CreatedFromEmdashPlugin *AiSearchInstanceMetadataCreatedFromEmdashPlugin `field:"optional" json:"createdFromEmdashPlugin" yaml:"createdFromEmdashPlugin"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.28.0/docs/resources/ai_search_instance#worker_domain AiSearchInstance#worker_domain}.
	WorkerDomain *string `field:"optional" json:"workerDomain" yaml:"workerDomain"`
}

