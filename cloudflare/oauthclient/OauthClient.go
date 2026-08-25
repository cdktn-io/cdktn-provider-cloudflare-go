// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oauthclient

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v16/oauthclient/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/oauth_client cloudflare_oauth_client}.
type OauthClient interface {
	cdktn.TerraformResource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	AllowedCorsOrigins() *[]*string
	SetAllowedCorsOrigins(val *[]*string)
	AllowedCorsOriginsInput() *[]*string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	ClientId() *string
	ClientName() *string
	SetClientName(val *string)
	ClientNameInput() *string
	ClientSecret() *string
	ClientUri() *string
	SetClientUri(val *string)
	ClientUriInput() *string
	ClientUriVerification() OauthClientClientUriVerificationOutputReference
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreatedAt() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GrantTypes() *[]*string
	SetGrantTypes(val *[]*string)
	GrantTypesInput() *[]*string
	HasRotatedSecret() cdktn.IResolvable
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	LogoUri() *string
	SetLogoUri(val *string)
	LogoUriInput() *string
	// The tree node.
	Node() constructs.Node
	OauthClientId() *string
	SetOauthClientId(val *string)
	OauthClientIdInput() *string
	PolicyUri() *string
	SetPolicyUri(val *string)
	PolicyUriInput() *string
	PostLogoutRedirectUris() *[]*string
	SetPostLogoutRedirectUris(val *[]*string)
	PostLogoutRedirectUrisInput() *[]*string
	PromotedAt() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	RedirectUris() *[]*string
	SetRedirectUris(val *[]*string)
	RedirectUrisInput() *[]*string
	ResponseTypes() *[]*string
	SetResponseTypes(val *[]*string)
	ResponseTypesInput() *[]*string
	Scopes() *[]*string
	SetScopes(val *[]*string)
	ScopesInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TokenEndpointAuthMethod() *string
	SetTokenEndpointAuthMethod(val *string)
	TokenEndpointAuthMethodInput() *string
	TosUri() *string
	SetTosUri(val *string)
	TosUriInput() *string
	UpdatedAt() *string
	Visibility() *string
	SetVisibility(val *string)
	VisibilityInput() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktn.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Wraps a write-only attribute's already-mapped value so that `ProviderFeature.WRITE_ONLY_ATTRIBUTES` usage is registered at *resolve* time instead of at mutation time (setter/constructor). Called by generated bindings from `synthesizeAttributes()` and `synthesizeHclAttributes()`, e.g. `secret_key_wo: this.markWriteOnlyAttribute(cdktn.stringToTerraform(this._secretKeyWo))`; not intended to be called directly.
	//
	// `undefined` passes through completely unchanged, so the existing
	// undefined-filtering that omits unset attributes from synthesized
	// output (see `resolve()` in `tokens/private/resolve.ts`, and the
	// `value.value !== undefined` filter in generated
	// `synthesizeHclAttributes()`) keeps working untouched. `null` is also
	// passed through unchanged: it already renders as an explicit
	// null-out and must not arm the validation either.
	//
	// Any other value - including one that will itself resolve to nothing
	// (e.g. a `Lazy`/`IResolvable` producer with no value to contribute) -
	// is wrapped in a token whose `resolve()` defers to the real resolver
	// first and registers usage only if what comes back is not
	// `null`/`undefined`; the resolved value is then returned unchanged,
	// so what actually renders is untouched by this wrapper. A producer
	// that resolves to `undefined` therefore neither registers usage nor
	// leaves anything behind in the synthesized attribute - the omission
	// behaves exactly as if the attribute had never been set.
	//
	// Registration goes through `_registerResolveDiscoveredProviderFeatureUsage`
	// rather than `registerProviderFeatureUsage`: usage here is only known at
	// resolve time, and a given element can be resolved across many
	// synthesis passes over its lifetime (repeated `app.synth()` calls,
	// tests reusing a construct tree), so it must represent only the CURRENT
	// pass rather than accumulate forever. Every validation-enabled entry
	// point (`App.synth`; `Testing.synth`/`synthHcl` with validations;
	// `StackSynthesizer.synthesize`) runs a prepare step that deactivates any
	// stale registration and then resolves every element's `toTerraform()`
	// before that same entry point's validations run - see
	// `TerraformStack._runPreparingResolve` - so whatever this closure
	// (re-)registers during that prepare step is always visible to the
	// validation that reads it afterwards, and nothing left over from an
	// earlier pass leaks into the current one.
	// Experimental.
	MarkWriteOnlyAttribute(value interface{}) interface{}
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using its instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Registers a synth-time validation that the project's declared targetVersions admit the given provider-protocol feature family.
	//
	// Called by generated provider bindings when a versioned feature is
	// structurally in use - the element's existence in the construct tree
	// already implies the feature is used, e.g. constructing a
	// `TerraformEphemeralResource` at all - so, unlike
	// `_registerResolveDiscoveredProviderFeatureUsage`, this registration is
	// never deactivated by `_resetResolveDiscoveredProviderFeatureUsage`. Not
	// intended to be called directly by user code. Lives on `TerraformElement`
	// (rather than `TerraformResource`) so it covers any element subclass
	// that needs it.
	// Experimental.
	RegisterProviderFeatureUsage(feature cdktn.ProviderFeature)
	ResetAllowedCorsOrigins()
	ResetClientUri()
	ResetLogoUri()
	ResetOauthClientId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPolicyUri()
	ResetPostLogoutRedirectUris()
	ResetTosUri()
	ResetVisibility()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for OauthClient
type jsiiProxy_OauthClient struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_OauthClient) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthClient) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthClient) AllowedCorsOrigins() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedCorsOrigins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthClient) AllowedCorsOriginsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedCorsOriginsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthClient) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthClient) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthClient) ClientName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthClient) ClientNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthClient) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthClient) ClientUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthClient) ClientUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthClient) ClientUriVerification() OauthClientClientUriVerificationOutputReference {
	var returns OauthClientClientUriVerificationOutputReference
	_jsii_.Get(
		j,
		"clientUriVerification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthClient) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthClient) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthClient) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthClient) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthClient) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthClient) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthClient) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthClient) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthClient) GrantTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"grantTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthClient) GrantTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"grantTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthClient) HasRotatedSecret() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"hasRotatedSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthClient) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthClient) LogoUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthClient) LogoUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthClient) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthClient) OauthClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthClientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthClient) OauthClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthClientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthClient) PolicyUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthClient) PolicyUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthClient) PostLogoutRedirectUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"postLogoutRedirectUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthClient) PostLogoutRedirectUrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"postLogoutRedirectUrisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthClient) PromotedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"promotedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthClient) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthClient) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthClient) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthClient) RedirectUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"redirectUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthClient) RedirectUrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"redirectUrisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthClient) ResponseTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"responseTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthClient) ResponseTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"responseTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthClient) Scopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthClient) ScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthClient) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthClient) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthClient) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthClient) TokenEndpointAuthMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenEndpointAuthMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthClient) TokenEndpointAuthMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenEndpointAuthMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthClient) TosUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tosUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthClient) TosUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tosUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthClient) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthClient) Visibility() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthClient) VisibilityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visibilityInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/oauth_client cloudflare_oauth_client} Resource.
