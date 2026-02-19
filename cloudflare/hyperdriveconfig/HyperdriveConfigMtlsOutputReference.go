// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package hyperdriveconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v14/jsii"

	"github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v14/hyperdriveconfig/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type HyperdriveConfigMtlsOutputReference interface {
	cdktn.ComplexObject
	CaCertificateId() *string
	SetCaCertificateId(val *string)
	CaCertificateIdInput() *string
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
	MtlsCertificateId() *string
	SetMtlsCertificateId(val *string)
	MtlsCertificateIdInput() *string
	Sslmode() *string
	SetSslmode(val *string)
	SslmodeInput() *string
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
	ResetCaCertificateId()
	ResetMtlsCertificateId()
	ResetSslmode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HyperdriveConfigMtlsOutputReference
type jsiiProxy_HyperdriveConfigMtlsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_HyperdriveConfigMtlsOutputReference) CaCertificateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HyperdriveConfigMtlsOutputReference) CaCertificateIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificateIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HyperdriveConfigMtlsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HyperdriveConfigMtlsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HyperdriveConfigMtlsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HyperdriveConfigMtlsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HyperdriveConfigMtlsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HyperdriveConfigMtlsOutputReference) MtlsCertificateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mtlsCertificateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HyperdriveConfigMtlsOutputReference) MtlsCertificateIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mtlsCertificateIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HyperdriveConfigMtlsOutputReference) Sslmode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslmode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HyperdriveConfigMtlsOutputReference) SslmodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslmodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HyperdriveConfigMtlsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HyperdriveConfigMtlsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewHyperdriveConfigMtlsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) HyperdriveConfigMtlsOutputReference {
	_init_.Initialize()

	if err := validateNewHyperdriveConfigMtlsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_HyperdriveConfigMtlsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-cloudflare.hyperdriveConfig.HyperdriveConfigMtlsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHyperdriveConfigMtlsOutputReference_Override(h HyperdriveConfigMtlsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-cloudflare.hyperdriveConfig.HyperdriveConfigMtlsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HyperdriveConfigMtlsOutputReference)SetCaCertificateId(val *string) {
	if err := j.validateSetCaCertificateIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caCertificateId",
		val,
	)
}

func (j *jsiiProxy_HyperdriveConfigMtlsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HyperdriveConfigMtlsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HyperdriveConfigMtlsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HyperdriveConfigMtlsOutputReference)SetMtlsCertificateId(val *string) {
	if err := j.validateSetMtlsCertificateIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mtlsCertificateId",
		val,
	)
}

func (j *jsiiProxy_HyperdriveConfigMtlsOutputReference)SetSslmode(val *string) {
	if err := j.validateSetSslmodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslmode",
		val,
	)
}

func (j *jsiiProxy_HyperdriveConfigMtlsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HyperdriveConfigMtlsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HyperdriveConfigMtlsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HyperdriveConfigMtlsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := h.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HyperdriveConfigMtlsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := h.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HyperdriveConfigMtlsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := h.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HyperdriveConfigMtlsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := h.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HyperdriveConfigMtlsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := h.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HyperdriveConfigMtlsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := h.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HyperdriveConfigMtlsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := h.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HyperdriveConfigMtlsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := h.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HyperdriveConfigMtlsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := h.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HyperdriveConfigMtlsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HyperdriveConfigMtlsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := h.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HyperdriveConfigMtlsOutputReference) ResetCaCertificateId() {
	_jsii_.InvokeVoid(
		h,
		"resetCaCertificateId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HyperdriveConfigMtlsOutputReference) ResetMtlsCertificateId() {
	_jsii_.InvokeVoid(
		h,
		"resetMtlsCertificateId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HyperdriveConfigMtlsOutputReference) ResetSslmode() {
	_jsii_.InvokeVoid(
		h,
		"resetSslmode",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HyperdriveConfigMtlsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := h.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HyperdriveConfigMtlsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

