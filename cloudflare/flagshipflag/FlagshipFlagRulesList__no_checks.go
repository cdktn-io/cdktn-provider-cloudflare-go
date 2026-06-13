// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package flagshipflag

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FlagshipFlagRulesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (f *jsiiProxy_FlagshipFlagRulesList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FlagshipFlagRulesList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FlagshipFlagRulesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_FlagshipFlagRulesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FlagshipFlagRulesList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FlagshipFlagRulesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFlagshipFlagRulesListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

