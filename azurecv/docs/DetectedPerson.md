# DetectedPerson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BoundingBox** | [**BoundingBox**](BoundingBox.md) |  | 
**Confidence** | **float64** | Confidence score of having observed the person in the image. Confidence scores span the range of 0.0 to 1.0 (inclusive), with higher values indicating a higher confidence of a match. | 

## Methods

### NewDetectedPerson

`func NewDetectedPerson(boundingBox BoundingBox, confidence float64, ) *DetectedPerson`

NewDetectedPerson instantiates a new DetectedPerson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetectedPersonWithDefaults

`func NewDetectedPersonWithDefaults() *DetectedPerson`

NewDetectedPersonWithDefaults instantiates a new DetectedPerson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoundingBox

`func (o *DetectedPerson) GetBoundingBox() BoundingBox`

GetBoundingBox returns the BoundingBox field if non-nil, zero value otherwise.

### GetBoundingBoxOk

`func (o *DetectedPerson) GetBoundingBoxOk() (*BoundingBox, bool)`

GetBoundingBoxOk returns a tuple with the BoundingBox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundingBox

`func (o *DetectedPerson) SetBoundingBox(v BoundingBox)`

SetBoundingBox sets BoundingBox field to given value.


### GetConfidence

`func (o *DetectedPerson) GetConfidence() float64`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *DetectedPerson) GetConfidenceOk() (*float64, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *DetectedPerson) SetConfidence(v float64)`

SetConfidence sets Confidence field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


