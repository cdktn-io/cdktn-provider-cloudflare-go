// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aisearchnamespace

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v16/aisearchnamespace/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AiSearchNamespacePublicEndpointParamsOutputReference interface {
	cdktn.ComplexObject
	AuthorizedHosts() *[]*string
	SetAuthorizedHosts(val *[]*string)
	AuthorizedHostsInput() *[]*string
	ChatCompletionsEndpoint() AiSearchNamespacePublicEndpointParamsChatCompletionsEndpointOutputReference
	ChatCompletionsEndpointInput() interface{}
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
	CustomDomains() *[]*string
	SetCustomDomains(val *[]*string)
	CustomDomainsInput() *[]*string
	DefaultDomainEnabled() interface{}
	SetDefaultDomainEnabled(val interface{})
	DefaultDomainEnabledInput() interface{}
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InstancesAllowed() *[]*string
	SetInstancesAllowed(val *[]*string)
	InstancesAllowedInput() *[]*string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Mcp() AiSearchNamespacePublicEndpointParamsMcpOutputReference
	McpInput() interface{}
	RateLimit() AiSearchNamespacePublicEndpointParamsRateLimitOutputReference
	RateLimitInput() interface{}
	SearchEndpoint() AiSearchNamespacePublicEndpointParamsSearchEndpointOutputReference
	SearchEndpointInput() interface{}
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
	PutChatCompletionsEndpoint(value *AiSearchNamespacePublicEndpointParamsChatCompletionsEndpoint)
	PutMcp(value *AiSearchNamespacePublicEndpointParamsMcp)
	PutRateLimit(value *AiSearchNamespacePublicEndpointParamsRateLimit)
	PutSearchEndpoint(value *AiSearchNamespacePublicEndpointParamsSearchEndpoint)
	ResetAuthorizedHosts()
	ResetChatCompletionsEndpoint()
	ResetCustomDomains()
	ResetDefaultDomainEnabled()
	ResetEnabled()
	ResetInstancesAllowed()
	ResetMcp()
	ResetRateLimit()
	ResetSearchEndpoint()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AiSearchNamespacePublicEndpointParamsOutputReference
type jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference) AuthorizedHosts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"authorizedHosts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference) AuthorizedHostsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"authorizedHostsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference) ChatCompletionsEndpoint() AiSearchNamespacePublicEndpointParamsChatCompletionsEndpointOutputReference {
	var returns AiSearchNamespacePublicEndpointParamsChatCompletionsEndpointOutputReference
	_jsii_.Get(
		j,
		"chatCompletionsEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference) ChatCompletionsEndpointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"chatCompletionsEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference) CustomDomains() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference) CustomDomainsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference) DefaultDomainEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultDomainEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference) DefaultDomainEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultDomainEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference) InstancesAllowed() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instancesAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference) InstancesAllowedInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instancesAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference) Mcp() AiSearchNamespacePublicEndpointParamsMcpOutputReference {
	var returns AiSearchNamespacePublicEndpointParamsMcpOutputReference
	_jsii_.Get(
		j,
		"mcp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference) McpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mcpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference) RateLimit() AiSearchNamespacePublicEndpointParamsRateLimitOutputReference {
	var returns AiSearchNamespacePublicEndpointParamsRateLimitOutputReference
	_jsii_.Get(
		j,
		"rateLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference) RateLimitInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rateLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference) SearchEndpoint() AiSearchNamespacePublicEndpointParamsSearchEndpointOutputReference {
	var returns AiSearchNamespacePublicEndpointParamsSearchEndpointOutputReference
	_jsii_.Get(
		j,
		"searchEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference) SearchEndpointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"searchEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAiSearchNamespacePublicEndpointParamsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AiSearchNamespacePublicEndpointParamsOutputReference {
	_init_.Initialize()

	if err := validateNewAiSearchNamespacePublicEndpointParamsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-cloudflare.aiSearchNamespace.AiSearchNamespacePublicEndpointParamsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAiSearchNamespacePublicEndpointParamsOutputReference_Override(a AiSearchNamespacePublicEndpointParamsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-cloudflare.aiSearchNamespace.AiSearchNamespacePublicEndpointParamsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference)SetAuthorizedHosts(val *[]*string) {
	if err := j.validateSetAuthorizedHostsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizedHosts",
		val,
	)
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference)SetCustomDomains(val *[]*string) {
	if err := j.validateSetCustomDomainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customDomains",
		val,
	)
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference)SetDefaultDomainEnabled(val interface{}) {
	if err := j.validateSetDefaultDomainEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultDomainEnabled",
		val,
	)
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference)SetInstancesAllowed(val *[]*string) {
	if err := j.validateSetInstancesAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instancesAllowed",
		val,
	)
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference) PutChatCompletionsEndpoint(value *AiSearchNamespacePublicEndpointParamsChatCompletionsEndpoint) {
	if err := a.validatePutChatCompletionsEndpointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putChatCompletionsEndpoint",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference) PutMcp(value *AiSearchNamespacePublicEndpointParamsMcp) {
	if err := a.validatePutMcpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMcp",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference) PutRateLimit(value *AiSearchNamespacePublicEndpointParamsRateLimit) {
	if err := a.validatePutRateLimitParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRateLimit",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference) PutSearchEndpoint(value *AiSearchNamespacePublicEndpointParamsSearchEndpoint) {
	if err := a.validatePutSearchEndpointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSearchEndpoint",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference) ResetAuthorizedHosts() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthorizedHosts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference) ResetChatCompletionsEndpoint() {
	_jsii_.InvokeVoid(
		a,
		"resetChatCompletionsEndpoint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference) ResetCustomDomains() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomDomains",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference) ResetDefaultDomainEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultDomainEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference) ResetInstancesAllowed() {
	_jsii_.InvokeVoid(
		a,
		"resetInstancesAllowed",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference) ResetMcp() {
	_jsii_.InvokeVoid(
		a,
		"resetMcp",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference) ResetRateLimit() {
	_jsii_.InvokeVoid(
		a,
		"resetRateLimit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference) ResetSearchEndpoint() {
	_jsii_.InvokeVoid(
		a,
		"resetSearchEndpoint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AiSearchNamespacePublicEndpointParamsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

