// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package snippetrules

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SnippetRulesRulesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SnippetRulesRulesList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SnippetRulesRulesList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SnippetRulesRulesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SnippetRulesRulesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SnippetRulesRulesList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SnippetRulesRulesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSnippetRulesRulesListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

