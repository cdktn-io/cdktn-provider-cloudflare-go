// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package workersdeployment

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WorkersDeploymentVersionsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (w *jsiiProxy_WorkersDeploymentVersionsList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WorkersDeploymentVersionsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WorkersDeploymentVersionsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WorkersDeploymentVersionsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WorkersDeploymentVersionsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WorkersDeploymentVersionsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWorkersDeploymentVersionsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

