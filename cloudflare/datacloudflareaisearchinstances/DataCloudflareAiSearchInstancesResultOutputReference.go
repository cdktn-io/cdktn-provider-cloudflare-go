// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacloudflareaisearchinstances

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v14/jsii"

	"github.com/cdktn-io/cdktn-provider-cloudflare-go/cloudflare/v14/datacloudflareaisearchinstances/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataCloudflareAiSearchInstancesResultOutputReference interface {
	cdktn.ComplexObject
	AccountId() *string
	AccountTag() *string
	AiGatewayId() *string
	AisearchModel() *string
	Cache() cdktn.IResolvable
	CacheThreshold() *string
	Chunk() cdktn.IResolvable
	ChunkOverlap() *float64
	ChunkSize() *float64
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	CreatedAt() *string
	CreatedBy() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomMetadata() DataCloudflareAiSearchInstancesResultCustomMetadataList
	EmbeddingModel() *string
	Enable() cdktn.IResolvable
	EngineVersion() *float64
	// Experimental.
	Fqn() *string
	HybridSearchEnabled() cdktn.IResolvable
	Id() *string
	InternalId() *string
	InternalValue() *DataCloudflareAiSearchInstancesResult
	SetInternalValue(val *DataCloudflareAiSearchInstancesResult)
	LastActivity() *string
	MaxNumResults() *float64
	Metadata() DataCloudflareAiSearchInstancesResultMetadataOutputReference
	ModifiedAt() *string
	ModifiedBy() *string
	Paused() cdktn.IResolvable
	PublicEndpointId() *string
	PublicEndpointParams() DataCloudflareAiSearchInstancesResultPublicEndpointParamsOutputReference
	Reranking() cdktn.IResolvable
	RerankingModel() *string
	RewriteModel() *string
	RewriteQuery() cdktn.IResolvable
	ScoreThreshold() *float64
	Source() *string
	SourceParams() DataCloudflareAiSearchInstancesResultSourceParamsOutputReference
	Status() *string
	Summarization() cdktn.IResolvable
	SummarizationModel() *string
	SystemPromptAisearch() *string
	SystemPromptIndexSummarization() *string
	SystemPromptRewriteQuery() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TokenId() *string
	Type() *string
	VectorizeActiveNamespace() *string
	VectorizeName() *string
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataCloudflareAiSearchInstancesResultOutputReference
type jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) AccountTag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) AiGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aiGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) AisearchModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aisearchModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) Cache() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"cache",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) CacheThreshold() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) Chunk() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"chunk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) ChunkOverlap() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"chunkOverlap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) ChunkSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"chunkSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) CreatedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) CustomMetadata() DataCloudflareAiSearchInstancesResultCustomMetadataList {
	var returns DataCloudflareAiSearchInstancesResultCustomMetadataList
	_jsii_.Get(
		j,
		"customMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) EmbeddingModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"embeddingModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) Enable() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"enable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) EngineVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) HybridSearchEnabled() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"hybridSearchEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) InternalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"internalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) InternalValue() *DataCloudflareAiSearchInstancesResult {
	var returns *DataCloudflareAiSearchInstancesResult
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) LastActivity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastActivity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) MaxNumResults() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxNumResults",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) Metadata() DataCloudflareAiSearchInstancesResultMetadataOutputReference {
	var returns DataCloudflareAiSearchInstancesResultMetadataOutputReference
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) ModifiedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modifiedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) ModifiedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modifiedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) Paused() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"paused",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) PublicEndpointId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicEndpointId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) PublicEndpointParams() DataCloudflareAiSearchInstancesResultPublicEndpointParamsOutputReference {
	var returns DataCloudflareAiSearchInstancesResultPublicEndpointParamsOutputReference
	_jsii_.Get(
		j,
		"publicEndpointParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) Reranking() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"reranking",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) RerankingModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rerankingModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) RewriteModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rewriteModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) RewriteQuery() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"rewriteQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) ScoreThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scoreThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) SourceParams() DataCloudflareAiSearchInstancesResultSourceParamsOutputReference {
	var returns DataCloudflareAiSearchInstancesResultSourceParamsOutputReference
	_jsii_.Get(
		j,
		"sourceParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) Summarization() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"summarization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) SummarizationModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"summarizationModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) SystemPromptAisearch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"systemPromptAisearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) SystemPromptIndexSummarization() *string {
	var returns *string
	_jsii_.Get(
		j,
		"systemPromptIndexSummarization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) SystemPromptRewriteQuery() *string {
	var returns *string
	_jsii_.Get(
		j,
		"systemPromptRewriteQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) TokenId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) VectorizeActiveNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vectorizeActiveNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) VectorizeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vectorizeName",
		&returns,
	)
	return returns
}


func NewDataCloudflareAiSearchInstancesResultOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataCloudflareAiSearchInstancesResultOutputReference {
	_init_.Initialize()

	if err := validateNewDataCloudflareAiSearchInstancesResultOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-cloudflare.dataCloudflareAiSearchInstances.DataCloudflareAiSearchInstancesResultOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataCloudflareAiSearchInstancesResultOutputReference_Override(d DataCloudflareAiSearchInstancesResultOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-cloudflare.dataCloudflareAiSearchInstances.DataCloudflareAiSearchInstancesResultOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference)SetInternalValue(val *DataCloudflareAiSearchInstancesResult) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

