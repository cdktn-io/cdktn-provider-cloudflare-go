// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v15/aigateway/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.1/docs/resources/ai_gateway cloudflare_ai_gateway}.
type AiGateway interface {
	cdktn.TerraformResource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	Authentication() interface{}
	SetAuthentication(val interface{})
	AuthenticationInput() interface{}
	CacheInvalidateOnUpdate() interface{}
	SetCacheInvalidateOnUpdate(val interface{})
	CacheInvalidateOnUpdateInput() interface{}
	CacheTtl() *float64
	SetCacheTtl(val *float64)
	CacheTtlInput() *float64
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	CollectLogs() interface{}
	SetCollectLogs(val interface{})
	CollectLogsInput() interface{}
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
	Dlp() AiGatewayDlpOutputReference
	DlpInput() interface{}
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Guardrails() AiGatewayGuardrailsOutputReference
	GuardrailsInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	IsDefault() cdktn.IResolvable
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	LogManagement() *float64
	SetLogManagement(val *float64)
	LogManagementInput() *float64
	LogManagementStrategy() *string
	SetLogManagementStrategy(val *string)
	LogManagementStrategyInput() *string
	Logpush() interface{}
	SetLogpush(val interface{})
	LogpushInput() interface{}
	LogpushPublicKey() *string
	SetLogpushPublicKey(val *string)
	LogpushPublicKeyInput() *string
	ModifiedAt() *string
	// The tree node.
	Node() constructs.Node
	Otel() AiGatewayOtelList
	OtelInput() interface{}
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	RateLimitingInterval() *float64
	SetRateLimitingInterval(val *float64)
	RateLimitingIntervalInput() *float64
	RateLimitingLimit() *float64
	SetRateLimitingLimit(val *float64)
	RateLimitingLimitInput() *float64
	RateLimitingTechnique() *string
	SetRateLimitingTechnique(val *string)
	RateLimitingTechniqueInput() *string
	// Experimental.
	RawOverrides() interface{}
	RetryBackoff() *string
	SetRetryBackoff(val *string)
	RetryBackoffInput() *string
	RetryDelay() *float64
	SetRetryDelay(val *float64)
	RetryDelayInput() *float64
	RetryMaxAttempts() *float64
	SetRetryMaxAttempts(val *float64)
	RetryMaxAttemptsInput() *float64
	SpendLimits() AiGatewaySpendLimitsOutputReference
	SpendLimitsInput() interface{}
	StoreId() *string
	SetStoreId(val *string)
	StoreIdInput() *string
	Stripe() AiGatewayStripeOutputReference
	StripeInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	WorkersAiBillingMode() *string
	SetWorkersAiBillingMode(val *string)
	WorkersAiBillingModeInput() *string
	Zdr() interface{}
	SetZdr(val interface{})
	ZdrInput() interface{}
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
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
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
	PutDlp(value *AiGatewayDlp)
	PutGuardrails(value *AiGatewayGuardrails)
	PutOtel(value interface{})
	PutSpendLimits(value *AiGatewaySpendLimits)
	PutStripe(value *AiGatewayStripe)
	ResetAuthentication()
	ResetDlp()
	ResetGuardrails()
	ResetLogManagement()
	ResetLogManagementStrategy()
	ResetLogpush()
	ResetLogpushPublicKey()
	ResetOtel()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRateLimitingTechnique()
	ResetRetryBackoff()
	ResetRetryDelay()
	ResetRetryMaxAttempts()
	ResetSpendLimits()
	ResetStoreId()
	ResetStripe()
	ResetWorkersAiBillingMode()
	ResetZdr()
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

// The jsii proxy struct for AiGateway
type jsiiProxy_AiGateway struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AiGateway) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) Authentication() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) AuthenticationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) CacheInvalidateOnUpdate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheInvalidateOnUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) CacheInvalidateOnUpdateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheInvalidateOnUpdateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) CacheTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cacheTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) CacheTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cacheTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) CollectLogs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"collectLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) CollectLogsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"collectLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) Dlp() AiGatewayDlpOutputReference {
	var returns AiGatewayDlpOutputReference
	_jsii_.Get(
		j,
		"dlp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) DlpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dlpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) Guardrails() AiGatewayGuardrailsOutputReference {
	var returns AiGatewayGuardrailsOutputReference
	_jsii_.Get(
		j,
		"guardrails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) GuardrailsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"guardrailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) IsDefault() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"isDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) LogManagement() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"logManagement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) LogManagementInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"logManagementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) LogManagementStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logManagementStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) LogManagementStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logManagementStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) Logpush() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logpush",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) LogpushInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logpushInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) LogpushPublicKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logpushPublicKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) LogpushPublicKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logpushPublicKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) ModifiedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modifiedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) Otel() AiGatewayOtelList {
	var returns AiGatewayOtelList
	_jsii_.Get(
		j,
		"otel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) OtelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"otelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) RateLimitingInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rateLimitingInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) RateLimitingIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rateLimitingIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) RateLimitingLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rateLimitingLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) RateLimitingLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rateLimitingLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) RateLimitingTechnique() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rateLimitingTechnique",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) RateLimitingTechniqueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rateLimitingTechniqueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) RetryBackoff() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retryBackoff",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) RetryBackoffInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retryBackoffInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) RetryDelay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) RetryDelayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryDelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) RetryMaxAttempts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryMaxAttempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) RetryMaxAttemptsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryMaxAttemptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) SpendLimits() AiGatewaySpendLimitsOutputReference {
	var returns AiGatewaySpendLimitsOutputReference
	_jsii_.Get(
		j,
		"spendLimits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) SpendLimitsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"spendLimitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) StoreId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) StoreIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) Stripe() AiGatewayStripeOutputReference {
	var returns AiGatewayStripeOutputReference
	_jsii_.Get(
		j,
		"stripe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) StripeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stripeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) WorkersAiBillingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workersAiBillingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) WorkersAiBillingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workersAiBillingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) Zdr() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"zdr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGateway) ZdrInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"zdrInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.1/docs/resources/ai_gateway cloudflare_ai_gateway} Resource.
func NewAiGateway(scope constructs.Construct, id *string, config *AiGatewayConfig) AiGateway {
	_init_.Initialize()

	if err := validateNewAiGatewayParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AiGateway{}

	_jsii_.Create(
		"@cdktn/provider-cloudflare.aiGateway.AiGateway",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.21.1/docs/resources/ai_gateway cloudflare_ai_gateway} Resource.
func NewAiGateway_Override(a AiGateway, scope constructs.Construct, id *string, config *AiGatewayConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-cloudflare.aiGateway.AiGateway",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AiGateway)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_AiGateway)SetAuthentication(val interface{}) {
	if err := j.validateSetAuthenticationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authentication",
		val,
	)
}

func (j *jsiiProxy_AiGateway)SetCacheInvalidateOnUpdate(val interface{}) {
	if err := j.validateSetCacheInvalidateOnUpdateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheInvalidateOnUpdate",
		val,
	)
}

