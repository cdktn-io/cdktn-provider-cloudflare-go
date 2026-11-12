// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package logexplorerdataset

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LogExplorerDatasetFieldsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LogExplorerDatasetFieldsList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LogExplorerDatasetFieldsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LogExplorerDatasetFieldsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LogExplorerDatasetFieldsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LogExplorerDatasetFieldsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LogExplorerDatasetFieldsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLogExplorerDatasetFieldsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

