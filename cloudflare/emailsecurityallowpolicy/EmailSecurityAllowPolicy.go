// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package emailsecurityallowpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v16/emailsecurityallowpolicy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.25.0/docs/resources/email_security_allow_policy cloudflare_email_security_allow_policy}.
type EmailSecurityAllowPolicy interface {
	cdktn.TerraformResource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	Comments() *string
	SetComments(val *string)
	CommentsInput() *string
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
	Id() *string
	IsAcceptableSender() interface{}
	SetIsAcceptableSender(val interface{})
	IsAcceptableSenderInput() interface{}
	IsExemptRecipient() interface{}
	SetIsExemptRecipient(val interface{})
	IsExemptRecipientInput() interface{}
	IsRecipient() interface{}
	SetIsRecipient(val interface{})
	IsRecipientInput() interface{}
	IsRegex() interface{}
	SetIsRegex(val interface{})
	IsRegexInput() interface{}
	IsSender() interface{}
	SetIsSender(val interface{})
	IsSenderInput() interface{}
	IsSpoof() interface{}
	SetIsSpoof(val interface{})
	IsSpoofInput() interface{}
	IsTrustedSender() interface{}
	SetIsTrustedSender(val interface{})
	IsTrustedSenderInput() interface{}
	LastModified() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	ModifiedAt() *string
	// The tree node.
	Node() constructs.Node
	Pattern() *string
	SetPattern(val *string)
	PatternInput() *string
	PatternType() *string
	SetPatternType(val *string)
	PatternTypeInput() *string
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	VerifySender() interface{}
	SetVerifySender(val interface{})
	VerifySenderInput() interface{}
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
	ResetComments()
	ResetIsRecipient()
	ResetIsSender()
	ResetIsSpoof()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for EmailSecurityAllowPolicy
type jsiiProxy_EmailSecurityAllowPolicy struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_EmailSecurityAllowPolicy) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityAllowPolicy) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityAllowPolicy) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityAllowPolicy) Comments() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityAllowPolicy) CommentsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityAllowPolicy) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityAllowPolicy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityAllowPolicy) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityAllowPolicy) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityAllowPolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityAllowPolicy) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityAllowPolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityAllowPolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityAllowPolicy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityAllowPolicy) IsAcceptableSender() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isAcceptableSender",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityAllowPolicy) IsAcceptableSenderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isAcceptableSenderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityAllowPolicy) IsExemptRecipient() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isExemptRecipient",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityAllowPolicy) IsExemptRecipientInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isExemptRecipientInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityAllowPolicy) IsRecipient() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isRecipient",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityAllowPolicy) IsRecipientInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isRecipientInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityAllowPolicy) IsRegex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityAllowPolicy) IsRegexInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityAllowPolicy) IsSender() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSender",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityAllowPolicy) IsSenderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSenderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityAllowPolicy) IsSpoof() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSpoof",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityAllowPolicy) IsSpoofInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSpoofInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityAllowPolicy) IsTrustedSender() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isTrustedSender",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityAllowPolicy) IsTrustedSenderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isTrustedSenderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityAllowPolicy) LastModified() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModified",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityAllowPolicy) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityAllowPolicy) ModifiedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modifiedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityAllowPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityAllowPolicy) Pattern() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityAllowPolicy) PatternInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityAllowPolicy) PatternType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patternType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityAllowPolicy) PatternTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patternTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityAllowPolicy) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityAllowPolicy) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityAllowPolicy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityAllowPolicy) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityAllowPolicy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityAllowPolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityAllowPolicy) VerifySender() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verifySender",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityAllowPolicy) VerifySenderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verifySenderInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.25.0/docs/resources/email_security_allow_policy cloudflare_email_security_allow_policy} Resource.
func NewEmailSecurityAllowPolicy(scope constructs.Construct, id *string, config *EmailSecurityAllowPolicyConfig) EmailSecurityAllowPolicy {
	_init_.Initialize()

	if err := validateNewEmailSecurityAllowPolicyParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_EmailSecurityAllowPolicy{}

	_jsii_.Create(
		"@cdktn/provider-cloudflare.emailSecurityAllowPolicy.EmailSecurityAllowPolicy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.25.0/docs/resources/email_security_allow_policy cloudflare_email_security_allow_policy} Resource.
func NewEmailSecurityAllowPolicy_Override(e EmailSecurityAllowPolicy, scope constructs.Construct, id *string, config *EmailSecurityAllowPolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-cloudflare.emailSecurityAllowPolicy.EmailSecurityAllowPolicy",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_EmailSecurityAllowPolicy)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityAllowPolicy)SetComments(val *string) {
	if err := j.validateSetCommentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comments",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityAllowPolicy)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityAllowPolicy)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityAllowPolicy)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityAllowPolicy)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityAllowPolicy)SetIsAcceptableSender(val interface{}) {
	if err := j.validateSetIsAcceptableSenderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isAcceptableSender",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityAllowPolicy)SetIsExemptRecipient(val interface{}) {
	if err := j.validateSetIsExemptRecipientParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isExemptRecipient",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityAllowPolicy)SetIsRecipient(val interface{}) {
	if err := j.validateSetIsRecipientParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isRecipient",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityAllowPolicy)SetIsRegex(val interface{}) {
	if err := j.validateSetIsRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isRegex",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityAllowPolicy)SetIsSender(val interface{}) {
	if err := j.validateSetIsSenderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isSender",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityAllowPolicy)SetIsSpoof(val interface{}) {
	if err := j.validateSetIsSpoofParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isSpoof",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityAllowPolicy)SetIsTrustedSender(val interface{}) {
	if err := j.validateSetIsTrustedSenderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isTrustedSender",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityAllowPolicy)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityAllowPolicy)SetPattern(val *string) {
	if err := j.validateSetPatternParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pattern",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityAllowPolicy)SetPatternType(val *string) {
	if err := j.validateSetPatternTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"patternType",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityAllowPolicy)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityAllowPolicy)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityAllowPolicy)SetVerifySender(val interface{}) {
	if err := j.validateSetVerifySenderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verifySender",
		val,
	)
}