func (j *jsiiProxy_AiGateway)SetCacheTtl(val *float64) {
	if err := j.validateSetCacheTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheTtl",
		val,
	)
}

func (j *jsiiProxy_AiGateway)SetCollectLogs(val interface{}) {
	if err := j.validateSetCollectLogsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"collectLogs",
		val,
	)
}

func (j *jsiiProxy_AiGateway)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AiGateway)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AiGateway)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AiGateway)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AiGateway)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AiGateway)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AiGateway)SetLogManagement(val *float64) {
	if err := j.validateSetLogManagementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logManagement",
		val,
	)
}

func (j *jsiiProxy_AiGateway)SetLogManagementStrategy(val *string) {
	if err := j.validateSetLogManagementStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logManagementStrategy",
		val,
	)
}

func (j *jsiiProxy_AiGateway)SetLogpush(val interface{}) {
	if err := j.validateSetLogpushParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logpush",
		val,
	)
}

func (j *jsiiProxy_AiGateway)SetLogpushPublicKey(val *string) {
	if err := j.validateSetLogpushPublicKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logpushPublicKey",
		val,
	)
}

func (j *jsiiProxy_AiGateway)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AiGateway)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AiGateway)SetRateLimitingInterval(val *float64) {
	if err := j.validateSetRateLimitingIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rateLimitingInterval",
		val,
	)
}

func (j *jsiiProxy_AiGateway)SetRateLimitingLimit(val *float64) {
	if err := j.validateSetRateLimitingLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rateLimitingLimit",
		val,
	)
}

func (j *jsiiProxy_AiGateway)SetRateLimitingTechnique(val *string) {
	if err := j.validateSetRateLimitingTechniqueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rateLimitingTechnique",
		val,
	)
}

func (j *jsiiProxy_AiGateway)SetRetryBackoff(val *string) {
	if err := j.validateSetRetryBackoffParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retryBackoff",
		val,
	)
}

