# DenseCaption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** | The text of the caption. | 
**Confidence** | **float64** | The level of confidence the service has in the caption. Confidence scores span the range of 0.0 to 1.0 (inclusive), with higher values indicating a higher confidence of a match. | 
**BoundingBox** | Pointer to [**BoundingBox**](BoundingBox.md) |  | [optional] 

## Methods

### NewDenseCaption

`func NewDenseCaption(text string, confidence float64, ) *DenseCaption`

NewDenseCaption instantiates a new DenseCaption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDenseCaptionWithDefaults

`func NewDenseCaptionWithDefaults() *DenseCaption`

NewDenseCaptionWithDefaults instantiates a new DenseCaption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *DenseCaption) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *DenseCaption) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *DenseCaption) SetText(v string)`

SetText sets Text field to given value.


### GetConfidence

`func (o *DenseCaption) GetConfidence() float64`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *DenseCaption) GetConfidenceOk() (*float64, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *DenseCaption) SetConfidence(v float64)`

SetConfidence sets Confidence field to given value.


### GetBoundingBox

`func (o *DenseCaption) GetBoundingBox() BoundingBox`

GetBoundingBox returns the BoundingBox field if non-nil, zero value otherwise.

### GetBoundingBoxOk

`func (o *DenseCaption) GetBoundingBoxOk() (*BoundingBox, bool)`

GetBoundingBoxOk returns a tuple with the BoundingBox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundingBox

`func (o *DenseCaption) SetBoundingBox(v BoundingBox)`

SetBoundingBox sets BoundingBox field to given value.

### HasBoundingBox

`func (o *DenseCaption) HasBoundingBox() bool`

HasBoundingBox returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


