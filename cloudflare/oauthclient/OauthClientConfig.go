// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oauthclient

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OauthClientConfig struct {
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
	// Account identifier tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.0/docs/resources/oauth_client#account_id OauthClient#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Human-readable name of the OAuth client.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.0/docs/resources/oauth_client#client_name OauthClient#client_name}
	ClientName *string `field:"required" json:"clientName" yaml:"clientName"`
	// Array of OAuth grant types the client is allowed to use. `authorization_code` is required; `refresh_token` may be included optionally.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.0/docs/resources/oauth_client#grant_types OauthClient#grant_types}
	GrantTypes *[]*string `field:"required" json:"grantTypes" yaml:"grantTypes"`
	// Array of allowed redirect URIs for the client.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.0/docs/resources/oauth_client#redirect_uris OauthClient#redirect_uris}
	RedirectUris *[]*string `field:"required" json:"redirectUris" yaml:"redirectUris"`
	// Array of OAuth response types the client is allowed to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.0/docs/resources/oauth_client#response_types OauthClient#response_types}
	ResponseTypes *[]*string `field:"required" json:"responseTypes" yaml:"responseTypes"`
	// Array of OAuth scopes the client is allowed to request.
	//
	// Colon-delimited scopes are not accepted. Dot-delimited scopes are validated against available OAuth API scopes; simple identity scopes are allowed. Protocol scopes `offline_access` and `openid` are added or removed automatically based on `grant_types` and `response_types`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.0/docs/resources/oauth_client#scopes OauthClient#scopes}
	Scopes *[]*string `field:"required" json:"scopes" yaml:"scopes"`
	// The authentication method the client uses at the token endpoint. Available values: "none", "client_secret_basic", "client_secret_post".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.0/docs/resources/oauth_client#token_endpoint_auth_method OauthClient#token_endpoint_auth_method}
	TokenEndpointAuthMethod *string `field:"required" json:"tokenEndpointAuthMethod" yaml:"tokenEndpointAuthMethod"`
	// Array of allowed CORS origins.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.0/docs/resources/oauth_client#allowed_cors_origins OauthClient#allowed_cors_origins}
	AllowedCorsOrigins *[]*string `field:"optional" json:"allowedCorsOrigins" yaml:"allowedCorsOrigins"`
	// URL of the home page of the client.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.0/docs/resources/oauth_client#client_uri OauthClient#client_uri}
	ClientUri *string `field:"optional" json:"clientUri" yaml:"clientUri"`
	// URL of the client's logo.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.0/docs/resources/oauth_client#logo_uri OauthClient#logo_uri}
	LogoUri *string `field:"optional" json:"logoUri" yaml:"logoUri"`
	// The unique identifier for an OAuth client.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.0/docs/resources/oauth_client#oauth_client_id OauthClient#oauth_client_id}
	OauthClientId *string `field:"optional" json:"oauthClientId" yaml:"oauthClientId"`
	// URL that points to a privacy policy document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.0/docs/resources/oauth_client#policy_uri OauthClient#policy_uri}
	PolicyUri *string `field:"optional" json:"policyUri" yaml:"policyUri"`
	// Array of allowed post-logout redirect URIs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.0/docs/resources/oauth_client#post_logout_redirect_uris OauthClient#post_logout_redirect_uris}
	PostLogoutRedirectUris *[]*string `field:"optional" json:"postLogoutRedirectUris" yaml:"postLogoutRedirectUris"`
	// URL that points to a terms of service document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.0/docs/resources/oauth_client#tos_uri OauthClient#tos_uri}
	TosUri *string `field:"optional" json:"tosUri" yaml:"tosUri"`
	// Promote the OAuth client from private to public visibility.
	//
	// Only `public` is accepted; demotion to `private` is not supported. Promotion requires a non-empty client name, logo URI, verified client URI host, and at least one non-identity scope.
	// Available values: "public".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.0/docs/resources/oauth_client#visibility OauthClient#visibility}
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

