// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package flagshipflag

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v15/jsii"

	"github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v15/flagshipflag/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type FlagshipFlagRulesConditionsOutputReference interface {
	cdktn.ComplexObject
	Attribute() *string
	SetAttribute(val *string)
	AttributeInput() *string
	Clauses() FlagshipFlagRulesConditionsClausesList
	ClausesInput() interface{}
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
	LogicalOperator() *string
	SetLogicalOperator(val *string)
	LogicalOperatorInput() *string
	Operator() *string
	SetOperator(val *string)
	OperatorInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Value() *string
	SetValue(val *string)
	ValueInput() *string
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
	PutClauses(value interface{})
	ResetAttribute()
	ResetClauses()
	ResetLogicalOperator()
	ResetOperator()
	ResetValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FlagshipFlagRulesConditionsOutputReference
type jsiiProxy_FlagshipFlagRulesConditionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_FlagshipFlagRulesConditionsOutputReference) Attribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlagshipFlagRulesConditionsOutputReference) AttributeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlagshipFlagRulesConditionsOutputReference) Clauses() FlagshipFlagRulesConditionsClausesList {
	var returns FlagshipFlagRulesConditionsClausesList
	_jsii_.Get(
		j,
		"clauses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlagshipFlagRulesConditionsOutputReference) ClausesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clausesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlagshipFlagRulesConditionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlagshipFlagRulesConditionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlagshipFlagRulesConditionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlagshipFlagRulesConditionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlagshipFlagRulesConditionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlagshipFlagRulesConditionsOutputReference) LogicalOperator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalOperator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlagshipFlagRulesConditionsOutputReference) LogicalOperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalOperatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlagshipFlagRulesConditionsOutputReference) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlagshipFlagRulesConditionsOutputReference) OperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlagshipFlagRulesConditionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlagshipFlagRulesConditionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlagshipFlagRulesConditionsOutputReference) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlagshipFlagRulesConditionsOutputReference) ValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewFlagshipFlagRulesConditionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) FlagshipFlagRulesConditionsOutputReference {
	_init_.Initialize()

	if err := validateNewFlagshipFlagRulesConditionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_FlagshipFlagRulesConditionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-cloudflare.flagshipFlag.FlagshipFlagRulesConditionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewFlagshipFlagRulesConditionsOutputReference_Override(f FlagshipFlagRulesConditionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-cloudflare.flagshipFlag.FlagshipFlagRulesConditionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		f,
	)
}

func (j *jsiiProxy_FlagshipFlagRulesConditionsOutputReference)SetAttribute(val *string) {
	if err := j.validateSetAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attribute",
		val,
	)
}

func (j *jsiiProxy_FlagshipFlagRulesConditionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FlagshipFlagRulesConditionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FlagshipFlagRulesConditionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FlagshipFlagRulesConditionsOutputReference)SetLogicalOperator(val *string) {
	if err := j.validateSetLogicalOperatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logicalOperator",
		val,
	)
}

func (j *jsiiProxy_FlagshipFlagRulesConditionsOutputReference)SetOperator(val *string) {
	if err := j.validateSetOperatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operator",
		val,
	)
}

func (j *jsiiProxy_FlagshipFlagRulesConditionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FlagshipFlagRulesConditionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FlagshipFlagRulesConditionsOutputReference)SetValue(val *string) {
	if err := j.validateSetValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

func (f *jsiiProxy_FlagshipFlagRulesConditionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FlagshipFlagRulesConditionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FlagshipFlagRulesConditionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FlagshipFlagRulesConditionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FlagshipFlagRulesConditionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FlagshipFlagRulesConditionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FlagshipFlagRulesConditionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FlagshipFlagRulesConditionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FlagshipFlagRulesConditionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FlagshipFlagRulesConditionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FlagshipFlagRulesConditionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FlagshipFlagRulesConditionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FlagshipFlagRulesConditionsOutputReference) PutClauses(value interface{}) {
	if err := f.validatePutClausesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putClauses",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FlagshipFlagRulesConditionsOutputReference) ResetAttribute() {
	_jsii_.InvokeVoid(
		f,
		"resetAttribute",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FlagshipFlagRulesConditionsOutputReference) ResetClauses() {
	_jsii_.InvokeVoid(
		f,
		"resetClauses",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FlagshipFlagRulesConditionsOutputReference) ResetLogicalOperator() {
	_jsii_.InvokeVoid(
		f,
		"resetLogicalOperator",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FlagshipFlagRulesConditionsOutputReference) ResetOperator() {
	_jsii_.InvokeVoid(
		f,
		"resetOperator",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FlagshipFlagRulesConditionsOutputReference) ResetValue() {
	_jsii_.InvokeVoid(
		f,
		"resetValue",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FlagshipFlagRulesConditionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := f.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FlagshipFlagRulesConditionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

