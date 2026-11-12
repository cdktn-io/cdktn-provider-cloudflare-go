// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aisecuritycustomtopics


type AiSecurityCustomTopicsTopics struct {
	// Unique label identifier. Must contain only lowercase letters (a–z), digits (0–9), and hyphens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/ai_security_custom_topics#label AiSecurityCustomTopics#label}
	Label *string `field:"required" json:"label" yaml:"label"`
	// Description of the topic category.
	//
	// Must contain 2–50 UTF-8 bytes. Text in any language is supported, including spaces, numbers, punctuation, and symbols. Tabs, line breaks, invalid text, and invisible formatting characters are not allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/ai_security_custom_topics#topic AiSecurityCustomTopics#topic}
	Topic *string `field:"required" json:"topic" yaml:"topic"`
}

