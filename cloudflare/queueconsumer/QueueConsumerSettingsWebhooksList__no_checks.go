// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package queueconsumer

// Building without runtime type checking enabled, so all the below just return nil

func (q *jsiiProxy_QueueConsumerSettingsWebhooksList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (q *jsiiProxy_QueueConsumerSettingsWebhooksList) validateGetParameters(index *float64) error {
	return nil
}

func (q *jsiiProxy_QueueConsumerSettingsWebhooksList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_QueueConsumerSettingsWebhooksList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_QueueConsumerSettingsWebhooksList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_QueueConsumerSettingsWebhooksList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_QueueConsumerSettingsWebhooksList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewQueueConsumerSettingsWebhooksListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

