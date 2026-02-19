// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package registrardomain

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RegistrarDomainConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/registrar_domain#account_id RegistrarDomain#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Domain name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/registrar_domain#domain_name RegistrarDomain#domain_name}
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Auto-renew controls whether subscription is automatically renewed upon domain expiration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/registrar_domain#auto_renew RegistrarDomain#auto_renew}
	AutoRenew interface{} `field:"optional" json:"autoRenew" yaml:"autoRenew"`
	// Shows whether a registrar lock is in place for a domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/registrar_domain#locked RegistrarDomain#locked}
	Locked interface{} `field:"optional" json:"locked" yaml:"locked"`
	// Privacy option controls redacting WHOIS information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/registrar_domain#privacy RegistrarDomain#privacy}
	Privacy interface{} `field:"optional" json:"privacy" yaml:"privacy"`
}

