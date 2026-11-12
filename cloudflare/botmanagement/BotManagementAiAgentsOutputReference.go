// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package botmanagement

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v17/botmanagement/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BotManagementAiAgentsOutputReference interface {
	cdktn.ComplexObject
	AnthropicAgent() *string
	SetAnthropicAgent(val *string)
	AnthropicAgentInput() *string
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
	GoogleAgent() *string
	SetGoogleAgent(val *string)
	GoogleAgentInput() *string
	Instinct() *string
	SetInstinct(val *string)
	InstinctInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Muse() *string
	SetMuse(val *string)
	MuseInput() *string
	OpenaiAgent() *string
	SetOpenaiAgent(val *string)
	OpenaiAgentInput() *string
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
	ResetAnthropicAgent()
	ResetGoogleAgent()
	ResetInstinct()
	ResetMuse()
	ResetOpenaiAgent()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BotManagementAiAgentsOutputReference
type jsiiProxy_BotManagementAiAgentsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BotManagementAiAgentsOutputReference) AnthropicAgent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"anthropicAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotManagementAiAgentsOutputReference) AnthropicAgentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"anthropicAgentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotManagementAiAgentsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotManagementAiAgentsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotManagementAiAgentsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotManagementAiAgentsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotManagementAiAgentsOutputReference) GoogleAgent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"googleAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotManagementAiAgentsOutputReference) GoogleAgentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"googleAgentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotManagementAiAgentsOutputReference) Instinct() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instinct",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotManagementAiAgentsOutputReference) InstinctInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instinctInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotManagementAiAgentsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotManagementAiAgentsOutputReference) Muse() *string {
	var returns *string
	_jsii_.Get(
		j,
		"muse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotManagementAiAgentsOutputReference) MuseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"museInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotManagementAiAgentsOutputReference) OpenaiAgent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openaiAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotManagementAiAgentsOutputReference) OpenaiAgentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openaiAgentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotManagementAiAgentsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotManagementAiAgentsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBotManagementAiAgentsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) BotManagementAiAgentsOutputReference {
	_init_.Initialize()

	if err := validateNewBotManagementAiAgentsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BotManagementAiAgentsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-cloudflare.botManagement.BotManagementAiAgentsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBotManagementAiAgentsOutputReference_Override(b BotManagementAiAgentsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-cloudflare.botManagement.BotManagementAiAgentsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BotManagementAiAgentsOutputReference)SetAnthropicAgent(val *string) {
	if err := j.validateSetAnthropicAgentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"anthropicAgent",
		val,
	)
}

func (j *jsiiProxy_BotManagementAiAgentsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BotManagementAiAgentsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BotManagementAiAgentsOutputReference)SetGoogleAgent(val *string) {
	if err := j.validateSetGoogleAgentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"googleAgent",
		val,
	)
}

func (j *jsiiProxy_BotManagementAiAgentsOutputReference)SetInstinct(val *string) {
	if err := j.validateSetInstinctParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instinct",
		val,
	)
}

func (j *jsiiProxy_BotManagementAiAgentsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BotManagementAiAgentsOutputReference)SetMuse(val *string) {
	if err := j.validateSetMuseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"muse",
		val,
	)
}

func (j *jsiiProxy_BotManagementAiAgentsOutputReference)SetOpenaiAgent(val *string) {
	if err := j.validateSetOpenaiAgentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"openaiAgent",
		val,
	)
}

func (j *jsiiProxy_BotManagementAiAgentsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BotManagementAiAgentsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BotManagementAiAgentsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotManagementAiAgentsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotManagementAiAgentsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotManagementAiAgentsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotManagementAiAgentsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotManagementAiAgentsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotManagementAiAgentsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotManagementAiAgentsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotManagementAiAgentsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotManagementAiAgentsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotManagementAiAgentsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotManagementAiAgentsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotManagementAiAgentsOutputReference) ResetAnthropicAgent() {
	_jsii_.InvokeVoid(
		b,
		"resetAnthropicAgent",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotManagementAiAgentsOutputReference) ResetGoogleAgent() {
	_jsii_.InvokeVoid(
		b,
		"resetGoogleAgent",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotManagementAiAgentsOutputReference) ResetInstinct() {
	_jsii_.InvokeVoid(
		b,
		"resetInstinct",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotManagementAiAgentsOutputReference) ResetMuse() {
	_jsii_.InvokeVoid(
		b,
		"resetMuse",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotManagementAiAgentsOutputReference) ResetOpenaiAgent() {
	_jsii_.InvokeVoid(
		b,
		"resetOpenaiAgent",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotManagementAiAgentsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := b.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotManagementAiAgentsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

