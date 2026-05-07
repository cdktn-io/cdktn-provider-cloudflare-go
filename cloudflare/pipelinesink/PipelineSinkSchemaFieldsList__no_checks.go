// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package pipelinesink

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PipelineSinkSchemaFieldsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_PipelineSinkSchemaFieldsList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PipelineSinkSchemaFieldsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PipelineSinkSchemaFieldsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PipelineSinkSchemaFieldsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PipelineSinkSchemaFieldsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PipelineSinkSchemaFieldsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPipelineSinkSchemaFieldsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

