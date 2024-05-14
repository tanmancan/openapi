# ContentTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the entity. | 
**Confidence** | **float64** | The level of confidence that the entity was observed. Confidence scores span the range of 0.0 to 1.0 (inclusive), with higher values indicating a higher confidence of a match. | 

## Methods

### NewContentTag

`func NewContentTag(name string, confidence float64, ) *ContentTag`

NewContentTag instantiates a new ContentTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentTagWithDefaults

`func NewContentTagWithDefaults() *ContentTag`

NewContentTagWithDefaults instantiates a new ContentTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ContentTag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContentTag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContentTag) SetName(v string)`

SetName sets Name field to given value.


### GetConfidence

`func (o *ContentTag) GetConfidence() float64`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *ContentTag) GetConfidenceOk() (*float64, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *ContentTag) SetConfidence(v float64)`

SetConfidence sets Confidence field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


