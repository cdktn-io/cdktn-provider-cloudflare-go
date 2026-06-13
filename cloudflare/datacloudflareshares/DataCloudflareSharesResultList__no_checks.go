// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datacloudflareshares

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataCloudflareSharesResultList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataCloudflareSharesResultList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataCloudflareSharesResultList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataCloudflareSharesResultList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataCloudflareSharesResultList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataCloudflareSharesResultList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataCloudflareSharesResultListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

