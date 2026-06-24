// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package ruleset

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RulesetRulesActionParametersVaryHeadersMap) validateGetParameters(key *string) error {
	return nil
}

func (r *jsiiProxy_RulesetRulesActionParametersVaryHeadersMap) validateInterpolationForAttributeParameters(property *string) error {
	return nil
}

func (r *jsiiProxy_RulesetRulesActionParametersVaryHeadersMap) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RulesetRulesActionParametersVaryHeadersMap) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RulesetRulesActionParametersVaryHeadersMap) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RulesetRulesActionParametersVaryHeadersMap) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func validateNewRulesetRulesActionParametersVaryHeadersMapParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) error {
	return nil
}

