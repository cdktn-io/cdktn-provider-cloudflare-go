// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package botmanagement


type BotManagementAiAgents struct {
	// Action to apply to one supported personal AI agent. Available values: "disabled", "block", "only_on_ad_pages".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/bot_management#anthropic_agent BotManagement#anthropic_agent}
	AnthropicAgent *string `field:"optional" json:"anthropicAgent" yaml:"anthropicAgent"`
	// Action to apply to one supported personal AI agent. Available values: "disabled", "block", "only_on_ad_pages".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/bot_management#google_agent BotManagement#google_agent}
	GoogleAgent *string `field:"optional" json:"googleAgent" yaml:"googleAgent"`
	// Action to apply to one supported personal AI agent. Available values: "disabled", "block", "only_on_ad_pages".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/bot_management#instinct BotManagement#instinct}
	Instinct *string `field:"optional" json:"instinct" yaml:"instinct"`
	// Action to apply to one supported personal AI agent. Available values: "disabled", "block", "only_on_ad_pages".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/bot_management#muse BotManagement#muse}
	Muse *string `field:"optional" json:"muse" yaml:"muse"`
	// Action to apply to one supported personal AI agent. Available values: "disabled", "block", "only_on_ad_pages".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/bot_management#openai_agent BotManagement#openai_agent}
	OpenaiAgent *string `field:"optional" json:"openaiAgent" yaml:"openaiAgent"`
}

