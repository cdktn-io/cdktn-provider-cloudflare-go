// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezerotrustdeviceposturerules

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v14/jsii"

	"github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v14/datacloudflarezerotrustdeviceposturerules/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference interface {
	cdktn.ComplexObject
	ActiveThreats() *float64
	CertificateId() *string
	CheckDisks() *[]*string
	CheckPrivateKey() cdktn.IResolvable
	Cn() *string
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
	ComplianceStatus() *string
	ConnectionId() *string
	CountOperator() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Domain() *string
	EidLastSeen() *string
	Enabled() cdktn.IResolvable
	Exists() cdktn.IResolvable
	ExtendedKeyUsage() *[]*string
	// Experimental.
	Fqn() *string
	Id() *string
	Infected() cdktn.IResolvable
	InternalValue() *DataCloudflareZeroTrustDevicePostureRulesResultInput
	SetInternalValue(val *DataCloudflareZeroTrustDevicePostureRulesResultInput)
	IsActive() cdktn.IResolvable
	IssueCount() *string
	LastSeen() *string
	Locations() DataCloudflareZeroTrustDevicePostureRulesResultInputLocationsOutputReference
	NetworkStatus() *string
	OperatingSystem() *string
	OperationalState() *string
	Operator() *string
	Os() *string
	OsDistroName() *string
	OsDistroRevision() *string
	OsVersionExtra() *string
	Overall() *string
	Path() *string
	RequireAll() cdktn.IResolvable
	RiskLevel() *string
	Score() *float64
	ScoreOperator() *string
	SensorConfig() *string
	Sha256() *string
	State() *string
	SubjectAlternativeNames() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Thumbprint() *string
	TotalScore() *float64
	UpdateWindowDays() *float64
	Version() *string
	VersionOperator() *string
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

// The jsii proxy struct for DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference
type jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) ActiveThreats() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"activeThreats",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) CertificateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) CheckDisks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"checkDisks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) CheckPrivateKey() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"checkPrivateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) Cn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) ComplianceStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complianceStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) ConnectionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) CountOperator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countOperator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) EidLastSeen() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eidLastSeen",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) Enabled() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) Exists() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"exists",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) ExtendedKeyUsage() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"extendedKeyUsage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) Infected() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"infected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) InternalValue() *DataCloudflareZeroTrustDevicePostureRulesResultInput {
	var returns *DataCloudflareZeroTrustDevicePostureRulesResultInput
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) IsActive() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"isActive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) IssueCount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issueCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) LastSeen() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastSeen",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) Locations() DataCloudflareZeroTrustDevicePostureRulesResultInputLocationsOutputReference {
	var returns DataCloudflareZeroTrustDevicePostureRulesResultInputLocationsOutputReference
	_jsii_.Get(
		j,
		"locations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) NetworkStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) OperatingSystem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatingSystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) OperationalState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationalState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) Os() *string {
	var returns *string
	_jsii_.Get(
		j,
		"os",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) OsDistroName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osDistroName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) OsDistroRevision() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osDistroRevision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) OsVersionExtra() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osVersionExtra",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) Overall() *string {
	var returns *string
	_jsii_.Get(
		j,
		"overall",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) RequireAll() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"requireAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) RiskLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"riskLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) Score() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"score",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) ScoreOperator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scoreOperator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) SensorConfig() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sensorConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) Sha256() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sha256",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) SubjectAlternativeNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subjectAlternativeNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) Thumbprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thumbprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) TotalScore() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalScore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) UpdateWindowDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"updateWindowDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) VersionOperator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionOperator",
		&returns,
	)
	return returns
}


func NewDataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference {
	_init_.Initialize()

	if err := validateNewDataCloudflareZeroTrustDevicePostureRulesResultInputOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-cloudflare.dataCloudflareZeroTrustDevicePostureRules.DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference_Override(d DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-cloudflare.dataCloudflareZeroTrustDevicePostureRules.DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference)SetInternalValue(val *DataCloudflareZeroTrustDevicePostureRulesResultInput) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataCloudflareZeroTrustDevicePostureRulesResultInputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

