# DetectedTextLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** | Text content of the detected text line. | 
**BoundingPolygon** | [**[]ImagePoint**](ImagePoint.md) | Bounding polygon of the text line. | 
**Words** | [**[]DetectedTextWord**](DetectedTextWord.md) | List of words in the text line. | 

## Methods

### NewDetectedTextLine

`func NewDetectedTextLine(text string, boundingPolygon []ImagePoint, words []DetectedTextWord, ) *DetectedTextLine`

NewDetectedTextLine instantiates a new DetectedTextLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetectedTextLineWithDefaults

`func NewDetectedTextLineWithDefaults() *DetectedTextLine`

NewDetectedTextLineWithDefaults instantiates a new DetectedTextLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *DetectedTextLine) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *DetectedTextLine) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *DetectedTextLine) SetText(v string)`

SetText sets Text field to given value.


### GetBoundingPolygon

`func (o *DetectedTextLine) GetBoundingPolygon() []ImagePoint`

GetBoundingPolygon returns the BoundingPolygon field if non-nil, zero value otherwise.

### GetBoundingPolygonOk

`func (o *DetectedTextLine) GetBoundingPolygonOk() (*[]ImagePoint, bool)`

GetBoundingPolygonOk returns a tuple with the BoundingPolygon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundingPolygon

`func (o *DetectedTextLine) SetBoundingPolygon(v []ImagePoint)`

SetBoundingPolygon sets BoundingPolygon field to given value.


### GetWords

`func (o *DetectedTextLine) GetWords() []DetectedTextWord`

GetWords returns the Words field if non-nil, zero value otherwise.

### GetWordsOk

`func (o *DetectedTextLine) GetWordsOk() (*[]DetectedTextWord, bool)`

GetWordsOk returns a tuple with the Words field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWords

`func (o *DetectedTextLine) SetWords(v []DetectedTextWord)`

SetWords sets Words field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


