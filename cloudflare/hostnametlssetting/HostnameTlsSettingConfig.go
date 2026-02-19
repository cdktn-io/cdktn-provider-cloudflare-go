// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package hostnametlssetting

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type HostnameTlsSettingConfig struct {
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
	// The hostname for which the tls settings are set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/hostname_tls_setting#hostname HostnameTlsSetting#hostname}
	Hostname *string `field:"required" json:"hostname" yaml:"hostname"`
	// The TLS Setting name. Available values: "ciphers", "min_tls_version", "http2".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/hostname_tls_setting#setting_id HostnameTlsSetting#setting_id}
	SettingId *string `field:"required" json:"settingId" yaml:"settingId"`
	// The tls setting value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/hostname_tls_setting#value HostnameTlsSetting#value}
	Value *map[string]interface{} `field:"required" json:"value" yaml:"value"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/hostname_tls_setting#zone_id HostnameTlsSetting#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
}

