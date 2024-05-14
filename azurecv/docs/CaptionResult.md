# CaptionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** | The text of the caption. | 
**Confidence** | **float64** | The level of confidence the service has in the caption. Confidence scores span the range of 0.0 to 1.0 (inclusive), with higher values indicating a higher confidence of a match. | 

## Methods

### NewCaptionResult

`func NewCaptionResult(text string, confidence float64, ) *CaptionResult`

NewCaptionResult instantiates a new CaptionResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptionResultWithDefaults

`func NewCaptionResultWithDefaults() *CaptionResult`

NewCaptionResultWithDefaults instantiates a new CaptionResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *CaptionResult) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CaptionResult) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CaptionResult) SetText(v string)`

SetText sets Text field to given value.


### GetConfidence

`func (o *CaptionResult) GetConfidence() float64`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *CaptionResult) GetConfidenceOk() (*float64, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *CaptionResult) SetConfidence(v float64)`

SetConfidence sets Confidence field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


