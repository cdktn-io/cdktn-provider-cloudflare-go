// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workersscript

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v14/jsii"

	"github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v14/workersscript/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WorkersScriptAssetsConfigOutputReference interface {
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
	Headers() *string
	SetHeaders(val *string)
	HeadersInput() *string
	HtmlHandling() *string
	SetHtmlHandling(val *string)
	HtmlHandlingInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	NotFoundHandling() *string
	SetNotFoundHandling(val *string)
	NotFoundHandlingInput() *string
	Redirects() *string
	SetRedirects(val *string)
	RedirectsInput() *string
	RunWorkerFirst() *map[string]interface{}
	SetRunWorkerFirst(val *map[string]interface{})
	RunWorkerFirstInput() *map[string]interface{}
	ServeDirectly() interface{}
	SetServeDirectly(val interface{})
	ServeDirectlyInput() interface{}
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
	ResetHeaders()
	ResetHtmlHandling()
	ResetNotFoundHandling()
	ResetRedirects()
	ResetRunWorkerFirst()
	ResetServeDirectly()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WorkersScriptAssetsConfigOutputReference
type jsiiProxy_WorkersScriptAssetsConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_WorkersScriptAssetsConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptAssetsConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptAssetsConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptAssetsConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptAssetsConfigOutputReference) Headers() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptAssetsConfigOutputReference) HeadersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptAssetsConfigOutputReference) HtmlHandling() *string {
	var returns *string
	_jsii_.Get(
		j,
		"htmlHandling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptAssetsConfigOutputReference) HtmlHandlingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"htmlHandlingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptAssetsConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptAssetsConfigOutputReference) NotFoundHandling() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notFoundHandling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptAssetsConfigOutputReference) NotFoundHandlingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notFoundHandlingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptAssetsConfigOutputReference) Redirects() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptAssetsConfigOutputReference) RedirectsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptAssetsConfigOutputReference) RunWorkerFirst() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"runWorkerFirst",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptAssetsConfigOutputReference) RunWorkerFirstInput() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"runWorkerFirstInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptAssetsConfigOutputReference) ServeDirectly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serveDirectly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptAssetsConfigOutputReference) ServeDirectlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serveDirectlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptAssetsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptAssetsConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWorkersScriptAssetsConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) WorkersScriptAssetsConfigOutputReference {
	_init_.Initialize()

	if err := validateNewWorkersScriptAssetsConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkersScriptAssetsConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-cloudflare.workersScript.WorkersScriptAssetsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWorkersScriptAssetsConfigOutputReference_Override(w WorkersScriptAssetsConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-cloudflare.workersScript.WorkersScriptAssetsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WorkersScriptAssetsConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WorkersScriptAssetsConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WorkersScriptAssetsConfigOutputReference)SetHeaders(val *string) {
	if err := j.validateSetHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"headers",
		val,
	)
}

func (j *jsiiProxy_WorkersScriptAssetsConfigOutputReference)SetHtmlHandling(val *string) {
	if err := j.validateSetHtmlHandlingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"htmlHandling",
		val,
	)
}

func (j *jsiiProxy_WorkersScriptAssetsConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WorkersScriptAssetsConfigOutputReference)SetNotFoundHandling(val *string) {
	if err := j.validateSetNotFoundHandlingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notFoundHandling",
		val,
	)
}

func (j *jsiiProxy_WorkersScriptAssetsConfigOutputReference)SetRedirects(val *string) {
	if err := j.validateSetRedirectsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redirects",
		val,
	)
}

func (j *jsiiProxy_WorkersScriptAssetsConfigOutputReference)SetRunWorkerFirst(val *map[string]interface{}) {
	if err := j.validateSetRunWorkerFirstParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runWorkerFirst",
		val,
	)
}

func (j *jsiiProxy_WorkersScriptAssetsConfigOutputReference)SetServeDirectly(val interface{}) {
	if err := j.validateSetServeDirectlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serveDirectly",
		val,
	)
}

func (j *jsiiProxy_WorkersScriptAssetsConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WorkersScriptAssetsConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WorkersScriptAssetsConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScriptAssetsConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScriptAssetsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScriptAssetsConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScriptAssetsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScriptAssetsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScriptAssetsConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScriptAssetsConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScriptAssetsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScriptAssetsConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScriptAssetsConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScriptAssetsConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScriptAssetsConfigOutputReference) ResetHeaders() {
	_jsii_.InvokeVoid(
		w,
		"resetHeaders",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScriptAssetsConfigOutputReference) ResetHtmlHandling() {
	_jsii_.InvokeVoid(
		w,
		"resetHtmlHandling",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScriptAssetsConfigOutputReference) ResetNotFoundHandling() {
	_jsii_.InvokeVoid(
		w,
		"resetNotFoundHandling",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScriptAssetsConfigOutputReference) ResetRedirects() {
	_jsii_.InvokeVoid(
		w,
		"resetRedirects",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScriptAssetsConfigOutputReference) ResetRunWorkerFirst() {
	_jsii_.InvokeVoid(
		w,
		"resetRunWorkerFirst",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScriptAssetsConfigOutputReference) ResetServeDirectly() {
	_jsii_.InvokeVoid(
		w,
		"resetServeDirectly",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScriptAssetsConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := w.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScriptAssetsConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

