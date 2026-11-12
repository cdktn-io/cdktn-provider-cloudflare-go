// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflareorganization


type DataCloudflareOrganizationFilterMember struct {
	// Filter organizations to those where the given user, by email, is an active member of that exact organization.
	//
	// This does not include hierarchy-derived access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/data-sources/organization#email DataCloudflareOrganization#email}
	Email *string `field:"optional" json:"email" yaml:"email"`
}

