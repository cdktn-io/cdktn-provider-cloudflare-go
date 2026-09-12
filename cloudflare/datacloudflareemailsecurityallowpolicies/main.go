// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflareemailsecurityallowpolicies

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktn/provider-cloudflare.dataCloudflareEmailSecurityAllowPolicies.DataCloudflareEmailSecurityAllowPolicies",
		reflect.TypeOf((*DataCloudflareEmailSecurityAllowPolicies)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accountId", GoGetter: "AccountId"},
			_jsii_.MemberProperty{JsiiProperty: "accountIdInput", GoGetter: "AccountIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "direction", GoGetter: "Direction"},
			_jsii_.MemberProperty{JsiiProperty: "directionInput", GoGetter: "DirectionInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isAcceptableSender", GoGetter: "IsAcceptableSender"},
			_jsii_.MemberProperty{JsiiProperty: "isAcceptableSenderInput", GoGetter: "IsAcceptableSenderInput"},
			_jsii_.MemberProperty{JsiiProperty: "isExemptRecipient", GoGetter: "IsExemptRecipient"},
			_jsii_.MemberProperty{JsiiProperty: "isExemptRecipientInput", GoGetter: "IsExemptRecipientInput"},
			_jsii_.MemberProperty{JsiiProperty: "isTrustedSender", GoGetter: "IsTrustedSender"},
			_jsii_.MemberProperty{JsiiProperty: "isTrustedSenderInput", GoGetter: "IsTrustedSenderInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "maxItems", GoGetter: "MaxItems"},
			_jsii_.MemberProperty{JsiiProperty: "maxItemsInput", GoGetter: "MaxItemsInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "order", GoGetter: "Order"},
			_jsii_.MemberProperty{JsiiProperty: "orderInput", GoGetter: "OrderInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "pattern", GoGetter: "Pattern"},
			_jsii_.MemberProperty{JsiiProperty: "patternInput", GoGetter: "PatternInput"},
			_jsii_.MemberProperty{JsiiProperty: "patternType", GoGetter: "PatternType"},
			_jsii_.MemberProperty{JsiiProperty: "patternTypeInput", GoGetter: "PatternTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "registerProviderFeatureUsage", GoMethod: "RegisterProviderFeatureUsage"},
			_jsii_.MemberMethod{JsiiMethod: "resetDirection", GoMethod: "ResetDirection"},
			_jsii_.MemberMethod{JsiiMethod: "resetIsAcceptableSender", GoMethod: "ResetIsAcceptableSender"},
			_jsii_.MemberMethod{JsiiMethod: "resetIsExemptRecipient", GoMethod: "ResetIsExemptRecipient"},
			_jsii_.MemberMethod{JsiiMethod: "resetIsTrustedSender", GoMethod: "ResetIsTrustedSender"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxItems", GoMethod: "ResetMaxItems"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrder", GoMethod: "ResetOrder"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPattern", GoMethod: "ResetPattern"},
			_jsii_.MemberMethod{JsiiMethod: "resetPatternType", GoMethod: "ResetPatternType"},
			_jsii_.MemberMethod{JsiiMethod: "resetSearch", GoMethod: "ResetSearch"},
			_jsii_.MemberMethod{JsiiMethod: "resetVerifySender", GoMethod: "ResetVerifySender"},
			_jsii_.MemberProperty{JsiiProperty: "result", GoGetter: "Result"},
			_jsii_.MemberProperty{JsiiProperty: "search", GoGetter: "Search"},
			_jsii_.MemberProperty{JsiiProperty: "searchInput", GoGetter: "SearchInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "verifySender", GoGetter: "VerifySender"},
			_jsii_.MemberProperty{JsiiProperty: "verifySenderInput", GoGetter: "VerifySenderInput"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_DataCloudflareEmailSecurityAllowPolicies{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-cloudflare.dataCloudflareEmailSecurityAllowPolicies.DataCloudflareEmailSecurityAllowPoliciesConfig",
		reflect.TypeOf((*DataCloudflareEmailSecurityAllowPoliciesConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-cloudflare.dataCloudflareEmailSecurityAllowPolicies.DataCloudflareEmailSecurityAllowPoliciesResult",
		reflect.TypeOf((*DataCloudflareEmailSecurityAllowPoliciesResult)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-cloudflare.dataCloudflareEmailSecurityAllowPolicies.DataCloudflareEmailSecurityAllowPoliciesResultList",
		reflect.TypeOf((*DataCloudflareEmailSecurityAllowPoliciesResultList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_DataCloudflareEmailSecurityAllowPoliciesResultList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-cloudflare.dataCloudflareEmailSecurityAllowPolicies.DataCloudflareEmailSecurityAllowPoliciesResultOutputReference",
		reflect.TypeOf((*DataCloudflareEmailSecurityAllowPoliciesResultOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "comments", GoGetter: "Comments"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isAcceptableSender", GoGetter: "IsAcceptableSender"},
			_jsii_.MemberProperty{JsiiProperty: "isExemptRecipient", GoGetter: "IsExemptRecipient"},
			_jsii_.MemberProperty{JsiiProperty: "isRecipient", GoGetter: "IsRecipient"},
			_jsii_.MemberProperty{JsiiProperty: "isRegex", GoGetter: "IsRegex"},
			_jsii_.MemberProperty{JsiiProperty: "isSender", GoGetter: "IsSender"},
			_jsii_.MemberProperty{JsiiProperty: "isSpoof", GoGetter: "IsSpoof"},
			_jsii_.MemberProperty{JsiiProperty: "isTrustedSender", GoGetter: "IsTrustedSender"},
			_jsii_.MemberProperty{JsiiProperty: "lastModified", GoGetter: "LastModified"},
			_jsii_.MemberProperty{JsiiProperty: "modifiedAt", GoGetter: "ModifiedAt"},
			_jsii_.MemberProperty{JsiiProperty: "pattern", GoGetter: "Pattern"},
			_jsii_.MemberProperty{JsiiProperty: "patternType", GoGetter: "PatternType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "verifySender", GoGetter: "VerifySender"},
		},
		func() interface{} {
			j := jsiiProxy_DataCloudflareEmailSecurityAllowPoliciesResultOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
}
