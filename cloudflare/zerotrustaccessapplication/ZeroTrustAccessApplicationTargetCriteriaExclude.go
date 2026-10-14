// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessapplication


type ZeroTrustAccessApplicationTargetCriteriaExclude struct {
	// Map of target tag keys to values. Values within a key are OR'd.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.28.0/docs/resources/zero_trust_access_application#tags ZeroTrustAccessApplication#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Hostname selector map. Only the hostname key is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.28.0/docs/resources/zero_trust_access_application#target_attributes ZeroTrustAccessApplication#target_attributes}
	TargetAttributes *ZeroTrustAccessApplicationTargetCriteriaExcludeTargetAttributes `field:"optional" json:"targetAttributes" yaml:"targetAttributes"`
}

