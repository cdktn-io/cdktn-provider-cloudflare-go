// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package r2bucketlock

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_R2BucketLockRulesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_R2BucketLockRulesList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_R2BucketLockRulesList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_R2BucketLockRulesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_R2BucketLockRulesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_R2BucketLockRulesList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_R2BucketLockRulesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewR2BucketLockRulesListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

