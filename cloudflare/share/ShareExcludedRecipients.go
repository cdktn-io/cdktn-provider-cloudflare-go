// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package share


type ShareExcludedRecipients struct {
	// The account to exclude from the organization-targeted share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/share#account_id Share#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
}

