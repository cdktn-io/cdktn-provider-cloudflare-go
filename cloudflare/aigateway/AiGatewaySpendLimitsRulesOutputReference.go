// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v15/jsii"

	"github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v15/aigateway/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AiGatewaySpendLimitsRulesOutputReference interface {
	cdktn.ComplexObject
	AiGatewayProvider() AiGatewaySpendLimitsRulesAiGatewayProviderOutputReference
	AiGatewayProviderInput() interface{}
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Limit() *float64
	SetLimit(val *float64)
	LimitInput() *float64
	LimitType() *string
	SetLimitType(val *string)
	LimitTypeInput() *string
	Metadata() AiGatewaySpendLimitsRulesMetadataMap
	MetadataInput() interface{}
	Model() AiGatewaySpendLimitsRulesModelOutputReference
	ModelInput() interface{}
	Technique() *string
	SetTechnique(val *string)
	TechniqueInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Window() *float64
	SetWindow(val *float64)
	WindowInput() *float64
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
	PutAiGatewayProvider(value *AiGatewaySpendLimitsRulesAiGatewayProvider)
	PutMetadata(value interface{})
	PutModel(value *AiGatewaySpendLimitsRulesModel)
	ResetAiGatewayProvider()
	ResetEnabled()
	ResetId()
	ResetMetadata()
	ResetModel()
	ResetTechnique()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AiGatewaySpendLimitsRulesOutputReference
type jsiiProxy_AiGatewaySpendLimitsRulesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference) AiGatewayProvider() AiGatewaySpendLimitsRulesAiGatewayProviderOutputReference {
	var returns AiGatewaySpendLimitsRulesAiGatewayProviderOutputReference
	_jsii_.Get(
		j,
		"aiGatewayProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference) AiGatewayProviderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"aiGatewayProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference) Limit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"limit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference) LimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"limitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference) LimitType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"limitType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference) LimitTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"limitTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference) Metadata() AiGatewaySpendLimitsRulesMetadataMap {
	var returns AiGatewaySpendLimitsRulesMetadataMap
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference) MetadataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference) Model() AiGatewaySpendLimitsRulesModelOutputReference {
	var returns AiGatewaySpendLimitsRulesModelOutputReference
	_jsii_.Get(
		j,
		"model",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference) ModelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference) Technique() *string {
	var returns *string
	_jsii_.Get(
		j,
		"technique",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference) TechniqueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"techniqueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference) Window() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"window",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference) WindowInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"windowInput",
		&returns,
	)
	return returns
}


func NewAiGatewaySpendLimitsRulesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AiGatewaySpendLimitsRulesOutputReference {
	_init_.Initialize()

	if err := validateNewAiGatewaySpendLimitsRulesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AiGatewaySpendLimitsRulesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-cloudflare.aiGateway.AiGatewaySpendLimitsRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAiGatewaySpendLimitsRulesOutputReference_Override(a AiGatewaySpendLimitsRulesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-cloudflare.aiGateway.AiGatewaySpendLimitsRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference)SetLimit(val *float64) {
	if err := j.validateSetLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"limit",
		val,
	)
}

func (j *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference)SetLimitType(val *string) {
	if err := j.validateSetLimitTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"limitType",
		val,
	)
}

func (j *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference)SetTechnique(val *string) {
	if err := j.validateSetTechniqueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"technique",
		val,
	)
}

func (j *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference)SetWindow(val *float64) {
	if err := j.validateSetWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"window",
		val,
	)
}

func (a *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference) PutAiGatewayProvider(value *AiGatewaySpendLimitsRulesAiGatewayProvider) {
	if err := a.validatePutAiGatewayProviderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAiGatewayProvider",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference) PutMetadata(value interface{}) {
	if err := a.validatePutMetadataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMetadata",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference) PutModel(value *AiGatewaySpendLimitsRulesModel) {
	if err := a.validatePutModelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putModel",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference) ResetAiGatewayProvider() {
	_jsii_.InvokeVoid(
		a,
		"resetAiGatewayProvider",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference) ResetMetadata() {
	_jsii_.InvokeVoid(
		a,
		"resetMetadata",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference) ResetModel() {
	_jsii_.InvokeVoid(
		a,
		"resetModel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference) ResetTechnique() {
	_jsii_.InvokeVoid(
		a,
		"resetTechnique",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiGatewaySpendLimitsRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

