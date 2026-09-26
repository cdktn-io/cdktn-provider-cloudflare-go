// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datacloudflareworker

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataCloudflareWorkerPreviewsBaseConfigEnvMap) validateGetParameters(key *string) error {
	return nil
}

func (d *jsiiProxy_DataCloudflareWorkerPreviewsBaseConfigEnvMap) validateInterpolationForAttributeParameters(property *string) error {
	return nil
}

func (d *jsiiProxy_DataCloudflareWorkerPreviewsBaseConfigEnvMap) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataCloudflareWorkerPreviewsBaseConfigEnvMap) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataCloudflareWorkerPreviewsBaseConfigEnvMap) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func validateNewDataCloudflareWorkerPreviewsBaseConfigEnvMapParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) error {
	return nil
}

