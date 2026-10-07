// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustcasbintegration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v16/zerotrustcasbintegration/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ZeroTrustCasbIntegrationAnthropicOutputReference interface {
	cdktn.ComplexObject
	AnthropicAdminApiKey() ZeroTrustCasbIntegrationAnthropicAnthropicAdminApiKeyOutputReference
	AnthropicAdminApiKeyInput() interface{}
	AnthropicComplianceApiKey() ZeroTrustCasbIntegrationAnthropicAnthropicComplianceApiKeyOutputReference
	AnthropicComplianceApiKeyInput() interface{}
	AnthropicWorkspaceApiKey() ZeroTrustCasbIntegrationAnthropicAnthropicWorkspaceApiKeyOutputReference
	AnthropicWorkspaceApiKeyInput() interface{}
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
	PutAnthropicAdminApiKey(value *ZeroTrustCasbIntegrationAnthropicAnthropicAdminApiKey)
	PutAnthropicComplianceApiKey(value *ZeroTrustCasbIntegrationAnthropicAnthropicComplianceApiKey)
	PutAnthropicWorkspaceApiKey(value *ZeroTrustCasbIntegrationAnthropicAnthropicWorkspaceApiKey)
	ResetAnthropicAdminApiKey()
	ResetAnthropicComplianceApiKey()
	ResetAnthropicWorkspaceApiKey()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ZeroTrustCasbIntegrationAnthropicOutputReference
type jsiiProxy_ZeroTrustCasbIntegrationAnthropicOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ZeroTrustCasbIntegrationAnthropicOutputReference) AnthropicAdminApiKey() ZeroTrustCasbIntegrationAnthropicAnthropicAdminApiKeyOutputReference {
	var returns ZeroTrustCasbIntegrationAnthropicAnthropicAdminApiKeyOutputReference
	_jsii_.Get(
		j,
		"anthropicAdminApiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustCasbIntegrationAnthropicOutputReference) AnthropicAdminApiKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"anthropicAdminApiKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustCasbIntegrationAnthropicOutputReference) AnthropicComplianceApiKey() ZeroTrustCasbIntegrationAnthropicAnthropicComplianceApiKeyOutputReference {
	var returns ZeroTrustCasbIntegrationAnthropicAnthropicComplianceApiKeyOutputReference
	_jsii_.Get(
		j,
		"anthropicComplianceApiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustCasbIntegrationAnthropicOutputReference) AnthropicComplianceApiKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"anthropicComplianceApiKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustCasbIntegrationAnthropicOutputReference) AnthropicWorkspaceApiKey() ZeroTrustCasbIntegrationAnthropicAnthropicWorkspaceApiKeyOutputReference {
	var returns ZeroTrustCasbIntegrationAnthropicAnthropicWorkspaceApiKeyOutputReference
	_jsii_.Get(
		j,
		"anthropicWorkspaceApiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustCasbIntegrationAnthropicOutputReference) AnthropicWorkspaceApiKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"anthropicWorkspaceApiKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustCasbIntegrationAnthropicOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustCasbIntegrationAnthropicOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustCasbIntegrationAnthropicOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustCasbIntegrationAnthropicOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustCasbIntegrationAnthropicOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustCasbIntegrationAnthropicOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustCasbIntegrationAnthropicOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewZeroTrustCasbIntegrationAnthropicOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ZeroTrustCasbIntegrationAnthropicOutputReference {
	_init_.Initialize()

	if err := validateNewZeroTrustCasbIntegrationAnthropicOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZeroTrustCasbIntegrationAnthropicOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-cloudflare.zeroTrustCasbIntegration.ZeroTrustCasbIntegrationAnthropicOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewZeroTrustCasbIntegrationAnthropicOutputReference_Override(z ZeroTrustCasbIntegrationAnthropicOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-cloudflare.zeroTrustCasbIntegration.ZeroTrustCasbIntegrationAnthropicOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		z,
	)
}

func (j *jsiiProxy_ZeroTrustCasbIntegrationAnthropicOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustCasbIntegrationAnthropicOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustCasbIntegrationAnthropicOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustCasbIntegrationAnthropicOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustCasbIntegrationAnthropicOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (z *jsiiProxy_ZeroTrustCasbIntegrationAnthropicOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustCasbIntegrationAnthropicOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (z *jsiiProxy_ZeroTrustCasbIntegrationAnthropicOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (z *jsiiProxy_ZeroTrustCasbIntegrationAnthropicOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (z *jsiiProxy_ZeroTrustCasbIntegrationAnthropicOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (z *jsiiProxy_ZeroTrustCasbIntegrationAnthropicOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (z *jsiiProxy_ZeroTrustCasbIntegrationAnthropicOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (z *jsiiProxy_ZeroTrustCasbIntegrationAnthropicOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (z *jsiiProxy_ZeroTrustCasbIntegrationAnthropicOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (z *jsiiProxy_ZeroTrustCasbIntegrationAnthropicOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (z *jsiiProxy_ZeroTrustCasbIntegrationAnthropicOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustCasbIntegrationAnthropicOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (z *jsiiProxy_ZeroTrustCasbIntegrationAnthropicOutputReference) PutAnthropicAdminApiKey(value *ZeroTrustCasbIntegrationAnthropicAnthropicAdminApiKey) {
	if err := z.validatePutAnthropicAdminApiKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putAnthropicAdminApiKey",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustCasbIntegrationAnthropicOutputReference) PutAnthropicComplianceApiKey(value *ZeroTrustCasbIntegrationAnthropicAnthropicComplianceApiKey) {
	if err := z.validatePutAnthropicComplianceApiKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putAnthropicComplianceApiKey",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustCasbIntegrationAnthropicOutputReference) PutAnthropicWorkspaceApiKey(value *ZeroTrustCasbIntegrationAnthropicAnthropicWorkspaceApiKey) {
	if err := z.validatePutAnthropicWorkspaceApiKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putAnthropicWorkspaceApiKey",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustCasbIntegrationAnthropicOutputReference) ResetAnthropicAdminApiKey() {
	_jsii_.InvokeVoid(
		z,
		"resetAnthropicAdminApiKey",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustCasbIntegrationAnthropicOutputReference) ResetAnthropicComplianceApiKey() {
	_jsii_.InvokeVoid(
		z,
		"resetAnthropicComplianceApiKey",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustCasbIntegrationAnthropicOutputReference) ResetAnthropicWorkspaceApiKey() {
	_jsii_.InvokeVoid(
		z,
		"resetAnthropicWorkspaceApiKey",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustCasbIntegrationAnthropicOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (z *jsiiProxy_ZeroTrustCasbIntegrationAnthropicOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

