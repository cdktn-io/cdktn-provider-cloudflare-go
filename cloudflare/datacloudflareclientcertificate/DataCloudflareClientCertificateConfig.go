// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflareclientcertificate

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataCloudflareClientCertificateConfig struct {
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
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/data-sources/client_certificate#client_certificate_id DataCloudflareClientCertificate#client_certificate_id}
	ClientCertificateId *string `field:"optional" json:"clientCertificateId" yaml:"clientCertificateId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/data-sources/client_certificate#filter DataCloudflareClientCertificate#filter}.
	Filter *DataCloudflareClientCertificateFilter `field:"optional" json:"filter" yaml:"filter"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/data-sources/client_certificate#zone_id DataCloudflareClientCertificate#zone_id}
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}

