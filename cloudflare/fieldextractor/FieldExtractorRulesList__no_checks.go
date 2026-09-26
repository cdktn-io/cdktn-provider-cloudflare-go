// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package fieldextractor

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FieldExtractorRulesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (f *jsiiProxy_FieldExtractorRulesList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FieldExtractorRulesList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FieldExtractorRulesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_FieldExtractorRulesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FieldExtractorRulesList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FieldExtractorRulesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFieldExtractorRulesListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

