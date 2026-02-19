// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package tokenvalidationconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v14/jsii"

	"github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v14/tokenvalidationconfig/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type TokenValidationConfigCredentialsKeysOutputReference interface {
	cdktn.ComplexObject
	Alg() *string
	SetAlg(val *string)
	AlgInput() *string
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
	Crv() *string
	SetCrv(val *string)
	CrvInput() *string
	E() *string
	SetE(val *string)
	EInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Kid() *string
	SetKid(val *string)
	KidInput() *string
	Kty() *string
	SetKty(val *string)
	KtyInput() *string
	N() *string
	SetN(val *string)
	NInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	X() *string
	SetX(val *string)
	XInput() *string
	Y() *string
	SetY(val *string)
	YInput() *string
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
	ResetCrv()
	ResetE()
	ResetN()
	ResetX()
	ResetY()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TokenValidationConfigCredentialsKeysOutputReference
type jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference) Alg() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alg",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference) AlgInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"algInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference) Crv() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crv",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference) CrvInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crvInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference) E() *string {
	var returns *string
	_jsii_.Get(
		j,
		"e",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference) EInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference) Kid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference) KidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference) Kty() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference) KtyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ktyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference) N() *string {
	var returns *string
	_jsii_.Get(
		j,
		"n",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference) NInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference) X() *string {
	var returns *string
	_jsii_.Get(
		j,
		"x",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference) XInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"xInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference) Y() *string {
	var returns *string
	_jsii_.Get(
		j,
		"y",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference) YInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"yInput",
		&returns,
	)
	return returns
}


func NewTokenValidationConfigCredentialsKeysOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TokenValidationConfigCredentialsKeysOutputReference {
	_init_.Initialize()

	if err := validateNewTokenValidationConfigCredentialsKeysOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-cloudflare.tokenValidationConfig.TokenValidationConfigCredentialsKeysOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewTokenValidationConfigCredentialsKeysOutputReference_Override(t TokenValidationConfigCredentialsKeysOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-cloudflare.tokenValidationConfig.TokenValidationConfigCredentialsKeysOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference)SetAlg(val *string) {
	if err := j.validateSetAlgParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alg",
		val,
	)
}

func (j *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference)SetCrv(val *string) {
	if err := j.validateSetCrvParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crv",
		val,
	)
}

func (j *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference)SetE(val *string) {
	if err := j.validateSetEParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"e",
		val,
	)
}

func (j *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference)SetKid(val *string) {
	if err := j.validateSetKidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kid",
		val,
	)
}

func (j *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference)SetKty(val *string) {
	if err := j.validateSetKtyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kty",
		val,
	)
}

func (j *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference)SetN(val *string) {
	if err := j.validateSetNParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"n",
		val,
	)
}

func (j *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference)SetX(val *string) {
	if err := j.validateSetXParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"x",
		val,
	)
}

func (j *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference)SetY(val *string) {
	if err := j.validateSetYParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"y",
		val,
	)
}

func (t *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference) ResetCrv() {
	_jsii_.InvokeVoid(
		t,
		"resetCrv",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference) ResetE() {
	_jsii_.InvokeVoid(
		t,
		"resetE",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference) ResetN() {
	_jsii_.InvokeVoid(
		t,
		"resetN",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference) ResetX() {
	_jsii_.InvokeVoid(
		t,
		"resetX",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference) ResetY() {
	_jsii_.InvokeVoid(
		t,
		"resetY",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := t.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TokenValidationConfigCredentialsKeysOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

