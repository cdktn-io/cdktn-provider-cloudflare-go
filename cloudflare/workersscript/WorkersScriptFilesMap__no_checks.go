// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package workersscript

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WorkersScriptFilesMap) validateGetParameters(key *string) error {
	return nil
}

func (w *jsiiProxy_WorkersScriptFilesMap) validateInterpolationForAttributeParameters(property *string) error {
	return nil
}

func (w *jsiiProxy_WorkersScriptFilesMap) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WorkersScriptFilesMap) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WorkersScriptFilesMap) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WorkersScriptFilesMap) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func validateNewWorkersScriptFilesMapParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) error {
	return nil
}

