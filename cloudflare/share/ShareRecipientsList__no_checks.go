// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package share

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ShareRecipientsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_ShareRecipientsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_ShareRecipientsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ShareRecipientsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ShareRecipientsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ShareRecipientsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ShareRecipientsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewShareRecipientsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

