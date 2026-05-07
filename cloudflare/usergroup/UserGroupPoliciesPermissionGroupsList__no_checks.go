// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package usergroup

// Building without runtime type checking enabled, so all the below just return nil

func (u *jsiiProxy_UserGroupPoliciesPermissionGroupsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (u *jsiiProxy_UserGroupPoliciesPermissionGroupsList) validateGetParameters(index *float64) error {
	return nil
}

func (u *jsiiProxy_UserGroupPoliciesPermissionGroupsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_UserGroupPoliciesPermissionGroupsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_UserGroupPoliciesPermissionGroupsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_UserGroupPoliciesPermissionGroupsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_UserGroupPoliciesPermissionGroupsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewUserGroupPoliciesPermissionGroupsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

