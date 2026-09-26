// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package worker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v16/worker/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WorkerPreviewsBaseConfigObservabilityOutputReference interface {
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	HeadSamplingRate() *float64
	SetHeadSamplingRate(val *float64)
	HeadSamplingRateInput() *float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Issues() WorkerPreviewsBaseConfigObservabilityIssuesOutputReference
	IssuesInput() interface{}
	Logs() WorkerPreviewsBaseConfigObservabilityLogsOutputReference
	LogsInput() interface{}
	RedactQueryString() interface{}
	SetRedactQueryString(val interface{})
	RedactQueryStringInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Traces() WorkerPreviewsBaseConfigObservabilityTracesOutputReference
	TracesInput() interface{}
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
	PutIssues(value *WorkerPreviewsBaseConfigObservabilityIssues)
	PutLogs(value *WorkerPreviewsBaseConfigObservabilityLogs)
	PutTraces(value *WorkerPreviewsBaseConfigObservabilityTraces)
	ResetEnabled()
	ResetHeadSamplingRate()
	ResetIssues()
	ResetLogs()
	ResetRedactQueryString()
	ResetTraces()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WorkerPreviewsBaseConfigObservabilityOutputReference
type jsiiProxy_WorkerPreviewsBaseConfigObservabilityOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigObservabilityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigObservabilityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigObservabilityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigObservabilityOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigObservabilityOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigObservabilityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigObservabilityOutputReference) HeadSamplingRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"headSamplingRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigObservabilityOutputReference) HeadSamplingRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"headSamplingRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigObservabilityOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigObservabilityOutputReference) Issues() WorkerPreviewsBaseConfigObservabilityIssuesOutputReference {
	var returns WorkerPreviewsBaseConfigObservabilityIssuesOutputReference
	_jsii_.Get(
		j,
		"issues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigObservabilityOutputReference) IssuesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"issuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigObservabilityOutputReference) Logs() WorkerPreviewsBaseConfigObservabilityLogsOutputReference {
	var returns WorkerPreviewsBaseConfigObservabilityLogsOutputReference
	_jsii_.Get(
		j,
		"logs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigObservabilityOutputReference) LogsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigObservabilityOutputReference) RedactQueryString() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redactQueryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigObservabilityOutputReference) RedactQueryStringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redactQueryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigObservabilityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigObservabilityOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigObservabilityOutputReference) Traces() WorkerPreviewsBaseConfigObservabilityTracesOutputReference {
	var returns WorkerPreviewsBaseConfigObservabilityTracesOutputReference
	_jsii_.Get(
		j,
		"traces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigObservabilityOutputReference) TracesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tracesInput",
		&returns,
	)
	return returns
}


func NewWorkerPreviewsBaseConfigObservabilityOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) WorkerPreviewsBaseConfigObservabilityOutputReference {
	_init_.Initialize()

	if err := validateNewWorkerPreviewsBaseConfigObservabilityOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkerPreviewsBaseConfigObservabilityOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-cloudflare.worker.WorkerPreviewsBaseConfigObservabilityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWorkerPreviewsBaseConfigObservabilityOutputReference_Override(w WorkerPreviewsBaseConfigObservabilityOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-cloudflare.worker.WorkerPreviewsBaseConfigObservabilityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigObservabilityOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigObservabilityOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigObservabilityOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigObservabilityOutputReference)SetHeadSamplingRate(val *float64) {
	if err := j.validateSetHeadSamplingRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"headSamplingRate",
		val,
	)
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigObservabilityOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigObservabilityOutputReference)SetRedactQueryString(val interface{}) {
	if err := j.validateSetRedactQueryStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redactQueryString",
		val,
	)
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigObservabilityOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigObservabilityOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WorkerPreviewsBaseConfigObservabilityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkerPreviewsBaseConfigObservabilityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_WorkerPreviewsBaseConfigObservabilityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (w *jsiiProxy_WorkerPreviewsBaseConfigObservabilityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_WorkerPreviewsBaseConfigObservabilityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_WorkerPreviewsBaseConfigObservabilityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_WorkerPreviewsBaseConfigObservabilityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_WorkerPreviewsBaseConfigObservabilityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_WorkerPreviewsBaseConfigObservabilityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_WorkerPreviewsBaseConfigObservabilityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_WorkerPreviewsBaseConfigObservabilityOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkerPreviewsBaseConfigObservabilityOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (w *jsiiProxy_WorkerPreviewsBaseConfigObservabilityOutputReference) PutIssues(value *WorkerPreviewsBaseConfigObservabilityIssues) {
	if err := w.validatePutIssuesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putIssues",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkerPreviewsBaseConfigObservabilityOutputReference) PutLogs(value *WorkerPreviewsBaseConfigObservabilityLogs) {
	if err := w.validatePutLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putLogs",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkerPreviewsBaseConfigObservabilityOutputReference) PutTraces(value *WorkerPreviewsBaseConfigObservabilityTraces) {
	if err := w.validatePutTracesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putTraces",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkerPreviewsBaseConfigObservabilityOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerPreviewsBaseConfigObservabilityOutputReference) ResetHeadSamplingRate() {
	_jsii_.InvokeVoid(
		w,
		"resetHeadSamplingRate",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerPreviewsBaseConfigObservabilityOutputReference) ResetIssues() {
	_jsii_.InvokeVoid(
		w,
		"resetIssues",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerPreviewsBaseConfigObservabilityOutputReference) ResetLogs() {
	_jsii_.InvokeVoid(
		w,
		"resetLogs",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerPreviewsBaseConfigObservabilityOutputReference) ResetRedactQueryString() {
	_jsii_.InvokeVoid(
		w,
		"resetRedactQueryString",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerPreviewsBaseConfigObservabilityOutputReference) ResetTraces() {
	_jsii_.InvokeVoid(
		w,
		"resetTraces",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerPreviewsBaseConfigObservabilityOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (w *jsiiProxy_WorkerPreviewsBaseConfigObservabilityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

