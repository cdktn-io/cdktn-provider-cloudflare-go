// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezerotrustaccessgroups

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v14/jsii"

	"github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v14/datacloudflarezerotrustaccessgroups/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference interface {
	cdktn.ComplexObject
	AnyValidServiceToken() DataCloudflareZeroTrustAccessGroupsResultRequireAnyValidServiceTokenOutputReference
	AuthContext() DataCloudflareZeroTrustAccessGroupsResultRequireAuthContextOutputReference
	AuthMethod() DataCloudflareZeroTrustAccessGroupsResultRequireAuthMethodOutputReference
	AzureAd() DataCloudflareZeroTrustAccessGroupsResultRequireAzureAdOutputReference
	Certificate() DataCloudflareZeroTrustAccessGroupsResultRequireCertificateOutputReference
	CommonName() DataCloudflareZeroTrustAccessGroupsResultRequireCommonNameOutputReference
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
	DevicePosture() DataCloudflareZeroTrustAccessGroupsResultRequireDevicePostureOutputReference
	Email() DataCloudflareZeroTrustAccessGroupsResultRequireEmailOutputReference
	EmailDomain() DataCloudflareZeroTrustAccessGroupsResultRequireEmailDomainOutputReference
	EmailList() DataCloudflareZeroTrustAccessGroupsResultRequireEmailListStructOutputReference
	Everyone() DataCloudflareZeroTrustAccessGroupsResultRequireEveryoneOutputReference
	ExternalEvaluation() DataCloudflareZeroTrustAccessGroupsResultRequireExternalEvaluationOutputReference
	// Experimental.
	Fqn() *string
	Geo() DataCloudflareZeroTrustAccessGroupsResultRequireGeoOutputReference
	GithubOrganization() DataCloudflareZeroTrustAccessGroupsResultRequireGithubOrganizationOutputReference
	Group() DataCloudflareZeroTrustAccessGroupsResultRequireGroupOutputReference
	Gsuite() DataCloudflareZeroTrustAccessGroupsResultRequireGsuiteOutputReference
	InternalValue() *DataCloudflareZeroTrustAccessGroupsResultRequire
	SetInternalValue(val *DataCloudflareZeroTrustAccessGroupsResultRequire)
	Ip() DataCloudflareZeroTrustAccessGroupsResultRequireIpOutputReference
	IpList() DataCloudflareZeroTrustAccessGroupsResultRequireIpListStructOutputReference
	LinkedAppToken() DataCloudflareZeroTrustAccessGroupsResultRequireLinkedAppTokenOutputReference
	LoginMethod() DataCloudflareZeroTrustAccessGroupsResultRequireLoginMethodOutputReference
	Oidc() DataCloudflareZeroTrustAccessGroupsResultRequireOidcOutputReference
	Okta() DataCloudflareZeroTrustAccessGroupsResultRequireOktaOutputReference
	Saml() DataCloudflareZeroTrustAccessGroupsResultRequireSamlOutputReference
	ServiceToken() DataCloudflareZeroTrustAccessGroupsResultRequireServiceTokenOutputReference
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

