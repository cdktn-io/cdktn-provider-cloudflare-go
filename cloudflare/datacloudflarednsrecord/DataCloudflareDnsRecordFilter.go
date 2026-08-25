// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflarednsrecord


type DataCloudflareDnsRecordFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/data-sources/dns_record#comment DataCloudflareDnsRecord#comment}.
	Comment *DataCloudflareDnsRecordFilterComment `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/data-sources/dns_record#content DataCloudflareDnsRecord#content}.
	Content *DataCloudflareDnsRecordFilterContent `field:"optional" json:"content" yaml:"content"`
	// Direction to order DNS records in. Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/data-sources/dns_record#direction DataCloudflareDnsRecord#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Whether to match all search requirements or at least one (any).
	//
	// If set to `all`, acts like a logical AND between filters. If set to `any`, acts like a logical OR instead. Note that the interaction between tag filters is controlled by the `tag-match` parameter instead.
	// Available values: "any", "all".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/data-sources/dns_record#match DataCloudflareDnsRecord#match}
	Match *string `field:"optional" json:"match" yaml:"match"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/data-sources/dns_record#name DataCloudflareDnsRecord#name}.
	Name *DataCloudflareDnsRecordFilterName `field:"optional" json:"name" yaml:"name"`
	// Field to order DNS records by. Available values: "type", "name", "content", "ttl", "proxied".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/data-sources/dns_record#order DataCloudflareDnsRecord#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
	// Whether the record is receiving the performance and security benefits of Cloudflare.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/data-sources/dns_record#proxied DataCloudflareDnsRecord#proxied}
	Proxied interface{} `field:"optional" json:"proxied" yaml:"proxied"`
	// Allows searching in multiple properties of a DNS record simultaneously.
	//
	// This parameter is intended for human users, not automation. Its exact behavior is intentionally left unspecified and is subject to change in the future. This parameter works independently of the `match` setting. For automated searches, please use the other available parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/data-sources/dns_record#search DataCloudflareDnsRecord#search}
	Search *string `field:"optional" json:"search" yaml:"search"`
	// Filters to records at or below the given NS delegation name, excluding the NS records that form the delegation itself.
	//
	// The value must be a subdomain of the zone; the zone apex is not accepted. Requires `include_shadow_metadata=true`. See [Shadowed records](https://developers.cloudflare.com/dns/manage-dns-records/reference/shadowed-records).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/data-sources/dns_record#shadowed_by_name DataCloudflareDnsRecord#shadowed_by_name}
	ShadowedByName *string `field:"optional" json:"shadowedByName" yaml:"shadowedByName"`
	// Returns NS records that shadow the given name, searching at the name itself and each of its ancestor names within the zone, excluding the zone apex.
	//
	// The value must be a subdomain of the zone; the zone apex is not accepted. See [Shadowed records](https://developers.cloudflare.com/dns/manage-dns-records/reference/shadowed-records).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/data-sources/dns_record#shadowing_name DataCloudflareDnsRecord#shadowing_name}
	ShadowingName *string `field:"optional" json:"shadowingName" yaml:"shadowingName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/data-sources/dns_record#tag DataCloudflareDnsRecord#tag}.
	Tag *DataCloudflareDnsRecordFilterTag `field:"optional" json:"tag" yaml:"tag"`
	// Whether to match all tag search requirements or at least one (any).
	//
	// If set to `all`, acts like a logical AND between tag filters. If set to `any`, acts like a logical OR instead. Note that the regular `match` parameter is still used to combine the resulting condition with other filters that aren't related to tags.
	// Available values: "any", "all".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/data-sources/dns_record#tag_match DataCloudflareDnsRecord#tag_match}
	TagMatch *string `field:"optional" json:"tagMatch" yaml:"tagMatch"`
	// Record type.
	//
	// Available values: "A", "AAAA", "CAA", "CERT", "CNAME", "DNSKEY", "DS", "HTTPS", "LOC", "MX", "NAPTR", "NS", "OPENPGPKEY", "PTR", "SMIMEA", "SRV", "SSHFP", "SVCB", "TLSA", "TXT", "URI".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/data-sources/dns_record#type DataCloudflareDnsRecord#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

