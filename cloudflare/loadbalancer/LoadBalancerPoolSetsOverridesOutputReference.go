// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package loadbalancer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v17/loadbalancer/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LoadBalancerPoolSetsOverridesOutputReference interface {
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
	FallbackPool() *string
	SetFallbackPool(val *string)
	FallbackPoolInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PoolDefaultWeight() *float64
	SetPoolDefaultWeight(val *float64)
	PoolDefaultWeightInput() *float64
	Pools() *[]*string
	SetPools(val *[]*string)
	PoolsInput() *[]*string
	PoolWeights() *map[string]*float64
	SetPoolWeights(val *map[string]*float64)
	PoolWeightsInput() *map[string]*float64
	SteeringPolicy() *string
	SetSteeringPolicy(val *string)
	SteeringPolicyInput() *string
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
	ResetFallbackPool()
	ResetPoolDefaultWeight()
	ResetPools()
	ResetPoolWeights()
	ResetSteeringPolicy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LoadBalancerPoolSetsOverridesOutputReference
type jsiiProxy_LoadBalancerPoolSetsOverridesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_LoadBalancerPoolSetsOverridesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPoolSetsOverridesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPoolSetsOverridesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPoolSetsOverridesOutputReference) FallbackPool() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fallbackPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPoolSetsOverridesOutputReference) FallbackPoolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fallbackPoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPoolSetsOverridesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPoolSetsOverridesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPoolSetsOverridesOutputReference) PoolDefaultWeight() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"poolDefaultWeight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPoolSetsOverridesOutputReference) PoolDefaultWeightInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"poolDefaultWeightInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPoolSetsOverridesOutputReference) Pools() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"pools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPoolSetsOverridesOutputReference) PoolsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"poolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPoolSetsOverridesOutputReference) PoolWeights() *map[string]*float64 {
	var returns *map[string]*float64
	_jsii_.Get(
		j,
		"poolWeights",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPoolSetsOverridesOutputReference) PoolWeightsInput() *map[string]*float64 {
	var returns *map[string]*float64
	_jsii_.Get(
		j,
		"poolWeightsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPoolSetsOverridesOutputReference) SteeringPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"steeringPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPoolSetsOverridesOutputReference) SteeringPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"steeringPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPoolSetsOverridesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPoolSetsOverridesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLoadBalancerPoolSetsOverridesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) LoadBalancerPoolSetsOverridesOutputReference {
	_init_.Initialize()

	if err := validateNewLoadBalancerPoolSetsOverridesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LoadBalancerPoolSetsOverridesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-cloudflare.loadBalancer.LoadBalancerPoolSetsOverridesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLoadBalancerPoolSetsOverridesOutputReference_Override(l LoadBalancerPoolSetsOverridesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-cloudflare.loadBalancer.LoadBalancerPoolSetsOverridesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LoadBalancerPoolSetsOverridesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerPoolSetsOverridesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerPoolSetsOverridesOutputReference)SetFallbackPool(val *string) {
	if err := j.validateSetFallbackPoolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fallbackPool",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerPoolSetsOverridesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerPoolSetsOverridesOutputReference)SetPoolDefaultWeight(val *float64) {
	if err := j.validateSetPoolDefaultWeightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"poolDefaultWeight",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerPoolSetsOverridesOutputReference)SetPools(val *[]*string) {
	if err := j.validateSetPoolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pools",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerPoolSetsOverridesOutputReference)SetPoolWeights(val *map[string]*float64) {
	if err := j.validateSetPoolWeightsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"poolWeights",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerPoolSetsOverridesOutputReference)SetSteeringPolicy(val *string) {
	if err := j.validateSetSteeringPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"steeringPolicy",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerPoolSetsOverridesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerPoolSetsOverridesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LoadBalancerPoolSetsOverridesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancerPoolSetsOverridesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancerPoolSetsOverridesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancerPoolSetsOverridesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancerPoolSetsOverridesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancerPoolSetsOverridesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancerPoolSetsOverridesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancerPoolSetsOverridesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancerPoolSetsOverridesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancerPoolSetsOverridesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancerPoolSetsOverridesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancerPoolSetsOverridesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancerPoolSetsOverridesOutputReference) ResetFallbackPool() {
	_jsii_.InvokeVoid(
		l,
		"resetFallbackPool",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancerPoolSetsOverridesOutputReference) ResetPoolDefaultWeight() {
	_jsii_.InvokeVoid(
		l,
		"resetPoolDefaultWeight",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancerPoolSetsOverridesOutputReference) ResetPools() {
	_jsii_.InvokeVoid(
		l,
		"resetPools",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancerPoolSetsOverridesOutputReference) ResetPoolWeights() {
	_jsii_.InvokeVoid(
		l,
		"resetPoolWeights",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancerPoolSetsOverridesOutputReference) ResetSteeringPolicy() {
	_jsii_.InvokeVoid(
		l,
		"resetSteeringPolicy",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancerPoolSetsOverridesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := l.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancerPoolSetsOverridesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

