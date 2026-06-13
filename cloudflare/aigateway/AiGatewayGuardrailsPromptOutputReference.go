// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v15/jsii"

	"github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v15/aigateway/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AiGatewayGuardrailsPromptOutputReference interface {
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
	P1() *string
	SetP1(val *string)
	P1Input() *string
	S1() *string
	SetS1(val *string)
	S10() *string
	SetS10(val *string)
	S10Input() *string
	S11() *string
	SetS11(val *string)
	S11Input() *string
	S12() *string
	SetS12(val *string)
	S12Input() *string
	S13() *string
	SetS13(val *string)
	S13Input() *string
	S1Input() *string
	S2() *string
	SetS2(val *string)
	S2Input() *string
	S3() *string
	SetS3(val *string)
	S3Input() *string
	S4() *string
	SetS4(val *string)
	S4Input() *string
	S5() *string
	SetS5(val *string)
	S5Input() *string
	S6() *string
	SetS6(val *string)
	S6Input() *string
	S7() *string
	SetS7(val *string)
	S7Input() *string
	S8() *string
	SetS8(val *string)
	S8Input() *string
	S9() *string
	SetS9(val *string)
	S9Input() *string
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
	ResetP1()
	ResetS1()
	ResetS10()
	ResetS11()
	ResetS12()
	ResetS13()
	ResetS2()
	ResetS3()
	ResetS4()
	ResetS5()
	ResetS6()
	ResetS7()
	ResetS8()
	ResetS9()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AiGatewayGuardrailsPromptOutputReference
type jsiiProxy_AiGatewayGuardrailsPromptOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) P1() *string {
	var returns *string
	_jsii_.Get(
		j,
		"p1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) P1Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"p1Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) S1() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) S10() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s10",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) S10Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s10Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) S11() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s11",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) S11Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s11Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) S12() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s12",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) S12Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s12Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) S13() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s13",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) S13Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s13Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) S1Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s1Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) S2() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) S2Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) S3() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) S3Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) S4() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) S4Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s4Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) S5() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s5",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) S5Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s5Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) S6() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) S6Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s6Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) S7() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s7",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) S7Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s7Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) S8() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s8",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) S8Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s8Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) S9() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s9",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) S9Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s9Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAiGatewayGuardrailsPromptOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AiGatewayGuardrailsPromptOutputReference {
	_init_.Initialize()

	if err := validateNewAiGatewayGuardrailsPromptOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AiGatewayGuardrailsPromptOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-cloudflare.aiGateway.AiGatewayGuardrailsPromptOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAiGatewayGuardrailsPromptOutputReference_Override(a AiGatewayGuardrailsPromptOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-cloudflare.aiGateway.AiGatewayGuardrailsPromptOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AiGatewayGuardrailsPromptOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AiGatewayGuardrailsPromptOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AiGatewayGuardrailsPromptOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AiGatewayGuardrailsPromptOutputReference)SetP1(val *string) {
	if err := j.validateSetP1Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"p1",
		val,
	)
}

func (j *jsiiProxy_AiGatewayGuardrailsPromptOutputReference)SetS1(val *string) {
	if err := j.validateSetS1Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s1",
		val,
	)
}

func (j *jsiiProxy_AiGatewayGuardrailsPromptOutputReference)SetS10(val *string) {
	if err := j.validateSetS10Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s10",
		val,
	)
}

func (j *jsiiProxy_AiGatewayGuardrailsPromptOutputReference)SetS11(val *string) {
	if err := j.validateSetS11Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s11",
		val,
	)
}

func (j *jsiiProxy_AiGatewayGuardrailsPromptOutputReference)SetS12(val *string) {
	if err := j.validateSetS12Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s12",
		val,
	)
}

func (j *jsiiProxy_AiGatewayGuardrailsPromptOutputReference)SetS13(val *string) {
	if err := j.validateSetS13Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s13",
		val,
	)
}

func (j *jsiiProxy_AiGatewayGuardrailsPromptOutputReference)SetS2(val *string) {
	if err := j.validateSetS2Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s2",
		val,
	)
}

func (j *jsiiProxy_AiGatewayGuardrailsPromptOutputReference)SetS3(val *string) {
	if err := j.validateSetS3Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3",
		val,
	)
}

func (j *jsiiProxy_AiGatewayGuardrailsPromptOutputReference)SetS4(val *string) {
	if err := j.validateSetS4Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s4",
		val,
	)
}

func (j *jsiiProxy_AiGatewayGuardrailsPromptOutputReference)SetS5(val *string) {
	if err := j.validateSetS5Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s5",
		val,
	)
}

func (j *jsiiProxy_AiGatewayGuardrailsPromptOutputReference)SetS6(val *string) {
	if err := j.validateSetS6Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s6",
		val,
	)
}

func (j *jsiiProxy_AiGatewayGuardrailsPromptOutputReference)SetS7(val *string) {
	if err := j.validateSetS7Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s7",
		val,
	)
}

func (j *jsiiProxy_AiGatewayGuardrailsPromptOutputReference)SetS8(val *string) {
	if err := j.validateSetS8Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s8",
		val,
	)
}

func (j *jsiiProxy_AiGatewayGuardrailsPromptOutputReference)SetS9(val *string) {
	if err := j.validateSetS9Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s9",
		val,
	)
}

func (j *jsiiProxy_AiGatewayGuardrailsPromptOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AiGatewayGuardrailsPromptOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) ResetP1() {
	_jsii_.InvokeVoid(
		a,
		"resetP1",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) ResetS1() {
	_jsii_.InvokeVoid(
		a,
		"resetS1",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) ResetS10() {
	_jsii_.InvokeVoid(
		a,
		"resetS10",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) ResetS11() {
	_jsii_.InvokeVoid(
		a,
		"resetS11",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) ResetS12() {
	_jsii_.InvokeVoid(
		a,
		"resetS12",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) ResetS13() {
	_jsii_.InvokeVoid(
		a,
		"resetS13",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) ResetS2() {
	_jsii_.InvokeVoid(
		a,
		"resetS2",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) ResetS3() {
	_jsii_.InvokeVoid(
		a,
		"resetS3",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) ResetS4() {
	_jsii_.InvokeVoid(
		a,
		"resetS4",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) ResetS5() {
	_jsii_.InvokeVoid(
		a,
		"resetS5",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) ResetS6() {
	_jsii_.InvokeVoid(
		a,
		"resetS6",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) ResetS7() {
	_jsii_.InvokeVoid(
		a,
		"resetS7",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) ResetS8() {
	_jsii_.InvokeVoid(
		a,
		"resetS8",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) ResetS9() {
	_jsii_.InvokeVoid(
		a,
		"resetS9",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AiGatewayGuardrailsPromptOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

