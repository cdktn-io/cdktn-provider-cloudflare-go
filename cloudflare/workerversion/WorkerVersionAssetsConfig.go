// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workerversion


type WorkerVersionAssetsConfig struct {
	// The public URL path prefix under which assets are served.
	//
	// A null request value resets it to `/`; responses represent the root as `/`. All versions in a gradual deployment must use the same canonical value. To change it, first deploy the version containing the change at 100%.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/worker_version#base_path WorkerVersion#base_path}
	BasePath *string `field:"optional" json:"basePath" yaml:"basePath"`
	// Determines the redirects and rewrites of requests for HTML content. Available values: "auto-trailing-slash", "force-trailing-slash", "drop-trailing-slash", "none".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/worker_version#html_handling WorkerVersion#html_handling}
	HtmlHandling *string `field:"optional" json:"htmlHandling" yaml:"htmlHandling"`
	// Determines the response when a request does not match a static asset, and there is no Worker script.
	//
	// Available values: "none", "404-page", "single-page-application".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/worker_version#not_found_handling WorkerVersion#not_found_handling}
	NotFoundHandling *string `field:"optional" json:"notFoundHandling" yaml:"notFoundHandling"`
	// When a boolean true, requests will always invoke the Worker script.
	//
	// Otherwise, attempt to serve an asset matching the request, falling back to the Worker script. When a list of strings, contains path rules to control routing to either the Worker or assets. Glob (*) and negative (!) rules are supported. Rules must start with either '/' or '!/'. At least one non-negative rule must be provided, and negative rules have higher precedence than non-negative rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.27.0/docs/resources/worker_version#run_worker_first WorkerVersion#run_worker_first}
	RunWorkerFirst *map[string]interface{} `field:"optional" json:"runWorkerFirst" yaml:"runWorkerFirst"`
}

