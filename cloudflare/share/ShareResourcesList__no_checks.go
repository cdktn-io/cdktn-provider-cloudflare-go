// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package share

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ShareResourcesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_ShareResourcesList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_ShareResourcesList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ShareResourcesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ShareResourcesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ShareResourcesList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ShareResourcesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewShareResourcesListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