// Generates CDKTN code for importing a EmailSecurityAllowPolicy resource upon running "cdktn plan <stack-name>".
func EmailSecurityAllowPolicy_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateEmailSecurityAllowPolicy_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-cloudflare.emailSecurityAllowPolicy.EmailSecurityAllowPolicy",
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
func EmailSecurityAllowPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEmailSecurityAllowPolicy_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-cloudflare.emailSecurityAllowPolicy.EmailSecurityAllowPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EmailSecurityAllowPolicy_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEmailSecurityAllowPolicy_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-cloudflare.emailSecurityAllowPolicy.EmailSecurityAllowPolicy",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EmailSecurityAllowPolicy_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEmailSecurityAllowPolicy_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-cloudflare.emailSecurityAllowPolicy.EmailSecurityAllowPolicy",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func EmailSecurityAllowPolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-cloudflare.emailSecurityAllowPolicy.EmailSecurityAllowPolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_EmailSecurityAllowPolicy) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_EmailSecurityAllowPolicy) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_EmailSecurityAllowPolicy) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EmailSecurityAllowPolicy) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_EmailSecurityAllowPolicy) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EmailSecurityAllowPolicy) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EmailSecurityAllowPolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EmailSecurityAllowPolicy) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EmailSecurityAllowPolicy) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EmailSecurityAllowPolicy) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EmailSecurityAllowPolicy) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EmailSecurityAllowPolicy) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailSecurityAllowPolicy) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_EmailSecurityAllowPolicy) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_EmailSecurityAllowPolicy) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (e *jsiiProxy_EmailSecurityAllowPolicy) MoveFromId(id *string) {
	if err := e.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveFromId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_EmailSecurityAllowPolicy) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_EmailSecurityAllowPolicy) MoveToId(id *string) {
	if err := e.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveToId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_EmailSecurityAllowPolicy) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_EmailSecurityAllowPolicy) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := e.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (e *jsiiProxy_EmailSecurityAllowPolicy) ResetComments() {
	_jsii_.InvokeVoid(
		e,
		"resetComments",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmailSecurityAllowPolicy) ResetIsRecipient() {
	_jsii_.InvokeVoid(
		e,
		"resetIsRecipient",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmailSecurityAllowPolicy) ResetIsSender() {
	_jsii_.InvokeVoid(
		e,
		"resetIsSender",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmailSecurityAllowPolicy) ResetIsSpoof() {
	_jsii_.InvokeVoid(
		e,
		"resetIsSpoof",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmailSecurityAllowPolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmailSecurityAllowPolicy) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailSecurityAllowPolicy) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailSecurityAllowPolicy) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailSecurityAllowPolicy) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailSecurityAllowPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailSecurityAllowPolicy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailSecurityAllowPolicy) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

