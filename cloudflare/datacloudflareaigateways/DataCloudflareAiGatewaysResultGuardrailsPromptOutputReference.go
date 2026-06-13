// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflareaigateways

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v15/jsii"

	"github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v15/datacloudflareaigateways/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataCloudflareAiGatewaysResultGuardrailsPromptOutputReference interface {
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
	InternalValue() *DataCloudflareAiGatewaysResultGuardrailsPrompt
	SetInternalValue(val *DataCloudflareAiGatewaysResultGuardrailsPrompt)
	P1() *string
	S1() *string
	S10() *string
	S11() *string
	S12() *string
	S13() *string
	S2() *string
	S3() *string
	S4() *string
	S5() *string
	S6() *string
	S7() *string
	S8() *string
	S9() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataCloudflareAiGatewaysResultGuardrailsPromptOutputReference
type jsiiProxy_DataCloudflareAiGatewaysResultGuardrailsPromptOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultGuardrailsPromptOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultGuardrailsPromptOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultGuardrailsPromptOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultGuardrailsPromptOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultGuardrailsPromptOutputReference) InternalValue() *DataCloudflareAiGatewaysResultGuardrailsPrompt {
	var returns *DataCloudflareAiGatewaysResultGuardrailsPrompt
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultGuardrailsPromptOutputReference) P1() *string {
	var returns *string
	_jsii_.Get(
		j,
		"p1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultGuardrailsPromptOutputReference) S1() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultGuardrailsPromptOutputReference) S10() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s10",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultGuardrailsPromptOutputReference) S11() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s11",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultGuardrailsPromptOutputReference) S12() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s12",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultGuardrailsPromptOutputReference) S13() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s13",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultGuardrailsPromptOutputReference) S2() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultGuardrailsPromptOutputReference) S3() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultGuardrailsPromptOutputReference) S4() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultGuardrailsPromptOutputReference) S5() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s5",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultGuardrailsPromptOutputReference) S6() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultGuardrailsPromptOutputReference) S7() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s7",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultGuardrailsPromptOutputReference) S8() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s8",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultGuardrailsPromptOutputReference) S9() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s9",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultGuardrailsPromptOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultGuardrailsPromptOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataCloudflareAiGatewaysResultGuardrailsPromptOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataCloudflareAiGatewaysResultGuardrailsPromptOutputReference {
	_init_.Initialize()

	if err := validateNewDataCloudflareAiGatewaysResultGuardrailsPromptOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareAiGatewaysResultGuardrailsPromptOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-cloudflare.dataCloudflareAiGateways.DataCloudflareAiGatewaysResultGuardrailsPromptOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataCloudflareAiGatewaysResultGuardrailsPromptOutputReference_Override(d DataCloudflareAiGatewaysResultGuardrailsPromptOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-cloudflare.dataCloudflareAiGateways.DataCloudflareAiGatewaysResultGuardrailsPromptOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultGuardrailsPromptOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultGuardrailsPromptOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultGuardrailsPromptOutputReference)SetInternalValue(val *DataCloudflareAiGatewaysResultGuardrailsPrompt) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultGuardrailsPromptOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareAiGatewaysResultGuardrailsPromptOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataCloudflareAiGatewaysResultGuardrailsPromptOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiGatewaysResultGuardrailsPromptOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataCloudflareAiGatewaysResultGuardrailsPromptOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataCloudflareAiGatewaysResultGuardrailsPromptOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataCloudflareAiGatewaysResultGuardrailsPromptOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataCloudflareAiGatewaysResultGuardrailsPromptOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataCloudflareAiGatewaysResultGuardrailsPromptOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataCloudflareAiGatewaysResultGuardrailsPromptOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataCloudflareAiGatewaysResultGuardrailsPromptOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataCloudflareAiGatewaysResultGuardrailsPromptOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataCloudflareAiGatewaysResultGuardrailsPromptOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiGatewaysResultGuardrailsPromptOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataCloudflareAiGatewaysResultGuardrailsPromptOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataCloudflareAiGatewaysResultGuardrailsPromptOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

