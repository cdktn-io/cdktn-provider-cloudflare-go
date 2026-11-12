// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apishield


type ApiShieldAuthIdCharacteristics struct {
	// For `type: header` or `type: cookie`: The name of the characteristic field, i.e., the header or cookie name.
	//
	// For `type: jwt`: Claim location expressed as `$(token_config_id):$(json_path)`, where `token_config_id`
	// is the ID of the token configuration used in validating the JWT, and `json_path` is a RFC 9535
	// JSONPath (https://goessner.net/articles/JsonPath/, https://www.rfc-editor.org/rfc/rfc9535.html).
	// The JSONPath expression may be in dot or bracket notation, may only specify literal keys
	// or array indexes, and must return a singleton value, which will be interpreted as a string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/api_shield#name ApiShield#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of characteristic. Available values: "header", "cookie", "jwt".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/api_shield#type ApiShield#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