// The jsii proxy struct for DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference
type jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference) AnyValidServiceToken() DataCloudflareZeroTrustAccessGroupsResultRequireAnyValidServiceTokenOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultRequireAnyValidServiceTokenOutputReference
	_jsii_.Get(
		j,
		"anyValidServiceToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference) AuthContext() DataCloudflareZeroTrustAccessGroupsResultRequireAuthContextOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultRequireAuthContextOutputReference
	_jsii_.Get(
		j,
		"authContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference) AuthMethod() DataCloudflareZeroTrustAccessGroupsResultRequireAuthMethodOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultRequireAuthMethodOutputReference
	_jsii_.Get(
		j,
		"authMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference) AzureAd() DataCloudflareZeroTrustAccessGroupsResultRequireAzureAdOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultRequireAzureAdOutputReference
	_jsii_.Get(
		j,
		"azureAd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference) Certificate() DataCloudflareZeroTrustAccessGroupsResultRequireCertificateOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultRequireCertificateOutputReference
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference) CommonName() DataCloudflareZeroTrustAccessGroupsResultRequireCommonNameOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultRequireCommonNameOutputReference
	_jsii_.Get(
		j,
		"commonName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference) DevicePosture() DataCloudflareZeroTrustAccessGroupsResultRequireDevicePostureOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultRequireDevicePostureOutputReference
	_jsii_.Get(
		j,
		"devicePosture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference) Email() DataCloudflareZeroTrustAccessGroupsResultRequireEmailOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultRequireEmailOutputReference
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference) EmailDomain() DataCloudflareZeroTrustAccessGroupsResultRequireEmailDomainOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultRequireEmailDomainOutputReference
	_jsii_.Get(
		j,
		"emailDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference) EmailList() DataCloudflareZeroTrustAccessGroupsResultRequireEmailListStructOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultRequireEmailListStructOutputReference
	_jsii_.Get(
		j,
		"emailList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference) Everyone() DataCloudflareZeroTrustAccessGroupsResultRequireEveryoneOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultRequireEveryoneOutputReference
	_jsii_.Get(
		j,
		"everyone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference) ExternalEvaluation() DataCloudflareZeroTrustAccessGroupsResultRequireExternalEvaluationOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultRequireExternalEvaluationOutputReference
	_jsii_.Get(
		j,
		"externalEvaluation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference) Geo() DataCloudflareZeroTrustAccessGroupsResultRequireGeoOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultRequireGeoOutputReference
	_jsii_.Get(
		j,
		"geo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference) GithubOrganization() DataCloudflareZeroTrustAccessGroupsResultRequireGithubOrganizationOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultRequireGithubOrganizationOutputReference
	_jsii_.Get(
		j,
		"githubOrganization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference) Group() DataCloudflareZeroTrustAccessGroupsResultRequireGroupOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultRequireGroupOutputReference
	_jsii_.Get(
		j,
		"group",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference) Gsuite() DataCloudflareZeroTrustAccessGroupsResultRequireGsuiteOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultRequireGsuiteOutputReference
	_jsii_.Get(
		j,
		"gsuite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference) InternalValue() *DataCloudflareZeroTrustAccessGroupsResultRequire {
	var returns *DataCloudflareZeroTrustAccessGroupsResultRequire
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference) Ip() DataCloudflareZeroTrustAccessGroupsResultRequireIpOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultRequireIpOutputReference
	_jsii_.Get(
		j,
		"ip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference) IpList() DataCloudflareZeroTrustAccessGroupsResultRequireIpListStructOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultRequireIpListStructOutputReference
	_jsii_.Get(
		j,
		"ipList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference) LinkedAppToken() DataCloudflareZeroTrustAccessGroupsResultRequireLinkedAppTokenOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultRequireLinkedAppTokenOutputReference
	_jsii_.Get(
		j,
		"linkedAppToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference) LoginMethod() DataCloudflareZeroTrustAccessGroupsResultRequireLoginMethodOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultRequireLoginMethodOutputReference
	_jsii_.Get(
		j,
		"loginMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference) Oidc() DataCloudflareZeroTrustAccessGroupsResultRequireOidcOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultRequireOidcOutputReference
	_jsii_.Get(
		j,
		"oidc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference) Okta() DataCloudflareZeroTrustAccessGroupsResultRequireOktaOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultRequireOktaOutputReference
	_jsii_.Get(
		j,
		"okta",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference) Saml() DataCloudflareZeroTrustAccessGroupsResultRequireSamlOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultRequireSamlOutputReference
	_jsii_.Get(
		j,
		"saml",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference) ServiceToken() DataCloudflareZeroTrustAccessGroupsResultRequireServiceTokenOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultRequireServiceTokenOutputReference
	_jsii_.Get(
		j,
		"serviceToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataCloudflareZeroTrustAccessGroupsResultRequireOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference {
	_init_.Initialize()

	if err := validateNewDataCloudflareZeroTrustAccessGroupsResultRequireOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-cloudflare.dataCloudflareZeroTrustAccessGroups.DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataCloudflareZeroTrustAccessGroupsResultRequireOutputReference_Override(d DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-cloudflare.dataCloudflareZeroTrustAccessGroups.DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference)SetInternalValue(val *DataCloudflareZeroTrustAccessGroupsResultRequire) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultRequireOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

