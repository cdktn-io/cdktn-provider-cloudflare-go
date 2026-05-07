// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package usergroup

// Building without runtime type checking enabled, so all the below just return nil

func (u *jsiiProxy_UserGroupPoliciesResourceGroupsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (u *jsiiProxy_UserGroupPoliciesResourceGroupsList) validateGetParameters(index *float64) error {
	return nil
}

func (u *jsiiProxy_UserGroupPoliciesResourceGroupsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_UserGroupPoliciesResourceGroupsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_UserGroupPoliciesResourceGroupsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_UserGroupPoliciesResourceGroupsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_UserGroupPoliciesResourceGroupsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewUserGroupPoliciesResourceGroupsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

