// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aisearchinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v16/aisearchinstance/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference interface {
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
	Depth() *float64
	SetDepth(val *float64)
	DepthInput() *float64
	// Experimental.
	Fqn() *string
	IncludeExternalLinks() interface{}
	SetIncludeExternalLinks(val interface{})
	IncludeExternalLinksInput() interface{}
	IncludeSubdomains() interface{}
	SetIncludeSubdomains(val interface{})
	IncludeSubdomainsInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Limit() *float64
	SetLimit(val *float64)
	LimitInput() *float64
	MaxAge() *float64
	SetMaxAge(val *float64)
	MaxAgeInput() *float64
	Source() *string
	SetSource(val *string)
	SourceInput() *string
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
	ResetDepth()
	ResetIncludeExternalLinks()
	ResetIncludeSubdomains()
	ResetLimit()
	ResetMaxAge()
	ResetSource()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference
type jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference) Depth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"depth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference) DepthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"depthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference) IncludeExternalLinks() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeExternalLinks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference) IncludeExternalLinksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeExternalLinksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference) IncludeSubdomains() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeSubdomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference) IncludeSubdomainsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeSubdomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference) Limit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"limit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference) LimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"limitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference) MaxAge() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference) MaxAgeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAgeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference) SourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewAiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-cloudflare.aiSearchInstance.AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference_Override(a AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-cloudflare.aiSearchInstance.AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference)SetDepth(val *float64) {
	if err := j.validateSetDepthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"depth",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference)SetIncludeExternalLinks(val interface{}) {
	if err := j.validateSetIncludeExternalLinksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeExternalLinks",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference)SetIncludeSubdomains(val interface{}) {
	if err := j.validateSetIncludeSubdomainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeSubdomains",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference)SetLimit(val *float64) {
	if err := j.validateSetLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"limit",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference)SetMaxAge(val *float64) {
	if err := j.validateSetMaxAgeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxAge",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference)SetSource(val *string) {
	if err := j.validateSetSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"source",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference) ResetDepth() {
	_jsii_.InvokeVoid(
		a,
		"resetDepth",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference) ResetIncludeExternalLinks() {
	_jsii_.InvokeVoid(
		a,
		"resetIncludeExternalLinks",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference) ResetIncludeSubdomains() {
	_jsii_.InvokeVoid(
		a,
		"resetIncludeSubdomains",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference) ResetLimit() {
	_jsii_.InvokeVoid(
		a,
		"resetLimit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference) ResetMaxAge() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxAge",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference) ResetSource() {
	_jsii_.InvokeVoid(
		a,
		"resetSource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerDiscoverOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

