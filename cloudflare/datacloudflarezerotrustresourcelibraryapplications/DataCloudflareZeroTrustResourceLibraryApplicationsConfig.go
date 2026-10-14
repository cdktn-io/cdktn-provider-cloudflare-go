// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezerotrustresourcelibraryapplications

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataCloudflareZeroTrustResourceLibraryApplicationsConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.28.0/docs/data-sources/zero_trust_resource_library_applications#account_id DataCloudflareZeroTrustResourceLibraryApplications#account_id}.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Return only the listed properties on each application, as a comma-separated list.
	//
	// Use this to keep responses small when you only need part of each application, for
	// example populating a picker with `fields=id,name` instead of downloading every
	// hostname and IP subnet.
	//
	// Omit this parameter to receive the full application object, including
	// when lookup is used.
	//
	// `id` is always returned.
	//
	// Selectable properties: `id`, `name`, `human_id`, `version`, `hostnames`,
	// `support_domains`, `ip_subnets`, `port_protocols`, `supported`, `gen_ai_score`,
	// `application_confidence_score`, `created_at`, `updated_at`, `review_status`.
	//
	// Unknown or empty property names return `400`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.28.0/docs/data-sources/zero_trust_resource_library_applications#fields DataCloudflareZeroTrustResourceLibraryApplications#fields}
	Fields *string `field:"optional" json:"fields" yaml:"fields"`
	// Filter applications using `key:value` format.
	//
	// Supported filter keys:
	// - `name`: application name (for example, `name:HR`).
	// - `id`: application ID (for example, `id:498`).
	// - `human_id`: human-readable ID (for example, `human_id:HR`).
	// - `hostname`: hostname or support domain (for example, `hostname:portal.example.com`).
	// - `source`: application source name (for example, `source:cloudflare`).
	// - `ip_subnet`: CIDR containment; returns applications where a stored subnet contains the search value (for example, `ip_subnet:10.0.1.5/32` matches `10.0.0.0/16`).
	// - `category_id`: category ID (for example, `category_id:12`).
	// - `category_name`: category name (for example, `category_name:HR`).
	// - `supported`: supported Cloudflare product. Values: `GATEWAY`, `ACCESS`, `CASB`.
	// - `review_status`: account-specific Gateway review status. Values: `approved`, `unapproved`, `in_review`, `unreviewed`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.28.0/docs/data-sources/zero_trust_resource_library_applications#filter DataCloudflareZeroTrustResourceLibraryApplications#filter}
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
	// Number of results to return. Defaults to 25; values are clamped to 1–250.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.28.0/docs/data-sources/zero_trust_resource_library_applications#limit DataCloudflareZeroTrustResourceLibraryApplications#limit}
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// Find complete, case-insensitive application name mentions in a phrase or rule sentence.
	//
	// Misspelled names do not match. Returns matching applications
	// ranked by relevance. Must be between 2 and 1000 characters. Uses the same
	// filters, fields, and pagination as listing. `search` cannot be combined
	// with lookup. Results are ranked by relevance, so `order_by` is ignored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.28.0/docs/data-sources/zero_trust_resource_library_applications#lookup DataCloudflareZeroTrustResourceLibraryApplications#lookup}
	Lookup *string `field:"optional" json:"lookup" yaml:"lookup"`
	// Max items to fetch, default: 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.28.0/docs/data-sources/zero_trust_resource_library_applications#max_items DataCloudflareZeroTrustResourceLibraryApplications#max_items}
	MaxItems *float64 `field:"optional" json:"maxItems" yaml:"maxItems"`
	// Number of results to skip. Defaults to 0; negative values are clamped to 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.28.0/docs/data-sources/zero_trust_resource_library_applications#offset DataCloudflareZeroTrustResourceLibraryApplications#offset}
	Offset *float64 `field:"optional" json:"offset" yaml:"offset"`
	// Order results using field:direction format.
	//
	// Supported fields are name, id, human_id,
	// category_id, application_type, application_confidence_score, and gen_ai_score.
	// Supported directions are asc and desc. Ignored when search is provided; results are
	// ranked by relevance instead. Also ignored when lookup is provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.28.0/docs/data-sources/zero_trust_resource_library_applications#order_by DataCloudflareZeroTrustResourceLibraryApplications#order_by}
	OrderBy *string `field:"optional" json:"orderBy" yaml:"orderBy"`
	// Fuzzy search across application name and hostnames.
	//
	// Results are ranked by relevance. Must be between 2 and 200 characters. Can be combined with filter parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.28.0/docs/data-sources/zero_trust_resource_library_applications#search DataCloudflareZeroTrustResourceLibraryApplications#search}
	Search *string `field:"optional" json:"search" yaml:"search"`
}

