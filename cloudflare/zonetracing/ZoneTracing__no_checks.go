// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package zonetracing

// Building without runtime type checking enabled, so all the below just return nil

func (z *jsiiProxy_ZoneTracing) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (z *jsiiProxy_ZoneTracing) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (z *jsiiProxy_ZoneTracing) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (z *jsiiProxy_ZoneTracing) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (z *jsiiProxy_ZoneTracing) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (z *jsiiProxy_ZoneTracing) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (z *jsiiProxy_ZoneTracing) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (z *jsiiProxy_ZoneTracing) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (z *jsiiProxy_ZoneTracing) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (z *jsiiProxy_ZoneTracing) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (z *jsiiProxy_ZoneTracing) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (z *jsiiProxy_ZoneTracing) validateImportFromParameters(id *string) error {
	return nil
}

func (z *jsiiProxy_ZoneTracing) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (z *jsiiProxy_ZoneTracing) validateMarkWriteOnlyAttributeParameters(value interface{}) error {
	return nil
}

func (z *jsiiProxy_ZoneTracing) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (z *jsiiProxy_ZoneTracing) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (z *jsiiProxy_ZoneTracing) validateMoveToIdParameters(id *string) error {
	return nil
}

func (z *jsiiProxy_ZoneTracing) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (z *jsiiProxy_ZoneTracing) validateRegisterProviderFeatureUsageParameters(feature cdktn.ProviderFeature) error {
	return nil
}

func validateZoneTracing_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateZoneTracing_IsConstructParameters(x interface{}) error {
	return nil
}

func validateZoneTracing_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateZoneTracing_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_ZoneTracing) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ZoneTracing) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ZoneTracing) validateSetDestinationsParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_ZoneTracing) validateSetEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ZoneTracing) validateSetForwardContextParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ZoneTracing) validateSetLifecycleParameters(val *cdktn.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_ZoneTracing) validateSetPersistParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ZoneTracing) validateSetPropagationPolicyParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ZoneTracing) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_ZoneTracing) validateSetSamplingRatioParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ZoneTracing) validateSetZoneIdParameters(val *string) error {
	return nil
}

func validateNewZoneTracingParameters(scope constructs.Construct, id *string, config *ZoneTracingConfig) error {
	return nil
}

