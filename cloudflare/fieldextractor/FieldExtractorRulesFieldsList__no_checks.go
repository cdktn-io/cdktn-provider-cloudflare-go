// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package fieldextractor

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FieldExtractorRulesFieldsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (f *jsiiProxy_FieldExtractorRulesFieldsList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FieldExtractorRulesFieldsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FieldExtractorRulesFieldsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_FieldExtractorRulesFieldsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FieldExtractorRulesFieldsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FieldExtractorRulesFieldsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFieldExtractorRulesFieldsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

