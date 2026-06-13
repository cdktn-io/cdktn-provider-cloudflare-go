// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datacloudflareshare

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataCloudflareShareResourcesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataCloudflareShareResourcesList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataCloudflareShareResourcesList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataCloudflareShareResourcesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataCloudflareShareResourcesList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataCloudflareShareResourcesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataCloudflareShareResourcesListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

