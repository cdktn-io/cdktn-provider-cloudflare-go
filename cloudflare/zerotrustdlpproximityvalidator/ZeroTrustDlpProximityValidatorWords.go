// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustdlpproximityvalidator


type ZeroTrustDlpProximityValidatorWords struct {
	// How much a word contributes to a match's confidence when it is found nearby.
	//
	// Separate from a profile's high/medium/low confidence threshold: this describes how telling the
	// word is, not how confident the resulting match is.
	// Available values: "strong", "moderate", "weak".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/zero_trust_dlp_proximity_validator#weight ZeroTrustDlpProximityValidator#weight}
	Weight *string `field:"required" json:"weight" yaml:"weight"`
	// The word or phrase to look for. Must contain 1–128 characters and is matched case-insensitively.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/zero_trust_dlp_proximity_validator#word ZeroTrustDlpProximityValidator#word}
	Word *string `field:"required" json:"word" yaml:"word"`
}

