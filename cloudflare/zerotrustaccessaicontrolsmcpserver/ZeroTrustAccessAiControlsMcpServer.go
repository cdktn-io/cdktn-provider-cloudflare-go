// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessaicontrolsmcpserver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v16/zerotrustaccessaicontrolsmcpserver/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/zero_trust_access_ai_controls_mcp_server cloudflare_zero_trust_access_ai_controls_mcp_server}.
type ZeroTrustAccessAiControlsMcpServer interface {
	cdktn.TerraformResource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	AuthConfigSummary() ZeroTrustAccessAiControlsMcpServerAuthConfigSummaryOutputReference
	AuthCredentials() *string
	SetAuthCredentials(val *string)
	AuthCredentialsInput() *string
	AuthenticationStatus() *string
	AuthType() *string
	SetAuthType(val *string)
	AuthTypeInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	ClientSecret() *string
	SetClientSecret(val *string)
	ClientSecretInput() *string
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
	CreatedBy() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Error() *string
	ErrorDetails() ZeroTrustAccessAiControlsMcpServerErrorDetailsOutputReference
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Hostname() *string
	SetHostname(val *string)
	HostnameInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IsSharedOauthCallbackEnabled() interface{}
	SetIsSharedOauthCallbackEnabled(val interface{})
	IsSharedOauthCallbackEnabledInput() interface{}
	LastSuccessfulSync() *string
	LastSynced() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	ModifiedAt() *string
	ModifiedBy() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Prompts() cdktn.StringMapList
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
	SecureWebGateway() interface{}
	SetSecureWebGateway(val interface{})
	SecureWebGatewayInput() interface{}
	Status() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Tools() cdktn.StringMapList
	UpdatedPrompts() ZeroTrustAccessAiControlsMcpServerUpdatedPromptsList
	UpdatedPromptsInput() interface{}
	UpdatedTools() ZeroTrustAccessAiControlsMcpServerUpdatedToolsList
	UpdatedToolsInput() interface{}
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
	PutUpdatedPrompts(value interface{})
	PutUpdatedTools(value interface{})
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
	ResetAuthCredentials()
	ResetClientSecret()
	ResetDescription()
	ResetIsSharedOauthCallbackEnabled()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSecureWebGateway()
	ResetUpdatedPrompts()
	ResetUpdatedTools()
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

// The jsii proxy struct for ZeroTrustAccessAiControlsMcpServer
type jsiiProxy_ZeroTrustAccessAiControlsMcpServer struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) AuthConfigSummary() ZeroTrustAccessAiControlsMcpServerAuthConfigSummaryOutputReference {
	var returns ZeroTrustAccessAiControlsMcpServerAuthConfigSummaryOutputReference
	_jsii_.Get(
		j,
		"authConfigSummary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) AuthCredentials() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) AuthCredentialsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) AuthenticationStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) AuthType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) AuthTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) ClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) CreatedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) Error() *string {
	var returns *string
	_jsii_.Get(
		j,
		"error",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) ErrorDetails() ZeroTrustAccessAiControlsMcpServerErrorDetailsOutputReference {
	var returns ZeroTrustAccessAiControlsMcpServerErrorDetailsOutputReference
	_jsii_.Get(
		j,
		"errorDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) HostnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) IsSharedOauthCallbackEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSharedOauthCallbackEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) IsSharedOauthCallbackEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSharedOauthCallbackEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) LastSuccessfulSync() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastSuccessfulSync",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) LastSynced() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastSynced",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) ModifiedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modifiedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) ModifiedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modifiedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) Prompts() cdktn.StringMapList {
	var returns cdktn.StringMapList
	_jsii_.Get(
		j,
		"prompts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) SecureWebGateway() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secureWebGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) SecureWebGatewayInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secureWebGatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) Tools() cdktn.StringMapList {
	var returns cdktn.StringMapList
	_jsii_.Get(
		j,
		"tools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) UpdatedPrompts() ZeroTrustAccessAiControlsMcpServerUpdatedPromptsList {
	var returns ZeroTrustAccessAiControlsMcpServerUpdatedPromptsList
	_jsii_.Get(
		j,
		"updatedPrompts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) UpdatedPromptsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"updatedPromptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) UpdatedTools() ZeroTrustAccessAiControlsMcpServerUpdatedToolsList {
	var returns ZeroTrustAccessAiControlsMcpServerUpdatedToolsList
	_jsii_.Get(
		j,
		"updatedTools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) UpdatedToolsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"updatedToolsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/zero_trust_access_ai_controls_mcp_server cloudflare_zero_trust_access_ai_controls_mcp_server} Resource.
func NewZeroTrustAccessAiControlsMcpServer(scope constructs.Construct, id *string, config *ZeroTrustAccessAiControlsMcpServerConfig) ZeroTrustAccessAiControlsMcpServer {
	_init_.Initialize()

	if err := validateNewZeroTrustAccessAiControlsMcpServerParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZeroTrustAccessAiControlsMcpServer{}

	_jsii_.Create(
		"@cdktn/provider-cloudflare.zeroTrustAccessAiControlsMcpServer.ZeroTrustAccessAiControlsMcpServer",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/zero_trust_access_ai_controls_mcp_server cloudflare_zero_trust_access_ai_controls_mcp_server} Resource.
func NewZeroTrustAccessAiControlsMcpServer_Override(z ZeroTrustAccessAiControlsMcpServer, scope constructs.Construct, id *string, config *ZeroTrustAccessAiControlsMcpServerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-cloudflare.zeroTrustAccessAiControlsMcpServer.ZeroTrustAccessAiControlsMcpServer",
		[]interface{}{scope, id, config},
		z,
	)
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer)SetAuthCredentials(val *string) {
	if err := j.validateSetAuthCredentialsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authCredentials",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer)SetAuthType(val *string) {
	if err := j.validateSetAuthTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authType",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer)SetClientSecret(val *string) {
	if err := j.validateSetClientSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSecret",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer)SetHostname(val *string) {
	if err := j.validateSetHostnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostname",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer)SetIsSharedOauthCallbackEnabled(val interface{}) {
	if err := j.validateSetIsSharedOauthCallbackEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isSharedOauthCallbackEnabled",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessAiControlsMcpServer)SetSecureWebGateway(val interface{}) {
	if err := j.validateSetSecureWebGatewayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secureWebGateway",
		val,
	)
}

// Generates CDKTN code for importing a ZeroTrustAccessAiControlsMcpServer resource upon running "cdktn plan <stack-name>".
func ZeroTrustAccessAiControlsMcpServer_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateZeroTrustAccessAiControlsMcpServer_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-cloudflare.zeroTrustAccessAiControlsMcpServer.ZeroTrustAccessAiControlsMcpServer",
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
func ZeroTrustAccessAiControlsMcpServer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateZeroTrustAccessAiControlsMcpServer_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-cloudflare.zeroTrustAccessAiControlsMcpServer.ZeroTrustAccessAiControlsMcpServer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ZeroTrustAccessAiControlsMcpServer_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateZeroTrustAccessAiControlsMcpServer_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-cloudflare.zeroTrustAccessAiControlsMcpServer.ZeroTrustAccessAiControlsMcpServer",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ZeroTrustAccessAiControlsMcpServer_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateZeroTrustAccessAiControlsMcpServer_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-cloudflare.zeroTrustAccessAiControlsMcpServer.ZeroTrustAccessAiControlsMcpServer",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ZeroTrustAccessAiControlsMcpServer_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-cloudflare.zeroTrustAccessAiControlsMcpServer.ZeroTrustAccessAiControlsMcpServer",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (z *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) AddMoveTarget(moveTarget *string) {
	if err := z.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (z *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) AddOverride(path *string, value interface{}) {
	if err := z.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := z.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		z,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := z.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		z,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := z.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		z,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := z.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		z,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := z.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		z,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := z.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		z,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := z.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		z,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) GetStringAttribute(terraformAttribute *string) *string {
	if err := z.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		z,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := z.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		z,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		z,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := z.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (z *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := z.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := z.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		z,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) MoveFromId(id *string) {
	if err := z.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"moveFromId",
		[]interface{}{id},
	)
}

func (z *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) MoveTo(moveTarget *string, index interface{}) {
	if err := z.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (z *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) MoveToId(id *string) {
	if err := z.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"moveToId",
		[]interface{}{id},
	)
}

func (z *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) OverrideLogicalId(newLogicalId *string) {
	if err := z.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (z *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) PutUpdatedPrompts(value interface{}) {
	if err := z.validatePutUpdatedPromptsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putUpdatedPrompts",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) PutUpdatedTools(value interface{}) {
	if err := z.validatePutUpdatedToolsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putUpdatedTools",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := z.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (z *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) ResetAuthCredentials() {
	_jsii_.InvokeVoid(
		z,
		"resetAuthCredentials",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) ResetClientSecret() {
	_jsii_.InvokeVoid(
		z,
		"resetClientSecret",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) ResetDescription() {
	_jsii_.InvokeVoid(
		z,
		"resetDescription",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) ResetIsSharedOauthCallbackEnabled() {
	_jsii_.InvokeVoid(
		z,
		"resetIsSharedOauthCallbackEnabled",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		z,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) ResetSecureWebGateway() {
	_jsii_.InvokeVoid(
		z,
		"resetSecureWebGateway",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) ResetUpdatedPrompts() {
	_jsii_.InvokeVoid(
		z,
		"resetUpdatedPrompts",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) ResetUpdatedTools() {
	_jsii_.InvokeVoid(
		z,
		"resetUpdatedTools",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		z,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		z,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		z,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		z,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		z,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessAiControlsMcpServer) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		z,
		"with",
		args,
		&returns,
	)

	return returns
}

