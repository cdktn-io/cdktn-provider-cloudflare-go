// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package queue

// Building without runtime type checking enabled, so all the below just return nil

func (q *jsiiProxy_QueueConsumersSettingsWebhooksList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (q *jsiiProxy_QueueConsumersSettingsWebhooksList) validateGetParameters(index *float64) error {
	return nil
}

func (q *jsiiProxy_QueueConsumersSettingsWebhooksList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_QueueConsumersSettingsWebhooksList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_QueueConsumersSettingsWebhooksList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_QueueConsumersSettingsWebhooksList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewQueueConsumersSettingsWebhooksListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

