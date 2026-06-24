// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package share


type ShareRecipients struct {
	// Organization identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.1/docs/resources/share#organization_id Share#organization_id}
	OrganizationId *string `field:"optional" json:"organizationId" yaml:"organizationId"`
	// The account that will receive the share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.1/docs/resources/share#recipient_account_id Share#recipient_account_id}
	RecipientAccountId *string `field:"optional" json:"recipientAccountId" yaml:"recipientAccountId"`
}

