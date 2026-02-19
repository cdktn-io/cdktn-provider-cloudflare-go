// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package accountmember

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AccountMemberPoliciesResourceGroupsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AccountMemberPoliciesResourceGroupsList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AccountMemberPoliciesResourceGroupsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AccountMemberPoliciesResourceGroupsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AccountMemberPoliciesResourceGroupsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AccountMemberPoliciesResourceGroupsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AccountMemberPoliciesResourceGroupsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAccountMemberPoliciesResourceGroupsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

