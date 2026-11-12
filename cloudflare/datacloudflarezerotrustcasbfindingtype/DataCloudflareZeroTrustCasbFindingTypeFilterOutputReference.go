// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezerotrustcasbfindingtype

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v17/datacloudflarezerotrustcasbfindingtype/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference interface {
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
	Direction() *string
	SetDirection(val *string)
	DirectionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaxUpdatedAt() *string
	SetMaxUpdatedAt(val *string)
	MaxUpdatedAtInput() *string
	MinUpdatedAt() *string
	SetMinUpdatedAt(val *string)
	MinUpdatedAtInput() *string
	Order() *string
	SetOrder(val *string)
	OrderInput() *string
	Origins() *[]*string
	SetOrigins(val *[]*string)
	OriginsInput() *[]*string
	Quarantined() interface{}
	SetQuarantined(val interface{})
	QuarantinedInput() interface{}
	Search() *string
	SetSearch(val *string)
	SearchInput() *string
	Severities() *[]*string
	SetSeverities(val *[]*string)
	SeveritiesInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Types() *[]*string
	SetTypes(val *[]*string)
	TypesInput() *[]*string
	Vendors() *[]*string
	SetVendors(val *[]*string)
	VendorsInput() *[]*string
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
	ResetDirection()
	ResetMaxUpdatedAt()
	ResetMinUpdatedAt()
	ResetOrder()
	ResetOrigins()
	ResetQuarantined()
	ResetSearch()
	ResetSeverities()
	ResetTypes()
	ResetVendors()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference
type jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference) Direction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"direction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference) DirectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference) MaxUpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxUpdatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference) MaxUpdatedAtInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxUpdatedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference) MinUpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minUpdatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference) MinUpdatedAtInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minUpdatedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference) Order() *string {
	var returns *string
	_jsii_.Get(
		j,
		"order",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference) OrderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference) Origins() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"origins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference) OriginsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"originsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference) Quarantined() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"quarantined",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference) QuarantinedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"quarantinedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference) Search() *string {
	var returns *string
	_jsii_.Get(
		j,
		"search",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference) SearchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"searchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference) Severities() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"severities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference) SeveritiesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"severitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference) Types() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"types",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference) TypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"typesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference) Vendors() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vendors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference) VendorsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vendorsInput",
		&returns,
	)
	return returns
}


func NewDataCloudflareZeroTrustCasbFindingTypeFilterOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference {
	_init_.Initialize()

	if err := validateNewDataCloudflareZeroTrustCasbFindingTypeFilterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-cloudflare.dataCloudflareZeroTrustCasbFindingType.DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataCloudflareZeroTrustCasbFindingTypeFilterOutputReference_Override(d DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-cloudflare.dataCloudflareZeroTrustCasbFindingType.DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference)SetDirection(val *string) {
	if err := j.validateSetDirectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"direction",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference)SetMaxUpdatedAt(val *string) {
	if err := j.validateSetMaxUpdatedAtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxUpdatedAt",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference)SetMinUpdatedAt(val *string) {
	if err := j.validateSetMinUpdatedAtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minUpdatedAt",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference)SetOrder(val *string) {
	if err := j.validateSetOrderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"order",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference)SetOrigins(val *[]*string) {
	if err := j.validateSetOriginsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"origins",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference)SetQuarantined(val interface{}) {
	if err := j.validateSetQuarantinedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"quarantined",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference)SetSearch(val *string) {
	if err := j.validateSetSearchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"search",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference)SetSeverities(val *[]*string) {
	if err := j.validateSetSeveritiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"severities",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference)SetTypes(val *[]*string) {
	if err := j.validateSetTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"types",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference)SetVendors(val *[]*string) {
	if err := j.validateSetVendorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vendors",
		val,
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference) ResetDirection() {
	_jsii_.InvokeVoid(
		d,
		"resetDirection",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference) ResetMaxUpdatedAt() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxUpdatedAt",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference) ResetMinUpdatedAt() {
	_jsii_.InvokeVoid(
		d,
		"resetMinUpdatedAt",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference) ResetOrder() {
	_jsii_.InvokeVoid(
		d,
		"resetOrder",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference) ResetOrigins() {
	_jsii_.InvokeVoid(
		d,
		"resetOrigins",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference) ResetQuarantined() {
	_jsii_.InvokeVoid(
		d,
		"resetQuarantined",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference) ResetSearch() {
	_jsii_.InvokeVoid(
		d,
		"resetSearch",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference) ResetSeverities() {
	_jsii_.InvokeVoid(
		d,
		"resetSeverities",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference) ResetTypes() {
	_jsii_.InvokeVoid(
		d,
		"resetTypes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference) ResetVendors() {
	_jsii_.InvokeVoid(
		d,
		"resetVendors",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataCloudflareZeroTrustCasbFindingTypeFilterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

