// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package worker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v16/worker/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WorkerPreviewsBaseConfigOutputReference interface {
	cdktn.ComplexObject
	CacheOptions() WorkerPreviewsBaseConfigCacheOptionsOutputReference
	CacheOptionsInput() interface{}
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
	Env() WorkerPreviewsBaseConfigEnvMap
	EnvInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Limits() WorkerPreviewsBaseConfigLimitsOutputReference
	LimitsInput() interface{}
	Logpush() interface{}
	SetLogpush(val interface{})
	LogpushInput() interface{}
	Observability() WorkerPreviewsBaseConfigObservabilityOutputReference
	ObservabilityInput() interface{}
	Placement() WorkerPreviewsBaseConfigPlacementOutputReference
	PlacementInput() interface{}
	TailConsumers() WorkerPreviewsBaseConfigTailConsumersList
	TailConsumersInput() interface{}
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
	PutCacheOptions(value *WorkerPreviewsBaseConfigCacheOptions)
	PutEnv(value interface{})
	PutLimits(value *WorkerPreviewsBaseConfigLimits)
	PutObservability(value *WorkerPreviewsBaseConfigObservability)
	PutPlacement(value *WorkerPreviewsBaseConfigPlacement)
	PutTailConsumers(value interface{})
	ResetCacheOptions()
	ResetEnv()
	ResetLimits()
	ResetLogpush()
	ResetObservability()
	ResetPlacement()
	ResetTailConsumers()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WorkerPreviewsBaseConfigOutputReference
type jsiiProxy_WorkerPreviewsBaseConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigOutputReference) CacheOptions() WorkerPreviewsBaseConfigCacheOptionsOutputReference {
	var returns WorkerPreviewsBaseConfigCacheOptionsOutputReference
	_jsii_.Get(
		j,
		"cacheOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigOutputReference) CacheOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigOutputReference) Env() WorkerPreviewsBaseConfigEnvMap {
	var returns WorkerPreviewsBaseConfigEnvMap
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigOutputReference) EnvInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"envInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigOutputReference) Limits() WorkerPreviewsBaseConfigLimitsOutputReference {
	var returns WorkerPreviewsBaseConfigLimitsOutputReference
	_jsii_.Get(
		j,
		"limits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigOutputReference) LimitsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"limitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigOutputReference) Logpush() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logpush",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigOutputReference) LogpushInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logpushInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigOutputReference) Observability() WorkerPreviewsBaseConfigObservabilityOutputReference {
	var returns WorkerPreviewsBaseConfigObservabilityOutputReference
	_jsii_.Get(
		j,
		"observability",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigOutputReference) ObservabilityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"observabilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigOutputReference) Placement() WorkerPreviewsBaseConfigPlacementOutputReference {
	var returns WorkerPreviewsBaseConfigPlacementOutputReference
	_jsii_.Get(
		j,
		"placement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigOutputReference) PlacementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"placementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigOutputReference) TailConsumers() WorkerPreviewsBaseConfigTailConsumersList {
	var returns WorkerPreviewsBaseConfigTailConsumersList
	_jsii_.Get(
		j,
		"tailConsumers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigOutputReference) TailConsumersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tailConsumersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWorkerPreviewsBaseConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) WorkerPreviewsBaseConfigOutputReference {
	_init_.Initialize()

	if err := validateNewWorkerPreviewsBaseConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkerPreviewsBaseConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-cloudflare.worker.WorkerPreviewsBaseConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWorkerPreviewsBaseConfigOutputReference_Override(w WorkerPreviewsBaseConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-cloudflare.worker.WorkerPreviewsBaseConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigOutputReference)SetLogpush(val interface{}) {
	if err := j.validateSetLogpushParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logpush",
		val,
	)
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WorkerPreviewsBaseConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkerPreviewsBaseConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_WorkerPreviewsBaseConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (w *jsiiProxy_WorkerPreviewsBaseConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_WorkerPreviewsBaseConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_WorkerPreviewsBaseConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_WorkerPreviewsBaseConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_WorkerPreviewsBaseConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_WorkerPreviewsBaseConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_WorkerPreviewsBaseConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_WorkerPreviewsBaseConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkerPreviewsBaseConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (w *jsiiProxy_WorkerPreviewsBaseConfigOutputReference) PutCacheOptions(value *WorkerPreviewsBaseConfigCacheOptions) {
	if err := w.validatePutCacheOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putCacheOptions",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkerPreviewsBaseConfigOutputReference) PutEnv(value interface{}) {
	if err := w.validatePutEnvParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putEnv",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkerPreviewsBaseConfigOutputReference) PutLimits(value *WorkerPreviewsBaseConfigLimits) {
	if err := w.validatePutLimitsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putLimits",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkerPreviewsBaseConfigOutputReference) PutObservability(value *WorkerPreviewsBaseConfigObservability) {
	if err := w.validatePutObservabilityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putObservability",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkerPreviewsBaseConfigOutputReference) PutPlacement(value *WorkerPreviewsBaseConfigPlacement) {
	if err := w.validatePutPlacementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putPlacement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkerPreviewsBaseConfigOutputReference) PutTailConsumers(value interface{}) {
	if err := w.validatePutTailConsumersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putTailConsumers",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkerPreviewsBaseConfigOutputReference) ResetCacheOptions() {
	_jsii_.InvokeVoid(
		w,
		"resetCacheOptions",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerPreviewsBaseConfigOutputReference) ResetEnv() {
	_jsii_.InvokeVoid(
		w,
		"resetEnv",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerPreviewsBaseConfigOutputReference) ResetLimits() {
	_jsii_.InvokeVoid(
		w,
		"resetLimits",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerPreviewsBaseConfigOutputReference) ResetLogpush() {
	_jsii_.InvokeVoid(
		w,
		"resetLogpush",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerPreviewsBaseConfigOutputReference) ResetObservability() {
	_jsii_.InvokeVoid(
		w,
		"resetObservability",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerPreviewsBaseConfigOutputReference) ResetPlacement() {
	_jsii_.InvokeVoid(
		w,
		"resetPlacement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerPreviewsBaseConfigOutputReference) ResetTailConsumers() {
	_jsii_.InvokeVoid(
		w,
		"resetTailConsumers",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerPreviewsBaseConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (w *jsiiProxy_WorkerPreviewsBaseConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

