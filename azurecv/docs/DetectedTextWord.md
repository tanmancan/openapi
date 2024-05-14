# DetectedTextWord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** | Text content of the word. | 
**BoundingPolygon** | [**[]ImagePoint**](ImagePoint.md) | Bounding polygon of the word. | 
**Confidence** | **float64** | The level of confidence that the word was detected. Confidence scores span the range of 0.0 to 1.0 (inclusive), with higher values indicating a higher confidence of a match. | 

## Methods

### NewDetectedTextWord

`func NewDetectedTextWord(text string, boundingPolygon []ImagePoint, confidence float64, ) *DetectedTextWord`

NewDetectedTextWord instantiates a new DetectedTextWord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetectedTextWordWithDefaults

`func NewDetectedTextWordWithDefaults() *DetectedTextWord`

NewDetectedTextWordWithDefaults instantiates a new DetectedTextWord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *DetectedTextWord) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *DetectedTextWord) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *DetectedTextWord) SetText(v string)`

SetText sets Text field to given value.


### GetBoundingPolygon

`func (o *DetectedTextWord) GetBoundingPolygon() []ImagePoint`

GetBoundingPolygon returns the BoundingPolygon field if non-nil, zero value otherwise.

### GetBoundingPolygonOk

`func (o *DetectedTextWord) GetBoundingPolygonOk() (*[]ImagePoint, bool)`

GetBoundingPolygonOk returns a tuple with the BoundingPolygon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundingPolygon

`func (o *DetectedTextWord) SetBoundingPolygon(v []ImagePoint)`

SetBoundingPolygon sets BoundingPolygon field to given value.


### GetConfidence

`func (o *DetectedTextWord) GetConfidence() float64`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *DetectedTextWord) GetConfidenceOk() (*float64, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *DetectedTextWord) SetConfidence(v float64)`

SetConfidence sets Confidence field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


