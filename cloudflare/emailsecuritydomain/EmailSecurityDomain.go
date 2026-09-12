// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package emailsecuritydomain

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v16/emailsecuritydomain/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.25.0/docs/resources/email_security_domain cloudflare_email_security_domain}.
type EmailSecurityDomain interface {
	cdktn.TerraformResource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	AllowedDeliveryModes() *[]*string
	SetAllowedDeliveryModes(val *[]*string)
	AllowedDeliveryModesInput() *[]*string
	Authorization() EmailSecurityDomainAuthorizationOutputReference
	// Experimental.
	CdktfStack() cdktn.TerraformStack
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
	DmarcStatus() *string
	Domain() *string
	SetDomain(val *string)
	DomainInput() *string
	DropDispositions() *[]*string
	SetDropDispositions(val *[]*string)
	DropDispositionsInput() *[]*string
	EmailsProcessed() EmailSecurityDomainEmailsProcessedOutputReference
	Folder() *string
	SetFolder(val *string)
	FolderInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	InboxProvider() *string
	IntegrationId() *string
	SetIntegrationId(val *string)
	IntegrationIdInput() *string
	IpRestrictions() *[]*string
	SetIpRestrictions(val *[]*string)
	IpRestrictionsInput() *[]*string
	LastModified() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	LookbackHops() *float64
	SetLookbackHops(val *float64)
	LookbackHopsInput() *float64
	ModifiedAt() *string
	// The tree node.
	Node() constructs.Node
	O365TenantId() *string
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
	Regions() *[]*string
	SetRegions(val *[]*string)
	RegionsInput() *[]*string
	RequireTlsInbound() interface{}
	SetRequireTlsInbound(val interface{})
	RequireTlsInboundInput() interface{}
	RequireTlsOutbound() interface{}
	SetRequireTlsOutbound(val interface{})
	RequireTlsOutboundInput() interface{}
	SpfStatus() *string
	Status() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Transport() *string
	SetTransport(val *string)
	TransportInput() *string
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
	ResetFolder()
	ResetIntegrationId()
	ResetLookbackHops()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRequireTlsInbound()
	ResetRequireTlsOutbound()
	ResetTransport()
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

// The jsii proxy struct for EmailSecurityDomain
type jsiiProxy_EmailSecurityDomain struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_EmailSecurityDomain) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityDomain) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityDomain) AllowedDeliveryModes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedDeliveryModes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityDomain) AllowedDeliveryModesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedDeliveryModesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityDomain) Authorization() EmailSecurityDomainAuthorizationOutputReference {
	var returns EmailSecurityDomainAuthorizationOutputReference
	_jsii_.Get(
		j,
		"authorization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityDomain) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityDomain) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityDomain) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityDomain) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityDomain) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityDomain) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityDomain) DmarcStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dmarcStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityDomain) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityDomain) DomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityDomain) DropDispositions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dropDispositions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityDomain) DropDispositionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dropDispositionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityDomain) EmailsProcessed() EmailSecurityDomainEmailsProcessedOutputReference {
	var returns EmailSecurityDomainEmailsProcessedOutputReference
	_jsii_.Get(
		j,
		"emailsProcessed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityDomain) Folder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityDomain) FolderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityDomain) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityDomain) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityDomain) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityDomain) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityDomain) InboxProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inboxProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityDomain) IntegrationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityDomain) IntegrationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityDomain) IpRestrictions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipRestrictions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityDomain) IpRestrictionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipRestrictionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityDomain) LastModified() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModified",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityDomain) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityDomain) LookbackHops() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lookbackHops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityDomain) LookbackHopsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lookbackHopsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityDomain) ModifiedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modifiedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityDomain) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityDomain) O365TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"o365TenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityDomain) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityDomain) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityDomain) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityDomain) Regions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityDomain) RegionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityDomain) RequireTlsInbound() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireTlsInbound",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityDomain) RequireTlsInboundInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireTlsInboundInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityDomain) RequireTlsOutbound() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireTlsOutbound",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityDomain) RequireTlsOutboundInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireTlsOutboundInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityDomain) SpfStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spfStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityDomain) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityDomain) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityDomain) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityDomain) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityDomain) Transport() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityDomain) TransportInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transportInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.25.0/docs/resources/email_security_domain cloudflare_email_security_domain} Resource.
