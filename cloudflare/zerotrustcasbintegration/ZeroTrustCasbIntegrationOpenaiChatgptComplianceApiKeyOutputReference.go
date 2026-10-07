// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package zerotrustcasbintegration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v16/zerotrustcasbintegration/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference interface {
	cdktn.ComplexObject
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	AdminApiKey() *string
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	SetAdminApiKey(val *string)
	AdminApiKeyInput() *string
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
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	ComplianceApiKey() *string
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	SetComplianceApiKey(val *string)
	ComplianceApiKeyInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	OrganizationId() *string
	SetOrganizationId(val *string)
	OrganizationIdInput() *string
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	ProjectApiKey() *string
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	SetProjectApiKey(val *string)
	ProjectApiKeyInput() *string
	ProjectId() *string
	SetProjectId(val *string)
	ProjectIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	WorkspaceId() *string
	SetWorkspaceId(val *string)
	WorkspaceIdInput() *string
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
	ResetProjectApiKey()
	ResetProjectId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference
type jsiiProxy_ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference) AdminApiKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminApiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference) AdminApiKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminApiKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference) ComplianceApiKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complianceApiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference) ComplianceApiKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complianceApiKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference) OrganizationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference) OrganizationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference) ProjectApiKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectApiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference) ProjectApiKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectApiKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference) WorkspaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference) WorkspaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceIdInput",
		&returns,
	)
	return returns
}


func NewZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference {
	_init_.Initialize()

	if err := validateNewZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-cloudflare.zeroTrustCasbIntegration.ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference_Override(z ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-cloudflare.zeroTrustCasbIntegration.ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		z,
	)
}

func (j *jsiiProxy_ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference)SetAdminApiKey(val *string) {
	if err := j.validateSetAdminApiKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminApiKey",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference)SetComplianceApiKey(val *string) {
	if err := j.validateSetComplianceApiKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complianceApiKey",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference)SetOrganizationId(val *string) {
	if err := j.validateSetOrganizationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organizationId",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference)SetProjectApiKey(val *string) {
	if err := j.validateSetProjectApiKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectApiKey",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference)SetWorkspaceId(val *string) {
	if err := j.validateSetWorkspaceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workspaceId",
		val,
	)
}

func (z *jsiiProxy_ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := z.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		z,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := z.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		z,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := z.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		z,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := z.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		z,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := z.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		z,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := z.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		z,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := z.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		z,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := z.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		z,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := z.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		z,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := z.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference) ResetProjectApiKey() {
	_jsii_.InvokeVoid(
		z,
		"resetProjectApiKey",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference) ResetProjectId() {
	_jsii_.InvokeVoid(
		z,
		"resetProjectId",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := z.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		z,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustCasbIntegrationOpenaiChatgptComplianceApiKeyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

