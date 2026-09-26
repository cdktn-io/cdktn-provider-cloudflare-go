// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package queueconsumer

// Building without runtime type checking enabled, so all the below just return nil

func (q *jsiiProxy_QueueConsumerSettingsPagerdutyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (q *jsiiProxy_QueueConsumerSettingsPagerdutyList) validateGetParameters(index *float64) error {
	return nil
}

func (q *jsiiProxy_QueueConsumerSettingsPagerdutyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_QueueConsumerSettingsPagerdutyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_QueueConsumerSettingsPagerdutyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_QueueConsumerSettingsPagerdutyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_QueueConsumerSettingsPagerdutyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewQueueConsumerSettingsPagerdutyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

