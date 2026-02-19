// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datacloudflarelists

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataCloudflareListsResultList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataCloudflareListsResultList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataCloudflareListsResultList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataCloudflareListsResultList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataCloudflareListsResultList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataCloudflareListsResultList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataCloudflareListsResultListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

