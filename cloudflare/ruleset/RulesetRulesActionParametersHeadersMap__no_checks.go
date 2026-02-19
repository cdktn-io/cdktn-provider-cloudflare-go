// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package ruleset

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RulesetRulesActionParametersHeadersMap) validateGetParameters(key *string) error {
	return nil
}

func (r *jsiiProxy_RulesetRulesActionParametersHeadersMap) validateInterpolationForAttributeParameters(property *string) error {
	return nil
}

func (r *jsiiProxy_RulesetRulesActionParametersHeadersMap) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RulesetRulesActionParametersHeadersMap) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RulesetRulesActionParametersHeadersMap) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RulesetRulesActionParametersHeadersMap) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func validateNewRulesetRulesActionParametersHeadersMapParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) error {
	return nil
}

