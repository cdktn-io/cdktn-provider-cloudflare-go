// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflareaigateways

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v15/jsii"

	"github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v15/datacloudflareaigateways/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataCloudflareAiGatewaysResultOutputReference interface {
	cdktn.ComplexObject
	Authentication() cdktn.IResolvable
	CacheInvalidateOnUpdate() cdktn.IResolvable
	CacheTtl() *float64
	CollectLogs() cdktn.IResolvable
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	CreatedAt() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Dlp() DataCloudflareAiGatewaysResultDlpOutputReference
	// Experimental.
	Fqn() *string
	Id() *string
	InternalValue() *DataCloudflareAiGatewaysResult
	SetInternalValue(val *DataCloudflareAiGatewaysResult)
	IsDefault() cdktn.IResolvable
	LogManagement() *float64
	LogManagementStrategy() *string
	Logpush() cdktn.IResolvable
	LogpushPublicKey() *string
	ModifiedAt() *string
	Otel() DataCloudflareAiGatewaysResultOtelList
	RateLimitingInterval() *float64
	RateLimitingLimit() *float64
	RateLimitingTechnique() *string
	RetryBackoff() *string
	RetryDelay() *float64
	RetryMaxAttempts() *float64
	StoreId() *string
	Stripe() DataCloudflareAiGatewaysResultStripeOutputReference
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	WorkersAiBillingMode() *string
	Zdr() cdktn.IResolvable
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataCloudflareAiGatewaysResultOutputReference
type jsiiProxy_DataCloudflareAiGatewaysResultOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultOutputReference) Authentication() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"authentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultOutputReference) CacheInvalidateOnUpdate() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"cacheInvalidateOnUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultOutputReference) CacheTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cacheTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultOutputReference) CollectLogs() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"collectLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultOutputReference) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultOutputReference) Dlp() DataCloudflareAiGatewaysResultDlpOutputReference {
	var returns DataCloudflareAiGatewaysResultDlpOutputReference
	_jsii_.Get(
		j,
		"dlp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultOutputReference) InternalValue() *DataCloudflareAiGatewaysResult {
	var returns *DataCloudflareAiGatewaysResult
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultOutputReference) IsDefault() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"isDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultOutputReference) LogManagement() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"logManagement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultOutputReference) LogManagementStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logManagementStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultOutputReference) Logpush() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"logpush",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultOutputReference) LogpushPublicKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logpushPublicKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultOutputReference) ModifiedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modifiedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultOutputReference) Otel() DataCloudflareAiGatewaysResultOtelList {
	var returns DataCloudflareAiGatewaysResultOtelList
	_jsii_.Get(
		j,
		"otel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultOutputReference) RateLimitingInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rateLimitingInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultOutputReference) RateLimitingLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rateLimitingLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultOutputReference) RateLimitingTechnique() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rateLimitingTechnique",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultOutputReference) RetryBackoff() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retryBackoff",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultOutputReference) RetryDelay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultOutputReference) RetryMaxAttempts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryMaxAttempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultOutputReference) StoreId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultOutputReference) Stripe() DataCloudflareAiGatewaysResultStripeOutputReference {
	var returns DataCloudflareAiGatewaysResultStripeOutputReference
	_jsii_.Get(
		j,
		"stripe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultOutputReference) WorkersAiBillingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workersAiBillingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultOutputReference) Zdr() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"zdr",
		&returns,
	)
	return returns
}


func NewDataCloudflareAiGatewaysResultOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataCloudflareAiGatewaysResultOutputReference {
	_init_.Initialize()

	if err := validateNewDataCloudflareAiGatewaysResultOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareAiGatewaysResultOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-cloudflare.dataCloudflareAiGateways.DataCloudflareAiGatewaysResultOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataCloudflareAiGatewaysResultOutputReference_Override(d DataCloudflareAiGatewaysResultOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-cloudflare.dataCloudflareAiGateways.DataCloudflareAiGatewaysResultOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultOutputReference)SetInternalValue(val *DataCloudflareAiGatewaysResult) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataCloudflareAiGatewaysResultOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiGatewaysResultOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataCloudflareAiGatewaysResultOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataCloudflareAiGatewaysResultOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataCloudflareAiGatewaysResultOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataCloudflareAiGatewaysResultOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataCloudflareAiGatewaysResultOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataCloudflareAiGatewaysResultOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataCloudflareAiGatewaysResultOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataCloudflareAiGatewaysResultOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataCloudflareAiGatewaysResultOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiGatewaysResultOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataCloudflareAiGatewaysResultOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiGatewaysResultOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