func (j *jsiiProxy_AiGateway)SetRetryDelay(val *float64) {
	if err := j.validateSetRetryDelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retryDelay",
		val,
	)
}

func (j *jsiiProxy_AiGateway)SetRetryMaxAttempts(val *float64) {
	if err := j.validateSetRetryMaxAttemptsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retryMaxAttempts",
		val,
	)
}

func (j *jsiiProxy_AiGateway)SetStoreId(val *string) {
	if err := j.validateSetStoreIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storeId",
		val,
	)
}

func (j *jsiiProxy_AiGateway)SetWorkersAiBillingMode(val *string) {
	if err := j.validateSetWorkersAiBillingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workersAiBillingMode",
		val,
	)
}

func (j *jsiiProxy_AiGateway)SetZdr(val interface{}) {
	if err := j.validateSetZdrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zdr",
		val,
	)
}

// Generates CDKTN code for importing a AiGateway resource upon running "cdktn plan <stack-name>".
func AiGateway_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAiGateway_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-cloudflare.aiGateway.AiGateway",
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
func AiGateway_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAiGateway_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-cloudflare.aiGateway.AiGateway",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AiGateway_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAiGateway_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-cloudflare.aiGateway.AiGateway",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AiGateway_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAiGateway_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-cloudflare.aiGateway.AiGateway",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AiGateway_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-cloudflare.aiGateway.AiGateway",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AiGateway) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AiGateway) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AiGateway) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiGateway) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiGateway) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiGateway) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiGateway) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiGateway) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiGateway) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiGateway) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiGateway) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiGateway) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiGateway) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AiGateway) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiGateway) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AiGateway) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AiGateway) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AiGateway) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AiGateway) PutDlp(value *AiGatewayDlp) {
	if err := a.validatePutDlpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDlp",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiGateway) PutGuardrails(value *AiGatewayGuardrails) {
	if err := a.validatePutGuardrailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGuardrails",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiGateway) PutOtel(value interface{}) {
	if err := a.validatePutOtelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOtel",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiGateway) PutSpendLimits(value *AiGatewaySpendLimits) {
	if err := a.validatePutSpendLimitsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSpendLimits",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiGateway) PutStripe(value *AiGatewayStripe) {
	if err := a.validatePutStripeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStripe",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiGateway) ResetAuthentication() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthentication",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGateway) ResetDlp() {
	_jsii_.InvokeVoid(
		a,
		"resetDlp",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGateway) ResetGuardrails() {
	_jsii_.InvokeVoid(
		a,
		"resetGuardrails",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGateway) ResetLogManagement() {
	_jsii_.InvokeVoid(
		a,
		"resetLogManagement",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGateway) ResetLogManagementStrategy() {
	_jsii_.InvokeVoid(
		a,
		"resetLogManagementStrategy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGateway) ResetLogpush() {
	_jsii_.InvokeVoid(
		a,
		"resetLogpush",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGateway) ResetLogpushPublicKey() {
	_jsii_.InvokeVoid(
		a,
		"resetLogpushPublicKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGateway) ResetOtel() {
	_jsii_.InvokeVoid(
		a,
		"resetOtel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGateway) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGateway) ResetRateLimitingTechnique() {
	_jsii_.InvokeVoid(
		a,
		"resetRateLimitingTechnique",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGateway) ResetRetryBackoff() {
	_jsii_.InvokeVoid(
		a,
		"resetRetryBackoff",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGateway) ResetRetryDelay() {
	_jsii_.InvokeVoid(
		a,
		"resetRetryDelay",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGateway) ResetRetryMaxAttempts() {
	_jsii_.InvokeVoid(
		a,
		"resetRetryMaxAttempts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGateway) ResetSpendLimits() {
	_jsii_.InvokeVoid(
		a,
		"resetSpendLimits",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGateway) ResetStoreId() {
	_jsii_.InvokeVoid(
		a,
		"resetStoreId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGateway) ResetStripe() {
	_jsii_.InvokeVoid(
		a,
		"resetStripe",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGateway) ResetWorkersAiBillingMode() {
	_jsii_.InvokeVoid(
		a,
		"resetWorkersAiBillingMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGateway) ResetZdr() {
	_jsii_.InvokeVoid(
		a,
		"resetZdr",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGateway) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiGateway) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiGateway) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiGateway) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiGateway) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiGateway) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiGateway) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		a,
		"with",
		args,
		&returns,
	)

	return returns
}

