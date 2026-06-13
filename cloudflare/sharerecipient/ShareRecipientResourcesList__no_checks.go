// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package sharerecipient

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ShareRecipientResourcesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_ShareRecipientResourcesList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_ShareRecipientResourcesList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ShareRecipientResourcesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ShareRecipientResourcesList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ShareRecipientResourcesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewShareRecipientResourcesListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

