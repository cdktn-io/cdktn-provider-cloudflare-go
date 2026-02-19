// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datacloudflarequeue

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataCloudflareQueueProducersList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataCloudflareQueueProducersList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataCloudflareQueueProducersList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataCloudflareQueueProducersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataCloudflareQueueProducersList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataCloudflareQueueProducersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataCloudflareQueueProducersListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

