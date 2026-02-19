// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package accounttoken

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AccountTokenPoliciesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AccountTokenPoliciesList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AccountTokenPoliciesList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AccountTokenPoliciesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AccountTokenPoliciesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AccountTokenPoliciesList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AccountTokenPoliciesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAccountTokenPoliciesListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

