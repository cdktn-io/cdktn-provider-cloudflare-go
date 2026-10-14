// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustresourcelibraryapplication

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ZeroTrustResourceLibraryApplicationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.28.0/docs/resources/zero_trust_resource_library_application#account_id ZeroTrustResourceLibraryApplication#account_id}.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Numeric identifier for an application category.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.28.0/docs/resources/zero_trust_resource_library_application#category_id ZeroTrustResourceLibraryApplication#category_id}
	CategoryId *float64 `field:"optional" json:"categoryId" yaml:"categoryId"`
	// Hostnames matched by the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.28.0/docs/resources/zero_trust_resource_library_application#hostnames ZeroTrustResourceLibraryApplication#hostnames}
	Hostnames *[]*string `field:"optional" json:"hostnames" yaml:"hostnames"`
	// Human-readable identifier for the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.28.0/docs/resources/zero_trust_resource_library_application#human_id ZeroTrustResourceLibraryApplication#human_id}
	HumanId *string `field:"optional" json:"humanId" yaml:"humanId"`
	// IP subnets for this application.
	//
	// Custom application create and update requests accept IPv4 prefix lengths /8 through /32 and IPv6 prefix lengths /32 through /128.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.28.0/docs/resources/zero_trust_resource_library_application#ip_subnets ZeroTrustResourceLibraryApplication#ip_subnets}
	IpSubnets *[]*string `field:"optional" json:"ipSubnets" yaml:"ipSubnets"`
	// Application name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.28.0/docs/resources/zero_trust_resource_library_application#name ZeroTrustResourceLibraryApplication#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Ports matched by the application, in `protocol/port` or inclusive `protocol/start-end` format (for example, `tcp/443` or `udp/10000-20000`).
	//
	// TCP and UDP ports must be between 0 and 65535.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.28.0/docs/resources/zero_trust_resource_library_application#port_protocols ZeroTrustResourceLibraryApplication#port_protocols}
	PortProtocols *[]*string `field:"optional" json:"portProtocols" yaml:"portProtocols"`
	// Support domains matched by the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.28.0/docs/resources/zero_trust_resource_library_application#support_domains ZeroTrustResourceLibraryApplication#support_domains}
	SupportDomains *[]*string `field:"optional" json:"supportDomains" yaml:"supportDomains"`
}

