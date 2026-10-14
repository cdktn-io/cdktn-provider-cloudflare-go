// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workersscript


type WorkersScriptBindingsRetry struct {
	// Maximum number of retries after the initial request, not the total number of attempts.
	//
	// Defaults to 4. Zero disables retries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.28.0/docs/resources/workers_script#max_attempts WorkersScript#max_attempts}
	MaxAttempts *float64 `field:"optional" json:"maxAttempts" yaml:"maxAttempts"`
	// Retry timeout in milliseconds, measured from the start of the call.
	//
	// No retry starts after it expires, and a retry still running when it expires is cancelled. This is not a request timeout; it does not limit the initial request, which is still subject to any timeouts set by your Worker. Defaults to 10000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.28.0/docs/resources/workers_script#timeout_ms WorkersScript#timeout_ms}
	TimeoutMs *float64 `field:"optional" json:"timeoutMs" yaml:"timeoutMs"`
}

