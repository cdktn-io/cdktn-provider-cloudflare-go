// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package secretsstore

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SecretsStore) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (s *jsiiProxy_SecretsStore) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (s *jsiiProxy_SecretsStore) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SecretsStore) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SecretsStore) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SecretsStore) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SecretsStore) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SecretsStore) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SecretsStore) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SecretsStore) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SecretsStore) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SecretsStore) validateImportFromParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_SecretsStore) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SecretsStore) validateMarkWriteOnlyAttributeParameters(value interface{}) error {
	return nil
}

func (s *jsiiProxy_SecretsStore) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_SecretsStore) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (s *jsiiProxy_SecretsStore) validateMoveToIdParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_SecretsStore) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (s *jsiiProxy_SecretsStore) validateRegisterProviderFeatureUsageParameters(feature cdktn.ProviderFeature) error {
	return nil
}

func validateSecretsStore_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateSecretsStore_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSecretsStore_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateSecretsStore_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_SecretsStore) validateSetAccountIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SecretsStore) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SecretsStore) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SecretsStore) validateSetLifecycleParameters(val *cdktn.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_SecretsStore) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SecretsStore) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func validateNewSecretsStoreParameters(scope constructs.Construct, id *string, config *SecretsStoreConfig) error {
	return nil
}

