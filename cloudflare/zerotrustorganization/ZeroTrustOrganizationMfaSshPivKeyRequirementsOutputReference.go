// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustorganization

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v16/zerotrustorganization/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference interface {
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
	PinPolicy() *string
	SetPinPolicy(val *string)
	PinPolicyInput() *string
	RequireFipsDevice() interface{}
	SetRequireFipsDevice(val interface{})
	RequireFipsDeviceInput() interface{}
	SshKeySize() *[]*float64
	SetSshKeySize(val *[]*float64)
	SshKeySizeInput() *[]*float64
	SshKeyType() *[]*string
	SetSshKeyType(val *[]*string)
	SshKeyTypeInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TouchPolicy() *string
	SetTouchPolicy(val *string)
	TouchPolicyInput() *string
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
	ResetPinPolicy()
	ResetRequireFipsDevice()
	ResetSshKeySize()
	ResetSshKeyType()
	ResetTouchPolicy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference
type jsiiProxy_ZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference) PinPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pinPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference) PinPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pinPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference) RequireFipsDevice() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireFipsDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference) RequireFipsDeviceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireFipsDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference) SshKeySize() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"sshKeySize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference) SshKeySizeInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"sshKeySizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference) SshKeyType() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshKeyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference) SshKeyTypeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshKeyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference) TouchPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"touchPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference) TouchPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"touchPolicyInput",
		&returns,
	)
	return returns
}


func NewZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference {
	_init_.Initialize()

	if err := validateNewZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-cloudflare.zeroTrustOrganization.ZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference_Override(z ZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-cloudflare.zeroTrustOrganization.ZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		z,
	)
}

func (j *jsiiProxy_ZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference)SetPinPolicy(val *string) {
	if err := j.validateSetPinPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pinPolicy",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference)SetRequireFipsDevice(val interface{}) {
	if err := j.validateSetRequireFipsDeviceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireFipsDevice",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference)SetSshKeySize(val *[]*float64) {
	if err := j.validateSetSshKeySizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sshKeySize",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference)SetSshKeyType(val *[]*string) {
	if err := j.validateSetSshKeyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sshKeyType",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference)SetTouchPolicy(val *string) {
	if err := j.validateSetTouchPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"touchPolicy",
		val,
	)
}

func (z *jsiiProxy_ZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := z.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		z,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := z.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		z,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := z.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		z,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := z.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		z,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := z.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		z,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := z.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		z,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := z.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		z,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := z.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		z,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := z.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		z,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := z.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference) ResetPinPolicy() {
	_jsii_.InvokeVoid(
		z,
		"resetPinPolicy",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference) ResetRequireFipsDevice() {
	_jsii_.InvokeVoid(
		z,
		"resetRequireFipsDevice",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference) ResetSshKeySize() {
	_jsii_.InvokeVoid(
		z,
		"resetSshKeySize",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference) ResetSshKeyType() {
	_jsii_.InvokeVoid(
		z,
		"resetSshKeyType",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference) ResetTouchPolicy() {
	_jsii_.InvokeVoid(
		z,
		"resetTouchPolicy",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := z.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		z,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustOrganizationMfaSshPivKeyRequirementsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

