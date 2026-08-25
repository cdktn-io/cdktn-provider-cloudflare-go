// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aisearchnamespace

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v16/aisearchnamespace/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AiSearchNamespacePublicEndpointParamsRateLimitOutputReference interface {
	cdktn.ComplexObject
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PeriodMs() *float64
	SetPeriodMs(val *float64)
	PeriodMsInput() *float64
	Requests() *float64
	SetRequests(val *float64)
	RequestsInput() *float64
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
	ResetPeriodMs()
	ResetRequests()
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

// The jsii proxy struct for AiSearchNamespacePublicEndpointParamsRateLimitOutputReference
type jsiiProxy_AiSearchNamespacePublicEndpointParamsRateLimitOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsRateLimitOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsRateLimitOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsRateLimitOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsRateLimitOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsRateLimitOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsRateLimitOutputReference) PeriodMs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"periodMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsRateLimitOutputReference) PeriodMsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"periodMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsRateLimitOutputReference) Requests() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsRateLimitOutputReference) RequestsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsRateLimitOutputReference) Technique() *string {
	var returns *string
	_jsii_.Get(
		j,
		"technique",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsRateLimitOutputReference) TechniqueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"techniqueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsRateLimitOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsRateLimitOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAiSearchNamespacePublicEndpointParamsRateLimitOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AiSearchNamespacePublicEndpointParamsRateLimitOutputReference {
	_init_.Initialize()

	if err := validateNewAiSearchNamespacePublicEndpointParamsRateLimitOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AiSearchNamespacePublicEndpointParamsRateLimitOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-cloudflare.aiSearchNamespace.AiSearchNamespacePublicEndpointParamsRateLimitOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAiSearchNamespacePublicEndpointParamsRateLimitOutputReference_Override(a AiSearchNamespacePublicEndpointParamsRateLimitOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-cloudflare.aiSearchNamespace.AiSearchNamespacePublicEndpointParamsRateLimitOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsRateLimitOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsRateLimitOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsRateLimitOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsRateLimitOutputReference)SetPeriodMs(val *float64) {
	if err := j.validateSetPeriodMsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"periodMs",
		val,
	)
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsRateLimitOutputReference)SetRequests(val *float64) {
	if err := j.validateSetRequestsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requests",
		val,
	)
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsRateLimitOutputReference)SetTechnique(val *string) {
	if err := j.validateSetTechniqueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"technique",
		val,
	)
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsRateLimitOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsRateLimitOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AiSearchNamespacePublicEndpointParamsRateLimitOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchNamespacePublicEndpointParamsRateLimitOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AiSearchNamespacePublicEndpointParamsRateLimitOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AiSearchNamespacePublicEndpointParamsRateLimitOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AiSearchNamespacePublicEndpointParamsRateLimitOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AiSearchNamespacePublicEndpointParamsRateLimitOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AiSearchNamespacePublicEndpointParamsRateLimitOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AiSearchNamespacePublicEndpointParamsRateLimitOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AiSearchNamespacePublicEndpointParamsRateLimitOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AiSearchNamespacePublicEndpointParamsRateLimitOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AiSearchNamespacePublicEndpointParamsRateLimitOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchNamespacePublicEndpointParamsRateLimitOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AiSearchNamespacePublicEndpointParamsRateLimitOutputReference) ResetPeriodMs() {
	_jsii_.InvokeVoid(
		a,
		"resetPeriodMs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchNamespacePublicEndpointParamsRateLimitOutputReference) ResetRequests() {
	_jsii_.InvokeVoid(
		a,
		"resetRequests",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchNamespacePublicEndpointParamsRateLimitOutputReference) ResetTechnique() {
	_jsii_.InvokeVoid(
		a,
		"resetTechnique",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchNamespacePublicEndpointParamsRateLimitOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AiSearchNamespacePublicEndpointParamsRateLimitOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

