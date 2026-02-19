// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package streamaudiotrack

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type StreamAudioTrackConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The account identifier tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/stream_audio_track#account_id StreamAudioTrack#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// A Cloudflare-generated unique identifier for a media item.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/stream_audio_track#identifier StreamAudioTrack#identifier}
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// The unique identifier for an additional audio track.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/stream_audio_track#audio_identifier StreamAudioTrack#audio_identifier}
	AudioIdentifier *string `field:"optional" json:"audioIdentifier" yaml:"audioIdentifier"`
	// Denotes whether the audio track will be played by default in a player.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/stream_audio_track#default StreamAudioTrack#default}
	Default interface{} `field:"optional" json:"default" yaml:"default"`
	// A string to uniquely identify the track amongst other audio track labels for the specified video.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/stream_audio_track#label StreamAudioTrack#label}
	Label *string `field:"optional" json:"label" yaml:"label"`
}

