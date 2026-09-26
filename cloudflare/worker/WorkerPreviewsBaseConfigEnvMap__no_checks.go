// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package worker

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WorkerPreviewsBaseConfigEnvMap) validateGetParameters(key *string) error {
	return nil
}

func (w *jsiiProxy_WorkerPreviewsBaseConfigEnvMap) validateInterpolationForAttributeParameters(property *string) error {
	return nil
}

func (w *jsiiProxy_WorkerPreviewsBaseConfigEnvMap) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigEnvMap) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigEnvMap) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigEnvMap) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func validateNewWorkerPreviewsBaseConfigEnvMapParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) error {
	return nil
}

