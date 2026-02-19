// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package certificatepack

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CertificatePackCertificatesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CertificatePackCertificatesList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CertificatePackCertificatesList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CertificatePackCertificatesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CertificatePackCertificatesList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CertificatePackCertificatesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCertificatePackCertificatesListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

