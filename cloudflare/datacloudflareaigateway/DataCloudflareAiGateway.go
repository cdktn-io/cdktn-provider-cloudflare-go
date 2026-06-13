// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflareaigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v15/datacloudflareaigateway/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/data-sources/ai_gateway cloudflare_ai_gateway}.
type DataCloudflareAiGateway interface {
	cdktn.TerraformDataSource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	Authentication() cdktn.IResolvable
	CacheInvalidateOnUpdate() cdktn.IResolvable
	CacheTtl() *float64
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	CollectLogs() cdktn.IResolvable
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
	Dlp() DataCloudflareAiGatewayDlpOutputReference
	Filter() DataCloudflareAiGatewayFilterOutputReference
	FilterInput() interface{}
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Guardrails() DataCloudflareAiGatewayGuardrailsOutputReference
	Id() *string
	SetId(val *string)
	IdInput() *string
	IsDefault() cdktn.IResolvable
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	LogManagement() *float64
	LogManagementStrategy() *string
	Logpush() cdktn.IResolvable
	LogpushPublicKey() *string
	ModifiedAt() *string
	// The tree node.
	Node() constructs.Node
	Otel() DataCloudflareAiGatewayOtelList
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	RateLimitingInterval() *float64
	RateLimitingLimit() *float64
	RateLimitingTechnique() *string
	// Experimental.
	RawOverrides() interface{}
	RetryBackoff() *string
	RetryDelay() *float64
	RetryMaxAttempts() *float64
	SpendLimits() DataCloudflareAiGatewaySpendLimitsOutputReference
	StoreId() *string
	Stripe() DataCloudflareAiGatewayStripeOutputReference
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	WorkersAiBillingMode() *string
	Zdr() cdktn.IResolvable
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
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutFilter(value *DataCloudflareAiGatewayFilter)
	ResetAccountId()
	ResetFilter()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
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

// The jsii proxy struct for DataCloudflareAiGateway
type jsiiProxy_DataCloudflareAiGateway struct {
	internal.Type__cdktnTerraformDataSource
}

func (j *jsiiProxy_DataCloudflareAiGateway) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGateway) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGateway) Authentication() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"authentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGateway) CacheInvalidateOnUpdate() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"cacheInvalidateOnUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGateway) CacheTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cacheTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGateway) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGateway) CollectLogs() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"collectLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGateway) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGateway) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGateway) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGateway) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGateway) Dlp() DataCloudflareAiGatewayDlpOutputReference {
	var returns DataCloudflareAiGatewayDlpOutputReference
	_jsii_.Get(
		j,
		"dlp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGateway) Filter() DataCloudflareAiGatewayFilterOutputReference {
	var returns DataCloudflareAiGatewayFilterOutputReference
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGateway) FilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGateway) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGateway) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGateway) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGateway) Guardrails() DataCloudflareAiGatewayGuardrailsOutputReference {
	var returns DataCloudflareAiGatewayGuardrailsOutputReference
	_jsii_.Get(
		j,
		"guardrails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGateway) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGateway) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGateway) IsDefault() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"isDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGateway) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGateway) LogManagement() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"logManagement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGateway) LogManagementStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logManagementStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGateway) Logpush() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"logpush",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGateway) LogpushPublicKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logpushPublicKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGateway) ModifiedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modifiedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGateway) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGateway) Otel() DataCloudflareAiGatewayOtelList {
	var returns DataCloudflareAiGatewayOtelList
	_jsii_.Get(
		j,
		"otel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGateway) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGateway) RateLimitingInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rateLimitingInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGateway) RateLimitingLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rateLimitingLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGateway) RateLimitingTechnique() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rateLimitingTechnique",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGateway) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGateway) RetryBackoff() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retryBackoff",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGateway) RetryDelay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGateway) RetryMaxAttempts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryMaxAttempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGateway) SpendLimits() DataCloudflareAiGatewaySpendLimitsOutputReference {
	var returns DataCloudflareAiGatewaySpendLimitsOutputReference
	_jsii_.Get(
		j,
		"spendLimits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGateway) StoreId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGateway) Stripe() DataCloudflareAiGatewayStripeOutputReference {
	var returns DataCloudflareAiGatewayStripeOutputReference
	_jsii_.Get(
		j,
		"stripe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGateway) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGateway) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGateway) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGateway) WorkersAiBillingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workersAiBillingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGateway) Zdr() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"zdr",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/data-sources/ai_gateway cloudflare_ai_gateway} Data Source.
func NewDataCloudflareAiGateway(scope constructs.Construct, id *string, config *DataCloudflareAiGatewayConfig) DataCloudflareAiGateway {
	_init_.Initialize()

	if err := validateNewDataCloudflareAiGatewayParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareAiGateway{}

	_jsii_.Create(
		"@cdktn/provider-cloudflare.dataCloudflareAiGateway.DataCloudflareAiGateway",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.20.0/docs/data-sources/ai_gateway cloudflare_ai_gateway} Data Source.
func NewDataCloudflareAiGateway_Override(d DataCloudflareAiGateway, scope constructs.Construct, id *string, config *DataCloudflareAiGatewayConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-cloudflare.dataCloudflareAiGateway.DataCloudflareAiGateway",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareAiGateway)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareAiGateway)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareAiGateway)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareAiGateway)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareAiGateway)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareAiGateway)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareAiGateway)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

// Generates CDKTN code for importing a DataCloudflareAiGateway resource upon running "cdktn plan <stack-name>".
func DataCloudflareAiGateway_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDataCloudflareAiGateway_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-cloudflare.dataCloudflareAiGateway.DataCloudflareAiGateway",
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
func DataCloudflareAiGateway_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataCloudflareAiGateway_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-cloudflare.dataCloudflareAiGateway.DataCloudflareAiGateway",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataCloudflareAiGateway_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataCloudflareAiGateway_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-cloudflare.dataCloudflareAiGateway.DataCloudflareAiGateway",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataCloudflareAiGateway_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataCloudflareAiGateway_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-cloudflare.dataCloudflareAiGateway.DataCloudflareAiGateway",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataCloudflareAiGateway_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-cloudflare.dataCloudflareAiGateway.DataCloudflareAiGateway",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataCloudflareAiGateway) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataCloudflareAiGateway) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiGateway) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiGateway) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiGateway) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiGateway) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiGateway) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiGateway) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiGateway) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiGateway) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiGateway) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiGateway) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataCloudflareAiGateway) PutFilter(value *DataCloudflareAiGatewayFilter) {
	if err := d.validatePutFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFilter",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataCloudflareAiGateway) ResetAccountId() {
	_jsii_.InvokeVoid(
		d,
		"resetAccountId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareAiGateway) ResetFilter() {
	_jsii_.InvokeVoid(
		d,
		"resetFilter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareAiGateway) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareAiGateway) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareAiGateway) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiGateway) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiGateway) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiGateway) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiGateway) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiGateway) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiGateway) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		d,
		"with",
		args,
		&returns,
	)

	return returns
}

