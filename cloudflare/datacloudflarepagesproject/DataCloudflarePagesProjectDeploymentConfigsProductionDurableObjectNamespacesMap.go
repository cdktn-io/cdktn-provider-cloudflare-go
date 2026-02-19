// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflarepagesproject

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v14/jsii"

	"github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v14/datacloudflarepagesproject/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataCloudflarePagesProjectDeploymentConfigsProductionDurableObjectNamespacesMap interface {
	cdktn.ComplexMap
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktn.IInterpolatingParent
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	Get(key *string) DataCloudflarePagesProjectDeploymentConfigsProductionDurableObjectNamespacesOutputReference
	// Experimental.
	InterpolationForAttribute(property *string) cdktn.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataCloudflarePagesProjectDeploymentConfigsProductionDurableObjectNamespacesMap
type jsiiProxy_DataCloudflarePagesProjectDeploymentConfigsProductionDurableObjectNamespacesMap struct {
	internal.Type__cdktnComplexMap
}

func (j *jsiiProxy_DataCloudflarePagesProjectDeploymentConfigsProductionDurableObjectNamespacesMap) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectDeploymentConfigsProductionDurableObjectNamespacesMap) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectDeploymentConfigsProductionDurableObjectNamespacesMap) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectDeploymentConfigsProductionDurableObjectNamespacesMap) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataCloudflarePagesProjectDeploymentConfigsProductionDurableObjectNamespacesMap(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataCloudflarePagesProjectDeploymentConfigsProductionDurableObjectNamespacesMap {
	_init_.Initialize()

	if err := validateNewDataCloudflarePagesProjectDeploymentConfigsProductionDurableObjectNamespacesMapParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflarePagesProjectDeploymentConfigsProductionDurableObjectNamespacesMap{}

	_jsii_.Create(
		"@cdktn/provider-cloudflare.dataCloudflarePagesProject.DataCloudflarePagesProjectDeploymentConfigsProductionDurableObjectNamespacesMap",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataCloudflarePagesProjectDeploymentConfigsProductionDurableObjectNamespacesMap_Override(d DataCloudflarePagesProjectDeploymentConfigsProductionDurableObjectNamespacesMap, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-cloudflare.dataCloudflarePagesProject.DataCloudflarePagesProjectDeploymentConfigsProductionDurableObjectNamespacesMap",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataCloudflarePagesProjectDeploymentConfigsProductionDurableObjectNamespacesMap)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflarePagesProjectDeploymentConfigsProductionDurableObjectNamespacesMap)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataCloudflarePagesProjectDeploymentConfigsProductionDurableObjectNamespacesMap) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflarePagesProjectDeploymentConfigsProductionDurableObjectNamespacesMap) Get(key *string) DataCloudflarePagesProjectDeploymentConfigsProductionDurableObjectNamespacesOutputReference {
	if err := d.validateGetParameters(key); err != nil {
		panic(err)
	}
	var returns DataCloudflarePagesProjectDeploymentConfigsProductionDurableObjectNamespacesOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflarePagesProjectDeploymentConfigsProductionDurableObjectNamespacesMap) InterpolationForAttribute(property *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflarePagesProjectDeploymentConfigsProductionDurableObjectNamespacesMap) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataCloudflarePagesProjectDeploymentConfigsProductionDurableObjectNamespacesMap) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

