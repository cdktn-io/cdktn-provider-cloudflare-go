// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflareoauthclients

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v15/jsii"

	"github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v15/datacloudflareoauthclients/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataCloudflareOauthClientsResultOutputReference interface {
	cdktn.ComplexObject
	AllowedCorsOrigins() *[]*string
	ClientId() *string
	ClientName() *string
	ClientUri() *string
	ClientUriVerification() DataCloudflareOauthClientsResultClientUriVerificationOutputReference
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
	CreatedAt() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	GrantTypes() *[]*string
	HasRotatedSecret() cdktn.IResolvable
	InternalValue() *DataCloudflareOauthClientsResult
	SetInternalValue(val *DataCloudflareOauthClientsResult)
	LogoUri() *string
	PolicyUri() *string
	PostLogoutRedirectUris() *[]*string
	PromotedAt() *string
	RedirectUris() *[]*string
	ResponseTypes() *[]*string
	Scopes() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TokenEndpointAuthMethod() *string
	TosUri() *string
	UpdatedAt() *string
	Visibility() *string
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

// The jsii proxy struct for DataCloudflareOauthClientsResultOutputReference
type jsiiProxy_DataCloudflareOauthClientsResultOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataCloudflareOauthClientsResultOutputReference) AllowedCorsOrigins() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedCorsOrigins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareOauthClientsResultOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareOauthClientsResultOutputReference) ClientName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareOauthClientsResultOutputReference) ClientUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareOauthClientsResultOutputReference) ClientUriVerification() DataCloudflareOauthClientsResultClientUriVerificationOutputReference {
	var returns DataCloudflareOauthClientsResultClientUriVerificationOutputReference
	_jsii_.Get(
		j,
		"clientUriVerification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareOauthClientsResultOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareOauthClientsResultOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareOauthClientsResultOutputReference) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareOauthClientsResultOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareOauthClientsResultOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareOauthClientsResultOutputReference) GrantTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"grantTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareOauthClientsResultOutputReference) HasRotatedSecret() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"hasRotatedSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareOauthClientsResultOutputReference) InternalValue() *DataCloudflareOauthClientsResult {
	var returns *DataCloudflareOauthClientsResult
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareOauthClientsResultOutputReference) LogoUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareOauthClientsResultOutputReference) PolicyUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareOauthClientsResultOutputReference) PostLogoutRedirectUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"postLogoutRedirectUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareOauthClientsResultOutputReference) PromotedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"promotedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareOauthClientsResultOutputReference) RedirectUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"redirectUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareOauthClientsResultOutputReference) ResponseTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"responseTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareOauthClientsResultOutputReference) Scopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareOauthClientsResultOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareOauthClientsResultOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareOauthClientsResultOutputReference) TokenEndpointAuthMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenEndpointAuthMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareOauthClientsResultOutputReference) TosUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tosUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareOauthClientsResultOutputReference) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareOauthClientsResultOutputReference) Visibility() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visibility",
		&returns,
	)
	return returns
}


func NewDataCloudflareOauthClientsResultOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataCloudflareOauthClientsResultOutputReference {
	_init_.Initialize()

	if err := validateNewDataCloudflareOauthClientsResultOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareOauthClientsResultOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-cloudflare.dataCloudflareOauthClients.DataCloudflareOauthClientsResultOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataCloudflareOauthClientsResultOutputReference_Override(d DataCloudflareOauthClientsResultOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-cloudflare.dataCloudflareOauthClients.DataCloudflareOauthClientsResultOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareOauthClientsResultOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareOauthClientsResultOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareOauthClientsResultOutputReference)SetInternalValue(val *DataCloudflareOauthClientsResult) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareOauthClientsResultOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareOauthClientsResultOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataCloudflareOauthClientsResultOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareOauthClientsResultOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataCloudflareOauthClientsResultOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataCloudflareOauthClientsResultOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataCloudflareOauthClientsResultOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataCloudflareOauthClientsResultOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataCloudflareOauthClientsResultOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataCloudflareOauthClientsResultOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataCloudflareOauthClientsResultOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataCloudflareOauthClientsResultOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataCloudflareOauthClientsResultOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareOauthClientsResultOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataCloudflareOauthClientsResultOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataCloudflareOauthClientsResultOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

