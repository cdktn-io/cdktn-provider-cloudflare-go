// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccesspolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v16/zerotrustaccesspolicy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/zero_trust_access_policy cloudflare_zero_trust_access_policy}.
type ZeroTrustAccessPolicy interface {
	cdktn.TerraformResource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	AppCount() *float64
	ApprovalGroups() ZeroTrustAccessPolicyApprovalGroupsList
	ApprovalGroupsInput() interface{}
	ApprovalRequired() interface{}
	SetApprovalRequired(val interface{})
	ApprovalRequiredInput() interface{}
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	ConnectionRules() ZeroTrustAccessPolicyConnectionRulesOutputReference
	ConnectionRulesInput() interface{}
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreatedAt() *string
	Decision() *string
	SetDecision(val *string)
	DecisionInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Exclude() ZeroTrustAccessPolicyExcludeList
	ExcludeInput() interface{}
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	Include() ZeroTrustAccessPolicyIncludeList
	IncludeInput() interface{}
	IsolationRequired() interface{}
	SetIsolationRequired(val interface{})
	IsolationRequiredInput() interface{}
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	MfaConfig() ZeroTrustAccessPolicyMfaConfigOutputReference
	MfaConfigInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PurposeJustificationPrompt() *string
	SetPurposeJustificationPrompt(val *string)
	PurposeJustificationPromptInput() *string
	PurposeJustificationRequired() interface{}
	SetPurposeJustificationRequired(val interface{})
	PurposeJustificationRequiredInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	Require() ZeroTrustAccessPolicyRequireList
	RequireInput() interface{}
	Reusable() cdktn.IResolvable
	SessionDuration() *string
	SetSessionDuration(val *string)
	SessionDurationInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UpdatedAt() *string
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
	PutApprovalGroups(value interface{})
	PutConnectionRules(value *ZeroTrustAccessPolicyConnectionRules)
	PutExclude(value interface{})
	PutInclude(value interface{})
	PutMfaConfig(value *ZeroTrustAccessPolicyMfaConfig)
	PutRequire(value interface{})
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
	ResetApprovalGroups()
	ResetApprovalRequired()
	ResetConnectionRules()
	ResetExclude()
	ResetInclude()
	ResetIsolationRequired()
	ResetMfaConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPurposeJustificationPrompt()
	ResetPurposeJustificationRequired()
	ResetRequire()
	ResetSessionDuration()
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

// The jsii proxy struct for ZeroTrustAccessPolicy
type jsiiProxy_ZeroTrustAccessPolicy struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) AppCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"appCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) ApprovalGroups() ZeroTrustAccessPolicyApprovalGroupsList {
	var returns ZeroTrustAccessPolicyApprovalGroupsList
	_jsii_.Get(
		j,
		"approvalGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) ApprovalGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"approvalGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) ApprovalRequired() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"approvalRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) ApprovalRequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"approvalRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) ConnectionRules() ZeroTrustAccessPolicyConnectionRulesOutputReference {
	var returns ZeroTrustAccessPolicyConnectionRulesOutputReference
	_jsii_.Get(
		j,
		"connectionRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) ConnectionRulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectionRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) Decision() *string {
	var returns *string
	_jsii_.Get(
		j,
		"decision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) DecisionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"decisionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) Exclude() ZeroTrustAccessPolicyExcludeList {
	var returns ZeroTrustAccessPolicyExcludeList
	_jsii_.Get(
		j,
		"exclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) ExcludeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) Include() ZeroTrustAccessPolicyIncludeList {
	var returns ZeroTrustAccessPolicyIncludeList
	_jsii_.Get(
		j,
		"include",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) IncludeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) IsolationRequired() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isolationRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) IsolationRequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isolationRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) MfaConfig() ZeroTrustAccessPolicyMfaConfigOutputReference {
	var returns ZeroTrustAccessPolicyMfaConfigOutputReference
	_jsii_.Get(
		j,
		"mfaConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) MfaConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mfaConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) PurposeJustificationPrompt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"purposeJustificationPrompt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) PurposeJustificationPromptInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"purposeJustificationPromptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) PurposeJustificationRequired() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"purposeJustificationRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) PurposeJustificationRequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"purposeJustificationRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) Require() ZeroTrustAccessPolicyRequireList {
	var returns ZeroTrustAccessPolicyRequireList
	_jsii_.Get(
		j,
		"require",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) RequireInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) Reusable() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"reusable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) SessionDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) SessionDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/zero_trust_access_policy cloudflare_zero_trust_access_policy} Resource.
