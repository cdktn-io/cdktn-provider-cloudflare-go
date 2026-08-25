// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflareaisearchnamespace

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v16/datacloudflareaisearchnamespace/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataCloudflareAiSearchNamespacePublicEndpointParamsOutputReference interface {
	cdktn.ComplexObject
	AuthorizedHosts() *[]*string
	ChatCompletionsEndpoint() DataCloudflareAiSearchNamespacePublicEndpointParamsChatCompletionsEndpointOutputReference
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
	DefaultDomainEnabled() cdktn.IResolvable
	Enabled() cdktn.IResolvable
	// Experimental.
	Fqn() *string
	InstancesAllowed() *[]*string
	InternalValue() *DataCloudflareAiSearchNamespacePublicEndpointParams
	SetInternalValue(val *DataCloudflareAiSearchNamespacePublicEndpointParams)
	Mcp() DataCloudflareAiSearchNamespacePublicEndpointParamsMcpOutputReference
	RateLimit() DataCloudflareAiSearchNamespacePublicEndpointParamsRateLimitOutputReference
	SearchEndpoint() DataCloudflareAiSearchNamespacePublicEndpointParamsSearchEndpointOutputReference
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

// The jsii proxy struct for DataCloudflareAiSearchNamespacePublicEndpointParamsOutputReference
type jsiiProxy_DataCloudflareAiSearchNamespacePublicEndpointParamsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataCloudflareAiSearchNamespacePublicEndpointParamsOutputReference) AuthorizedHosts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"authorizedHosts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchNamespacePublicEndpointParamsOutputReference) ChatCompletionsEndpoint() DataCloudflareAiSearchNamespacePublicEndpointParamsChatCompletionsEndpointOutputReference {
	var returns DataCloudflareAiSearchNamespacePublicEndpointParamsChatCompletionsEndpointOutputReference
	_jsii_.Get(
		j,
		"chatCompletionsEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchNamespacePublicEndpointParamsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchNamespacePublicEndpointParamsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchNamespacePublicEndpointParamsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchNamespacePublicEndpointParamsOutputReference) CustomDomains() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchNamespacePublicEndpointParamsOutputReference) DefaultDomainEnabled() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"defaultDomainEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchNamespacePublicEndpointParamsOutputReference) Enabled() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchNamespacePublicEndpointParamsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchNamespacePublicEndpointParamsOutputReference) InstancesAllowed() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instancesAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchNamespacePublicEndpointParamsOutputReference) InternalValue() *DataCloudflareAiSearchNamespacePublicEndpointParams {
	var returns *DataCloudflareAiSearchNamespacePublicEndpointParams
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchNamespacePublicEndpointParamsOutputReference) Mcp() DataCloudflareAiSearchNamespacePublicEndpointParamsMcpOutputReference {
	var returns DataCloudflareAiSearchNamespacePublicEndpointParamsMcpOutputReference
	_jsii_.Get(
		j,
		"mcp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchNamespacePublicEndpointParamsOutputReference) RateLimit() DataCloudflareAiSearchNamespacePublicEndpointParamsRateLimitOutputReference {
	var returns DataCloudflareAiSearchNamespacePublicEndpointParamsRateLimitOutputReference
	_jsii_.Get(
		j,
		"rateLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchNamespacePublicEndpointParamsOutputReference) SearchEndpoint() DataCloudflareAiSearchNamespacePublicEndpointParamsSearchEndpointOutputReference {
	var returns DataCloudflareAiSearchNamespacePublicEndpointParamsSearchEndpointOutputReference
	_jsii_.Get(
		j,
		"searchEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchNamespacePublicEndpointParamsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchNamespacePublicEndpointParamsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataCloudflareAiSearchNamespacePublicEndpointParamsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataCloudflareAiSearchNamespacePublicEndpointParamsOutputReference {
	_init_.Initialize()

	if err := validateNewDataCloudflareAiSearchNamespacePublicEndpointParamsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareAiSearchNamespacePublicEndpointParamsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-cloudflare.dataCloudflareAiSearchNamespace.DataCloudflareAiSearchNamespacePublicEndpointParamsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataCloudflareAiSearchNamespacePublicEndpointParamsOutputReference_Override(d DataCloudflareAiSearchNamespacePublicEndpointParamsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-cloudflare.dataCloudflareAiSearchNamespace.DataCloudflareAiSearchNamespacePublicEndpointParamsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareAiSearchNamespacePublicEndpointParamsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareAiSearchNamespacePublicEndpointParamsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareAiSearchNamespacePublicEndpointParamsOutputReference)SetInternalValue(val *DataCloudflareAiSearchNamespacePublicEndpointParams) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareAiSearchNamespacePublicEndpointParamsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareAiSearchNamespacePublicEndpointParamsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataCloudflareAiSearchNamespacePublicEndpointParamsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiSearchNamespacePublicEndpointParamsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataCloudflareAiSearchNamespacePublicEndpointParamsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataCloudflareAiSearchNamespacePublicEndpointParamsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataCloudflareAiSearchNamespacePublicEndpointParamsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataCloudflareAiSearchNamespacePublicEndpointParamsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataCloudflareAiSearchNamespacePublicEndpointParamsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataCloudflareAiSearchNamespacePublicEndpointParamsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataCloudflareAiSearchNamespacePublicEndpointParamsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataCloudflareAiSearchNamespacePublicEndpointParamsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataCloudflareAiSearchNamespacePublicEndpointParamsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiSearchNamespacePublicEndpointParamsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataCloudflareAiSearchNamespacePublicEndpointParamsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataCloudflareAiSearchNamespacePublicEndpointParamsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