func NewEmailSecurityDomain(scope constructs.Construct, id *string, config *EmailSecurityDomainConfig) EmailSecurityDomain {
	_init_.Initialize()

	if err := validateNewEmailSecurityDomainParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_EmailSecurityDomain{}

	_jsii_.Create(
		"@cdktn/provider-cloudflare.emailSecurityDomain.EmailSecurityDomain",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.25.0/docs/resources/email_security_domain cloudflare_email_security_domain} Resource.
func NewEmailSecurityDomain_Override(e EmailSecurityDomain, scope constructs.Construct, id *string, config *EmailSecurityDomainConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-cloudflare.emailSecurityDomain.EmailSecurityDomain",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_EmailSecurityDomain)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityDomain)SetAllowedDeliveryModes(val *[]*string) {
	if err := j.validateSetAllowedDeliveryModesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedDeliveryModes",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityDomain)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityDomain)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityDomain)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityDomain)SetDomain(val *string) {
	if err := j.validateSetDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityDomain)SetDropDispositions(val *[]*string) {
	if err := j.validateSetDropDispositionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dropDispositions",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityDomain)SetFolder(val *string) {
	if err := j.validateSetFolderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"folder",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityDomain)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityDomain)SetIntegrationId(val *string) {
	if err := j.validateSetIntegrationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"integrationId",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityDomain)SetIpRestrictions(val *[]*string) {
	if err := j.validateSetIpRestrictionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipRestrictions",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityDomain)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityDomain)SetLookbackHops(val *float64) {
	if err := j.validateSetLookbackHopsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lookbackHops",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityDomain)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityDomain)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityDomain)SetRegions(val *[]*string) {
	if err := j.validateSetRegionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regions",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityDomain)SetRequireTlsInbound(val interface{}) {
	if err := j.validateSetRequireTlsInboundParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireTlsInbound",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityDomain)SetRequireTlsOutbound(val interface{}) {
	if err := j.validateSetRequireTlsOutboundParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireTlsOutbound",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityDomain)SetTransport(val *string) {
	if err := j.validateSetTransportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transport",
		val,
	)
}

// Generates CDKTN code for importing a EmailSecurityDomain resource upon running "cdktn plan <stack-name>".
func EmailSecurityDomain_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateEmailSecurityDomain_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-cloudflare.emailSecurityDomain.EmailSecurityDomain",
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
func EmailSecurityDomain_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEmailSecurityDomain_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-cloudflare.emailSecurityDomain.EmailSecurityDomain",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EmailSecurityDomain_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEmailSecurityDomain_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-cloudflare.emailSecurityDomain.EmailSecurityDomain",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EmailSecurityDomain_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEmailSecurityDomain_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-cloudflare.emailSecurityDomain.EmailSecurityDomain",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func EmailSecurityDomain_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-cloudflare.emailSecurityDomain.EmailSecurityDomain",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_EmailSecurityDomain) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_EmailSecurityDomain) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_EmailSecurityDomain) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailSecurityDomain) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailSecurityDomain) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailSecurityDomain) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailSecurityDomain) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailSecurityDomain) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailSecurityDomain) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailSecurityDomain) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailSecurityDomain) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailSecurityDomain) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailSecurityDomain) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_EmailSecurityDomain) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailSecurityDomain) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := e.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailSecurityDomain) MoveFromId(id *string) {
	if err := e.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveFromId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_EmailSecurityDomain) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_EmailSecurityDomain) MoveToId(id *string) {
	if err := e.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveToId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_EmailSecurityDomain) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_EmailSecurityDomain) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := e.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (e *jsiiProxy_EmailSecurityDomain) ResetFolder() {
	_jsii_.InvokeVoid(
		e,
		"resetFolder",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmailSecurityDomain) ResetIntegrationId() {
	_jsii_.InvokeVoid(
		e,
		"resetIntegrationId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmailSecurityDomain) ResetLookbackHops() {
	_jsii_.InvokeVoid(
		e,
		"resetLookbackHops",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmailSecurityDomain) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmailSecurityDomain) ResetRequireTlsInbound() {
	_jsii_.InvokeVoid(
		e,
		"resetRequireTlsInbound",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmailSecurityDomain) ResetRequireTlsOutbound() {
	_jsii_.InvokeVoid(
		e,
		"resetRequireTlsOutbound",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmailSecurityDomain) ResetTransport() {
	_jsii_.InvokeVoid(
		e,
		"resetTransport",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmailSecurityDomain) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailSecurityDomain) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailSecurityDomain) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailSecurityDomain) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailSecurityDomain) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailSecurityDomain) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailSecurityDomain) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		e,
		"with",
		args,
		&returns,
	)

	return returns
}

