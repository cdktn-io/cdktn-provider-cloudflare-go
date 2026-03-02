// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package certificatepack

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CertificatePackDcvDelegationRecordsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CertificatePackDcvDelegationRecordsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CertificatePackDcvDelegationRecordsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CertificatePackDcvDelegationRecordsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CertificatePackDcvDelegationRecordsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CertificatePackDcvDelegationRecordsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCertificatePackDcvDelegationRecordsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

