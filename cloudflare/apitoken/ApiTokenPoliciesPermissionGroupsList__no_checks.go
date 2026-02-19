// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package apitoken

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApiTokenPoliciesPermissionGroupsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_ApiTokenPoliciesPermissionGroupsList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_ApiTokenPoliciesPermissionGroupsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ApiTokenPoliciesPermissionGroupsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApiTokenPoliciesPermissionGroupsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApiTokenPoliciesPermissionGroupsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ApiTokenPoliciesPermissionGroupsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewApiTokenPoliciesPermissionGroupsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

