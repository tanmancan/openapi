# ImageAnalysisResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModelVersion** | **string** | Model Version. | 
**CaptionResult** | Pointer to [**CaptionResult**](CaptionResult.md) |  | [optional] 
**DenseCaptionsResult** | Pointer to [**DenseCaptionsResult**](DenseCaptionsResult.md) |  | [optional] 
**Metadata** | [**ImageMetadata**](ImageMetadata.md) |  | 
**TagsResult** | Pointer to [**TagsResult**](TagsResult.md) |  | [optional] 
**ObjectsResult** | Pointer to [**ObjectsResult**](ObjectsResult.md) |  | [optional] 
**ReadResult** | Pointer to [**ReadResult**](ReadResult.md) |  | [optional] 
**SmartCropsResult** | Pointer to [**SmartCropsResult**](SmartCropsResult.md) |  | [optional] 
**PeopleResult** | Pointer to [**PeopleResult**](PeopleResult.md) |  | [optional] 

## Methods

### NewImageAnalysisResult

`func NewImageAnalysisResult(modelVersion string, metadata ImageMetadata, ) *ImageAnalysisResult`

NewImageAnalysisResult instantiates a new ImageAnalysisResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageAnalysisResultWithDefaults

`func NewImageAnalysisResultWithDefaults() *ImageAnalysisResult`

NewImageAnalysisResultWithDefaults instantiates a new ImageAnalysisResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModelVersion

`func (o *ImageAnalysisResult) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *ImageAnalysisResult) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *ImageAnalysisResult) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.


### GetCaptionResult

`func (o *ImageAnalysisResult) GetCaptionResult() CaptionResult`

GetCaptionResult returns the CaptionResult field if non-nil, zero value otherwise.

### GetCaptionResultOk

`func (o *ImageAnalysisResult) GetCaptionResultOk() (*CaptionResult, bool)`

GetCaptionResultOk returns a tuple with the CaptionResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptionResult

`func (o *ImageAnalysisResult) SetCaptionResult(v CaptionResult)`

SetCaptionResult sets CaptionResult field to given value.

### HasCaptionResult

`func (o *ImageAnalysisResult) HasCaptionResult() bool`

HasCaptionResult returns a boolean if a field has been set.

### GetDenseCaptionsResult

`func (o *ImageAnalysisResult) GetDenseCaptionsResult() DenseCaptionsResult`

GetDenseCaptionsResult returns the DenseCaptionsResult field if non-nil, zero value otherwise.

### GetDenseCaptionsResultOk

`func (o *ImageAnalysisResult) GetDenseCaptionsResultOk() (*DenseCaptionsResult, bool)`

GetDenseCaptionsResultOk returns a tuple with the DenseCaptionsResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenseCaptionsResult

`func (o *ImageAnalysisResult) SetDenseCaptionsResult(v DenseCaptionsResult)`

SetDenseCaptionsResult sets DenseCaptionsResult field to given value.

### HasDenseCaptionsResult

`func (o *ImageAnalysisResult) HasDenseCaptionsResult() bool`

HasDenseCaptionsResult returns a boolean if a field has been set.

### GetMetadata

`func (o *ImageAnalysisResult) GetMetadata() ImageMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ImageAnalysisResult) GetMetadataOk() (*ImageMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ImageAnalysisResult) SetMetadata(v ImageMetadata)`

SetMetadata sets Metadata field to given value.


### GetTagsResult

`func (o *ImageAnalysisResult) GetTagsResult() TagsResult`

GetTagsResult returns the TagsResult field if non-nil, zero value otherwise.

### GetTagsResultOk

`func (o *ImageAnalysisResult) GetTagsResultOk() (*TagsResult, bool)`

GetTagsResultOk returns a tuple with the TagsResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagsResult

`func (o *ImageAnalysisResult) SetTagsResult(v TagsResult)`

SetTagsResult sets TagsResult field to given value.

### HasTagsResult

`func (o *ImageAnalysisResult) HasTagsResult() bool`

HasTagsResult returns a boolean if a field has been set.

### GetObjectsResult

`func (o *ImageAnalysisResult) GetObjectsResult() ObjectsResult`

GetObjectsResult returns the ObjectsResult field if non-nil, zero value otherwise.

### GetObjectsResultOk

`func (o *ImageAnalysisResult) GetObjectsResultOk() (*ObjectsResult, bool)`

GetObjectsResultOk returns a tuple with the ObjectsResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectsResult

`func (o *ImageAnalysisResult) SetObjectsResult(v ObjectsResult)`

SetObjectsResult sets ObjectsResult field to given value.

### HasObjectsResult

`func (o *ImageAnalysisResult) HasObjectsResult() bool`

HasObjectsResult returns a boolean if a field has been set.

### GetReadResult

`func (o *ImageAnalysisResult) GetReadResult() ReadResult`

GetReadResult returns the ReadResult field if non-nil, zero value otherwise.

### GetReadResultOk

`func (o *ImageAnalysisResult) GetReadResultOk() (*ReadResult, bool)`

GetReadResultOk returns a tuple with the ReadResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadResult

`func (o *ImageAnalysisResult) SetReadResult(v ReadResult)`

SetReadResult sets ReadResult field to given value.

### HasReadResult

`func (o *ImageAnalysisResult) HasReadResult() bool`

HasReadResult returns a boolean if a field has been set.

### GetSmartCropsResult

`func (o *ImageAnalysisResult) GetSmartCropsResult() SmartCropsResult`

GetSmartCropsResult returns the SmartCropsResult field if non-nil, zero value otherwise.

### GetSmartCropsResultOk

`func (o *ImageAnalysisResult) GetSmartCropsResultOk() (*SmartCropsResult, bool)`

GetSmartCropsResultOk returns a tuple with the SmartCropsResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartCropsResult

`func (o *ImageAnalysisResult) SetSmartCropsResult(v SmartCropsResult)`

SetSmartCropsResult sets SmartCropsResult field to given value.

### HasSmartCropsResult

`func (o *ImageAnalysisResult) HasSmartCropsResult() bool`

HasSmartCropsResult returns a boolean if a field has been set.

### GetPeopleResult

`func (o *ImageAnalysisResult) GetPeopleResult() PeopleResult`

GetPeopleResult returns the PeopleResult field if non-nil, zero value otherwise.

### GetPeopleResultOk

`func (o *ImageAnalysisResult) GetPeopleResultOk() (*PeopleResult, bool)`

GetPeopleResultOk returns a tuple with the PeopleResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeopleResult

`func (o *ImageAnalysisResult) SetPeopleResult(v PeopleResult)`

SetPeopleResult sets PeopleResult field to given value.

### HasPeopleResult

`func (o *ImageAnalysisResult) HasPeopleResult() bool`

HasPeopleResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


