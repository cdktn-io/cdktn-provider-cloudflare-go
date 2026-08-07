// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v16/zerotrustaccessapplication/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistrationOutputReference interface {
	cdktn.ComplexObject
	AllowAnyOnLocalhost() interface{}
	SetAllowAnyOnLocalhost(val interface{})
	AllowAnyOnLocalhostInput() interface{}
	AllowAnyOnLoopback() interface{}
	SetAllowAnyOnLoopback(val interface{})
	AllowAnyOnLoopbackInput() interface{}
	AllowedUris() *[]*string
	SetAllowedUris(val *[]*string)
	AllowedUrisInput() *[]*string
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
	ResetAllowAnyOnLocalhost()
	ResetAllowAnyOnLoopback()
	ResetAllowedUris()
	ResetEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistrationOutputReference
type jsiiProxy_ZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistrationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistrationOutputReference) AllowAnyOnLocalhost() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowAnyOnLocalhost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistrationOutputReference) AllowAnyOnLocalhostInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowAnyOnLocalhostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistrationOutputReference) AllowAnyOnLoopback() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowAnyOnLoopback",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistrationOutputReference) AllowAnyOnLoopbackInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowAnyOnLoopbackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistrationOutputReference) AllowedUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistrationOutputReference) AllowedUrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedUrisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistrationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistrationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistrationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistrationOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistrationOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistrationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistrationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistrationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistrationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistrationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistrationOutputReference {
	_init_.Initialize()

	if err := validateNewZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistrationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistrationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistrationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistrationOutputReference_Override(z ZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistrationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistrationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		z,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistrationOutputReference)SetAllowAnyOnLocalhost(val interface{}) {
	if err := j.validateSetAllowAnyOnLocalhostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowAnyOnLocalhost",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistrationOutputReference)SetAllowAnyOnLoopback(val interface{}) {
	if err := j.validateSetAllowAnyOnLoopbackParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowAnyOnLoopback",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistrationOutputReference)SetAllowedUris(val *[]*string) {
	if err := j.validateSetAllowedUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedUris",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistrationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistrationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistrationOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistrationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistrationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistrationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistrationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistrationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistrationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistrationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistrationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistrationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistrationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistrationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistrationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistrationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistrationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistrationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistrationOutputReference) ResetAllowAnyOnLocalhost() {
	_jsii_.InvokeVoid(
		z,
		"resetAllowAnyOnLocalhost",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistrationOutputReference) ResetAllowAnyOnLoopback() {
	_jsii_.InvokeVoid(
		z,
		"resetAllowAnyOnLoopback",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistrationOutputReference) ResetAllowedUris() {
	_jsii_.InvokeVoid(
		z,
		"resetAllowedUris",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistrationOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		z,
		"resetEnabled",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistrationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistrationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

