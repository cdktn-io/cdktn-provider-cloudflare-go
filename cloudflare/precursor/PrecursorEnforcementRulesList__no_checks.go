// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package precursor

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PrecursorEnforcementRulesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_PrecursorEnforcementRulesList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PrecursorEnforcementRulesList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PrecursorEnforcementRulesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PrecursorEnforcementRulesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PrecursorEnforcementRulesList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PrecursorEnforcementRulesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPrecursorEnforcementRulesListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

