// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustdlpdataclass


type ZeroTrustDlpDataClassSensitivityLevels struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.1/docs/resources/zero_trust_dlp_data_class#group_id ZeroTrustDlpDataClass#group_id}.
	GroupId *string `field:"required" json:"groupId" yaml:"groupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.1/docs/resources/zero_trust_dlp_data_class#level_id ZeroTrustDlpDataClass#level_id}.
	LevelId *string `field:"required" json:"levelId" yaml:"levelId"`
}

