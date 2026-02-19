// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package tokenvalidationrules

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TokenValidationRulesSelectorExcludeList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (t *jsiiProxy_TokenValidationRulesSelectorExcludeList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TokenValidationRulesSelectorExcludeList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TokenValidationRulesSelectorExcludeList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_TokenValidationRulesSelectorExcludeList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TokenValidationRulesSelectorExcludeList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TokenValidationRulesSelectorExcludeList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTokenValidationRulesSelectorExcludeListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

