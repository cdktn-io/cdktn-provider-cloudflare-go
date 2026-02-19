// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package zerotrustdlpentry

// Building without runtime type checking enabled, so all the below just return nil

func (z *jsiiProxy_ZeroTrustDlpEntryProfilesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (z *jsiiProxy_ZeroTrustDlpEntryProfilesList) validateGetParameters(index *float64) error {
	return nil
}

func (z *jsiiProxy_ZeroTrustDlpEntryProfilesList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ZeroTrustDlpEntryProfilesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ZeroTrustDlpEntryProfilesList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ZeroTrustDlpEntryProfilesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewZeroTrustDlpEntryProfilesListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