func NewOauthClient(scope constructs.Construct, id *string, config *OauthClientConfig) OauthClient {
	_init_.Initialize()

	if err := validateNewOauthClientParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_OauthClient{}

	_jsii_.Create(
		"@cdktn/provider-cloudflare.oauthClient.OauthClient",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/oauth_client cloudflare_oauth_client} Resource.
func NewOauthClient_Override(o OauthClient, scope constructs.Construct, id *string, config *OauthClientConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-cloudflare.oauthClient.OauthClient",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OauthClient)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_OauthClient)SetAllowedCorsOrigins(val *[]*string) {
	if err := j.validateSetAllowedCorsOriginsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedCorsOrigins",
		val,
	)
}

func (j *jsiiProxy_OauthClient)SetClientName(val *string) {
	if err := j.validateSetClientNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientName",
		val,
	)
}

func (j *jsiiProxy_OauthClient)SetClientUri(val *string) {
	if err := j.validateSetClientUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientUri",
		val,
	)
}

func (j *jsiiProxy_OauthClient)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_OauthClient)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OauthClient)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OauthClient)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_OauthClient)SetGrantTypes(val *[]*string) {
	if err := j.validateSetGrantTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"grantTypes",
		val,
	)
}

func (j *jsiiProxy_OauthClient)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OauthClient)SetLogoUri(val *string) {
	if err := j.validateSetLogoUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logoUri",
		val,
	)
}

