// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package streamaudiotrack

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StreamAudioTrackAudioList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_StreamAudioTrackAudioList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_StreamAudioTrackAudioList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_StreamAudioTrackAudioList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_StreamAudioTrackAudioList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_StreamAudioTrackAudioList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewStreamAudioTrackAudioListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

