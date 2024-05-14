# SingleVectorResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vector** | Pointer to **[]float32** | Vector of the image. | [optional] 
**ModelVersion** | Pointer to **string** | Model version. | [optional] 

## Methods

### NewSingleVectorResult

`func NewSingleVectorResult() *SingleVectorResult`

NewSingleVectorResult instantiates a new SingleVectorResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSingleVectorResultWithDefaults

`func NewSingleVectorResultWithDefaults() *SingleVectorResult`

NewSingleVectorResultWithDefaults instantiates a new SingleVectorResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVector

`func (o *SingleVectorResult) GetVector() []float32`

GetVector returns the Vector field if non-nil, zero value otherwise.

### GetVectorOk

`func (o *SingleVectorResult) GetVectorOk() (*[]float32, bool)`

GetVectorOk returns a tuple with the Vector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVector

`func (o *SingleVectorResult) SetVector(v []float32)`

SetVector sets Vector field to given value.

### HasVector

`func (o *SingleVectorResult) HasVector() bool`

HasVector returns a boolean if a field has been set.

### GetModelVersion

`func (o *SingleVectorResult) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *SingleVectorResult) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *SingleVectorResult) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *SingleVectorResult) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


