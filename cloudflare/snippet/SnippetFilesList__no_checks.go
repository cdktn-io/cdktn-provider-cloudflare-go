// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package snippet

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SnippetFilesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SnippetFilesList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SnippetFilesList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SnippetFilesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SnippetFilesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SnippetFilesList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SnippetFilesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSnippetFilesListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

