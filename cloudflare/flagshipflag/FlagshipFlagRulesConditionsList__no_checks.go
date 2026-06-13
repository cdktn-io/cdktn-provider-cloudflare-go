// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package flagshipflag

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FlagshipFlagRulesConditionsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (f *jsiiProxy_FlagshipFlagRulesConditionsList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FlagshipFlagRulesConditionsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FlagshipFlagRulesConditionsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_FlagshipFlagRulesConditionsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FlagshipFlagRulesConditionsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FlagshipFlagRulesConditionsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFlagshipFlagRulesConditionsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

