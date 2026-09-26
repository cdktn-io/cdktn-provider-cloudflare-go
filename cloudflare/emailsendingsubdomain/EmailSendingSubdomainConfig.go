// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package emailsendingsubdomain

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EmailSendingSubdomainConfig struct {
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
	// The domain name within the zone.
	//
	// A wildcard is allowed only as the complete leftmost label (`*.example.com`) and requires the account wildcard Email Sending entitlement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/resources/email_sending_subdomain#name EmailSendingSubdomain#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/resources/email_sending_subdomain#zone_id EmailSendingSubdomain#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// Whether a send request that includes a recipient suppressed on this subdomain drops that recipient and still delivers to the rest, instead of failing the entire request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/resources/email_sending_subdomain#drop_suppressed_recipients EmailSendingSubdomain#drop_suppressed_recipients}
	DropSuppressedRecipients interface{} `field:"optional" json:"dropSuppressedRecipients" yaml:"dropSuppressedRecipients"`
	// Whether sent messages from this subdomain can be previewed in the activity log.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.26.0/docs/resources/email_sending_subdomain#preview_enabled EmailSendingSubdomain#preview_enabled}
	PreviewEnabled interface{} `field:"optional" json:"previewEnabled" yaml:"previewEnabled"`
}

