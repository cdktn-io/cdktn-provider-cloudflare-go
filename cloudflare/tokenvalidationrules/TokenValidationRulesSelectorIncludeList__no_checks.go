// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package tokenvalidationrules

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TokenValidationRulesSelectorIncludeList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (t *jsiiProxy_TokenValidationRulesSelectorIncludeList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TokenValidationRulesSelectorIncludeList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TokenValidationRulesSelectorIncludeList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_TokenValidationRulesSelectorIncludeList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TokenValidationRulesSelectorIncludeList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TokenValidationRulesSelectorIncludeList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTokenValidationRulesSelectorIncludeListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

