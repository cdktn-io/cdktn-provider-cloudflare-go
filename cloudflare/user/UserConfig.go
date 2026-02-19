// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package user

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type UserConfig struct {
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
	// The country in which the user lives.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/user#country User#country}
	Country *string `field:"optional" json:"country" yaml:"country"`
	// User's first name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/user#first_name User#first_name}
	FirstName *string `field:"optional" json:"firstName" yaml:"firstName"`
	// User's last name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/user#last_name User#last_name}
	LastName *string `field:"optional" json:"lastName" yaml:"lastName"`
	// User's telephone number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/user#telephone User#telephone}
	Telephone *string `field:"optional" json:"telephone" yaml:"telephone"`
	// The zipcode or postal code where the user lives.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/user#zipcode User#zipcode}
	Zipcode *string `field:"optional" json:"zipcode" yaml:"zipcode"`
}

