// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package worker

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WorkerPreviewsBaseConfigPlacementTargetList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (w *jsiiProxy_WorkerPreviewsBaseConfigPlacementTargetList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WorkerPreviewsBaseConfigPlacementTargetList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigPlacementTargetList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigPlacementTargetList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigPlacementTargetList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WorkerPreviewsBaseConfigPlacementTargetList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWorkerPreviewsBaseConfigPlacementTargetListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

