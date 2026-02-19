// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package queue

// Building without runtime type checking enabled, so all the below just return nil

func (q *jsiiProxy_QueueProducersList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (q *jsiiProxy_QueueProducersList) validateGetParameters(index *float64) error {
	return nil
}

func (q *jsiiProxy_QueueProducersList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_QueueProducersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_QueueProducersList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_QueueProducersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewQueueProducersListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

