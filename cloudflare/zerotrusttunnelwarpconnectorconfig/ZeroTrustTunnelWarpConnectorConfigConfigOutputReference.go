// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrusttunnelwarpconnectorconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v15/jsii"

	"github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v15/zerotrusttunnelwarpconnectorconfig/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ZeroTrustTunnelWarpConnectorConfigConfigOutputReference interface {
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
	FnrId() *string
	SetFnrId(val *string)
	FnrIdInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Vips() ZeroTrustTunnelWarpConnectorConfigConfigVipsList
	VipsInput() interface{}
	VipsPrevious() ZeroTrustTunnelWarpConnectorConfigConfigVipsPreviousList
	VipsPreviousInput() interface{}
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
	PutVips(value interface{})
	PutVipsPrevious(value interface{})
	ResetFnrId()
	ResetVips()
	ResetVipsPrevious()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ZeroTrustTunnelWarpConnectorConfigConfigOutputReference
type jsiiProxy_ZeroTrustTunnelWarpConnectorConfigConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ZeroTrustTunnelWarpConnectorConfigConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelWarpConnectorConfigConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelWarpConnectorConfigConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelWarpConnectorConfigConfigOutputReference) FnrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fnrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelWarpConnectorConfigConfigOutputReference) FnrIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fnrIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelWarpConnectorConfigConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelWarpConnectorConfigConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelWarpConnectorConfigConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelWarpConnectorConfigConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelWarpConnectorConfigConfigOutputReference) Vips() ZeroTrustTunnelWarpConnectorConfigConfigVipsList {
	var returns ZeroTrustTunnelWarpConnectorConfigConfigVipsList
	_jsii_.Get(
		j,
		"vips",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelWarpConnectorConfigConfigOutputReference) VipsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vipsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelWarpConnectorConfigConfigOutputReference) VipsPrevious() ZeroTrustTunnelWarpConnectorConfigConfigVipsPreviousList {
	var returns ZeroTrustTunnelWarpConnectorConfigConfigVipsPreviousList
	_jsii_.Get(
		j,
		"vipsPrevious",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelWarpConnectorConfigConfigOutputReference) VipsPreviousInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vipsPreviousInput",
		&returns,
	)
	return returns
}


func NewZeroTrustTunnelWarpConnectorConfigConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ZeroTrustTunnelWarpConnectorConfigConfigOutputReference {
	_init_.Initialize()

	if err := validateNewZeroTrustTunnelWarpConnectorConfigConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZeroTrustTunnelWarpConnectorConfigConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-cloudflare.zeroTrustTunnelWarpConnectorConfig.ZeroTrustTunnelWarpConnectorConfigConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewZeroTrustTunnelWarpConnectorConfigConfigOutputReference_Override(z ZeroTrustTunnelWarpConnectorConfigConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-cloudflare.zeroTrustTunnelWarpConnectorConfig.ZeroTrustTunnelWarpConnectorConfigConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		z,
	)
}

func (j *jsiiProxy_ZeroTrustTunnelWarpConnectorConfigConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustTunnelWarpConnectorConfigConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustTunnelWarpConnectorConfigConfigOutputReference)SetFnrId(val *string) {
	if err := j.validateSetFnrIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fnrId",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustTunnelWarpConnectorConfigConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustTunnelWarpConnectorConfigConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustTunnelWarpConnectorConfigConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (z *jsiiProxy_ZeroTrustTunnelWarpConnectorConfigConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustTunnelWarpConnectorConfigConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (z *jsiiProxy_ZeroTrustTunnelWarpConnectorConfigConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (z *jsiiProxy_ZeroTrustTunnelWarpConnectorConfigConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (z *jsiiProxy_ZeroTrustTunnelWarpConnectorConfigConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (z *jsiiProxy_ZeroTrustTunnelWarpConnectorConfigConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (z *jsiiProxy_ZeroTrustTunnelWarpConnectorConfigConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (z *jsiiProxy_ZeroTrustTunnelWarpConnectorConfigConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (z *jsiiProxy_ZeroTrustTunnelWarpConnectorConfigConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (z *jsiiProxy_ZeroTrustTunnelWarpConnectorConfigConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (z *jsiiProxy_ZeroTrustTunnelWarpConnectorConfigConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustTunnelWarpConnectorConfigConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (z *jsiiProxy_ZeroTrustTunnelWarpConnectorConfigConfigOutputReference) PutVips(value interface{}) {
	if err := z.validatePutVipsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putVips",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustTunnelWarpConnectorConfigConfigOutputReference) PutVipsPrevious(value interface{}) {
	if err := z.validatePutVipsPreviousParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putVipsPrevious",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustTunnelWarpConnectorConfigConfigOutputReference) ResetFnrId() {
	_jsii_.InvokeVoid(
		z,
		"resetFnrId",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustTunnelWarpConnectorConfigConfigOutputReference) ResetVips() {
	_jsii_.InvokeVoid(
		z,
		"resetVips",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustTunnelWarpConnectorConfigConfigOutputReference) ResetVipsPrevious() {
	_jsii_.InvokeVoid(
		z,
		"resetVipsPrevious",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustTunnelWarpConnectorConfigConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (z *jsiiProxy_ZeroTrustTunnelWarpConnectorConfigConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

