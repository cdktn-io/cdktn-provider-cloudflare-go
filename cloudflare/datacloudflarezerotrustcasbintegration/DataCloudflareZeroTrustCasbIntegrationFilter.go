// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezerotrustcasbintegration


type DataCloudflareZeroTrustCasbIntegrationFilter struct {
	// Filter by application/vendor (e.g., GOOGLE_WORKSPACE, MICROSOFT_INTERNAL).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/zero_trust_casb_integration#application DataCloudflareZeroTrustCasbIntegration#application}
	Application *string `field:"optional" json:"application" yaml:"application"`
	// Direction to order results. Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/zero_trust_casb_integration#direction DataCloudflareZeroTrustCasbIntegration#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Filter by DLP enabled status (true/false).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/zero_trust_casb_integration#dlp_enabled DataCloudflareZeroTrustCasbIntegration#dlp_enabled}
	DlpEnabled interface{} `field:"optional" json:"dlpEnabled" yaml:"dlpEnabled"`
	// Field to order results by. Available values: "application", "created", "name", "status".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/zero_trust_casb_integration#order DataCloudflareZeroTrustCasbIntegration#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
	// Page number within the paginated result set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/zero_trust_casb_integration#page DataCloudflareZeroTrustCasbIntegration#page}
	Page *float64 `field:"optional" json:"page" yaml:"page"`
	// Number of results per page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/zero_trust_casb_integration#page_size DataCloudflareZeroTrustCasbIntegration#page_size}
	PageSize *float64 `field:"optional" json:"pageSize" yaml:"pageSize"`
	// Search integrations by name or application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/zero_trust_casb_integration#search DataCloudflareZeroTrustCasbIntegration#search}
	Search *string `field:"optional" json:"search" yaml:"search"`
	// Filter by integration status. Available values: "Healthy", "Initializing", "Offline", "Unhealthy".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/zero_trust_casb_integration#status DataCloudflareZeroTrustCasbIntegration#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Filter by one enabled use case (for example, casb or ces).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/data-sources/zero_trust_casb_integration#use_cases DataCloudflareZeroTrustCasbIntegration#use_cases}
	UseCases *string `field:"optional" json:"useCases" yaml:"useCases"`
}