func (j *jsiiProxy_OauthClient)SetOauthClientId(val *string) {
	if err := j.validateSetOauthClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauthClientId",
		val,
	)
}

func (j *jsiiProxy_OauthClient)SetPolicyUri(val *string) {
	if err := j.validateSetPolicyUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyUri",
		val,
	)
}

func (j *jsiiProxy_OauthClient)SetPostLogoutRedirectUris(val *[]*string) {
	if err := j.validateSetPostLogoutRedirectUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postLogoutRedirectUris",
		val,
	)
}

func (j *jsiiProxy_OauthClient)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OauthClient)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_OauthClient)SetRedirectUris(val *[]*string) {
	if err := j.validateSetRedirectUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redirectUris",
		val,
	)
}

func (j *jsiiProxy_OauthClient)SetResponseTypes(val *[]*string) {
	if err := j.validateSetResponseTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"responseTypes",
		val,
	)
}

func (j *jsiiProxy_OauthClient)SetScopes(val *[]*string) {
	if err := j.validateSetScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scopes",
		val,
	)
}

func (j *jsiiProxy_OauthClient)SetTokenEndpointAuthMethod(val *string) {
	if err := j.validateSetTokenEndpointAuthMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenEndpointAuthMethod",
		val,
	)
}

func (j *jsiiProxy_OauthClient)SetTosUri(val *string) {
	if err := j.validateSetTosUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tosUri",
		val,
	)
}

func (j *jsiiProxy_OauthClient)SetVisibility(val *string) {
	if err := j.validateSetVisibilityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visibility",
		val,
	)
}

// Generates CDKTN code for importing a OauthClient resource upon running "cdktn plan <stack-name>".
func OauthClient_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateOauthClient_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-cloudflare.oauthClient.OauthClient",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func OauthClient_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOauthClient_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-cloudflare.oauthClient.OauthClient",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OauthClient_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOauthClient_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-cloudflare.oauthClient.OauthClient",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OauthClient_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOauthClient_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-cloudflare.oauthClient.OauthClient",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OauthClient_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-cloudflare.oauthClient.OauthClient",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OauthClient) AddMoveTarget(moveTarget *string) {
	if err := o.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (o *jsiiProxy_OauthClient) AddOverride(path *string, value interface{}) {
	if err := o.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (o *jsiiProxy_OauthClient) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthClient) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthClient) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthClient) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthClient) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthClient) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthClient) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthClient) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthClient) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthClient) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthClient) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := o.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (o *jsiiProxy_OauthClient) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthClient) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := o.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthClient) MoveFromId(id *string) {
	if err := o.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveFromId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_OauthClient) MoveTo(moveTarget *string, index interface{}) {
	if err := o.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (o *jsiiProxy_OauthClient) MoveToId(id *string) {
	if err := o.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveToId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_OauthClient) OverrideLogicalId(newLogicalId *string) {
	if err := o.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OauthClient) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := o.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (o *jsiiProxy_OauthClient) ResetAllowedCorsOrigins() {
	_jsii_.InvokeVoid(
		o,
		"resetAllowedCorsOrigins",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthClient) ResetClientUri() {
	_jsii_.InvokeVoid(
		o,
		"resetClientUri",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthClient) ResetLogoUri() {
	_jsii_.InvokeVoid(
		o,
		"resetLogoUri",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthClient) ResetOauthClientId() {
	_jsii_.InvokeVoid(
		o,
		"resetOauthClientId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthClient) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthClient) ResetPolicyUri() {
	_jsii_.InvokeVoid(
		o,
		"resetPolicyUri",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthClient) ResetPostLogoutRedirectUris() {
	_jsii_.InvokeVoid(
		o,
		"resetPostLogoutRedirectUris",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthClient) ResetTosUri() {
	_jsii_.InvokeVoid(
		o,
		"resetTosUri",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthClient) ResetVisibility() {
	_jsii_.InvokeVoid(
		o,
		"resetVisibility",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthClient) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthClient) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthClient) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthClient) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthClient) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthClient) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthClient) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		o,
		"with",
		args,
		&returns,
	)

	return returns
}

