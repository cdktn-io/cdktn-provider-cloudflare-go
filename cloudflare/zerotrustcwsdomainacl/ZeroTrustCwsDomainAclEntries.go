// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustcwsdomainacl


type ZeroTrustCwsDomainAclEntries struct {
	// Per-domain decision overriding the parent ACL's default_action Available values: "ALLOW", "BLOCK".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/zero_trust_cws_domain_acl#action ZeroTrustCwsDomainAcl#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// Hostname or domain pattern, capped at 255 UTF-8 bytes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/zero_trust_cws_domain_acl#domain ZeroTrustCwsDomainAcl#domain}
	Domain *string `field:"required" json:"domain" yaml:"domain"`
}

