// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package share

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Share) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (s *jsiiProxy_Share) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (s *jsiiProxy_Share) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Share) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Share) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Share) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Share) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Share) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Share) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Share) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Share) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Share) validateImportFromParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_Share) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Share) validateMarkWriteOnlyAttributeParameters(value interface{}) error {
	return nil
}

func (s *jsiiProxy_Share) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_Share) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (s *jsiiProxy_Share) validateMoveToIdParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_Share) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (s *jsiiProxy_Share) validatePutRecipientsParameters(value interface{}) error {
	return nil
}

func (s *jsiiProxy_Share) validatePutResourcesParameters(value interface{}) error {
	return nil
}

func (s *jsiiProxy_Share) validateRegisterProviderFeatureUsageParameters(feature cdktn.ProviderFeature) error {
	return nil
}

func validateShare_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateShare_IsConstructParameters(x interface{}) error {
	return nil
}

func validateShare_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateShare_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Share) validateSetAccountIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Share) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Share) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Share) validateSetLifecycleParameters(val *cdktn.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Share) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Share) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func validateNewShareParameters(scope constructs.Construct, id *string, config *ShareConfig) error {
	return nil
}