func NewZeroTrustAccessPolicy(scope constructs.Construct, id *string, config *ZeroTrustAccessPolicyConfig) ZeroTrustAccessPolicy {
	_init_.Initialize()

	if err := validateNewZeroTrustAccessPolicyParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZeroTrustAccessPolicy{}

	_jsii_.Create(
		"@cdktn/provider-cloudflare.zeroTrustAccessPolicy.ZeroTrustAccessPolicy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.24.0/docs/resources/zero_trust_access_policy cloudflare_zero_trust_access_policy} Resource.
func NewZeroTrustAccessPolicy_Override(z ZeroTrustAccessPolicy, scope constructs.Construct, id *string, config *ZeroTrustAccessPolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-cloudflare.zeroTrustAccessPolicy.ZeroTrustAccessPolicy",
		[]interface{}{scope, id, config},
		z,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicy)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicy)SetApprovalRequired(val interface{}) {
	if err := j.validateSetApprovalRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"approvalRequired",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicy)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicy)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicy)SetDecision(val *string) {
	if err := j.validateSetDecisionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"decision",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicy)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicy)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicy)SetIsolationRequired(val interface{}) {
	if err := j.validateSetIsolationRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isolationRequired",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicy)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicy)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicy)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicy)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicy)SetPurposeJustificationPrompt(val *string) {
	if err := j.validateSetPurposeJustificationPromptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"purposeJustificationPrompt",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicy)SetPurposeJustificationRequired(val interface{}) {
	if err := j.validateSetPurposeJustificationRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"purposeJustificationRequired",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicy)SetSessionDuration(val *string) {
	if err := j.validateSetSessionDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionDuration",
		val,
	)
}

// Generates CDKTN code for importing a ZeroTrustAccessPolicy resource upon running "cdktn plan <stack-name>".
func ZeroTrustAccessPolicy_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateZeroTrustAccessPolicy_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-cloudflare.zeroTrustAccessPolicy.ZeroTrustAccessPolicy",
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
func ZeroTrustAccessPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateZeroTrustAccessPolicy_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-cloudflare.zeroTrustAccessPolicy.ZeroTrustAccessPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ZeroTrustAccessPolicy_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateZeroTrustAccessPolicy_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-cloudflare.zeroTrustAccessPolicy.ZeroTrustAccessPolicy",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ZeroTrustAccessPolicy_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateZeroTrustAccessPolicy_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-cloudflare.zeroTrustAccessPolicy.ZeroTrustAccessPolicy",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ZeroTrustAccessPolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-cloudflare.zeroTrustAccessPolicy.ZeroTrustAccessPolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) AddMoveTarget(moveTarget *string) {
	if err := z.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) AddOverride(path *string, value interface{}) {
	if err := z.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (z *jsiiProxy_ZeroTrustAccessPolicy) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (z *jsiiProxy_ZeroTrustAccessPolicy) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (z *jsiiProxy_ZeroTrustAccessPolicy) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (z *jsiiProxy_ZeroTrustAccessPolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (z *jsiiProxy_ZeroTrustAccessPolicy) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (z *jsiiProxy_ZeroTrustAccessPolicy) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (z *jsiiProxy_ZeroTrustAccessPolicy) GetStringAttribute(terraformAttribute *string) *string {
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

func (z *jsiiProxy_ZeroTrustAccessPolicy) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (z *jsiiProxy_ZeroTrustAccessPolicy) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		z,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := z.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (z *jsiiProxy_ZeroTrustAccessPolicy) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (z *jsiiProxy_ZeroTrustAccessPolicy) MoveFromId(id *string) {
	if err := z.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"moveFromId",
		[]interface{}{id},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) MoveTo(moveTarget *string, index interface{}) {
	if err := z.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) MoveToId(id *string) {
	if err := z.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"moveToId",
		[]interface{}{id},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) OverrideLogicalId(newLogicalId *string) {
	if err := z.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) PutApprovalGroups(value interface{}) {
	if err := z.validatePutApprovalGroupsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putApprovalGroups",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) PutConnectionRules(value *ZeroTrustAccessPolicyConnectionRules) {
	if err := z.validatePutConnectionRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putConnectionRules",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) PutExclude(value interface{}) {
	if err := z.validatePutExcludeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putExclude",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) PutInclude(value interface{}) {
	if err := z.validatePutIncludeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putInclude",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) PutMfaConfig(value *ZeroTrustAccessPolicyMfaConfig) {
	if err := z.validatePutMfaConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putMfaConfig",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) PutRequire(value interface{}) {
	if err := z.validatePutRequireParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putRequire",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := z.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) ResetApprovalGroups() {
	_jsii_.InvokeVoid(
		z,
		"resetApprovalGroups",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) ResetApprovalRequired() {
	_jsii_.InvokeVoid(
		z,
		"resetApprovalRequired",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) ResetConnectionRules() {
	_jsii_.InvokeVoid(
		z,
		"resetConnectionRules",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) ResetExclude() {
	_jsii_.InvokeVoid(
		z,
		"resetExclude",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) ResetInclude() {
	_jsii_.InvokeVoid(
		z,
		"resetInclude",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) ResetIsolationRequired() {
	_jsii_.InvokeVoid(
		z,
		"resetIsolationRequired",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) ResetMfaConfig() {
	_jsii_.InvokeVoid(
		z,
		"resetMfaConfig",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		z,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) ResetPurposeJustificationPrompt() {
	_jsii_.InvokeVoid(
		z,
		"resetPurposeJustificationPrompt",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) ResetPurposeJustificationRequired() {
	_jsii_.InvokeVoid(
		z,
		"resetPurposeJustificationRequired",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) ResetRequire() {
	_jsii_.InvokeVoid(
		z,
		"resetRequire",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) ResetSessionDuration() {
	_jsii_.InvokeVoid(
		z,
		"resetSessionDuration",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		z,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		z,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		z,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		z,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		z,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

