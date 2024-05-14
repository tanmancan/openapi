# DetectedObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Id of the detected object. | [optional] 
**BoundingBox** | [**BoundingBox**](BoundingBox.md) |  | 
**Tags** | [**[]ContentTag**](ContentTag.md) | Classification confidences of the detected object. | 

## Methods

### NewDetectedObject

`func NewDetectedObject(boundingBox BoundingBox, tags []ContentTag, ) *DetectedObject`

NewDetectedObject instantiates a new DetectedObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetectedObjectWithDefaults

`func NewDetectedObjectWithDefaults() *DetectedObject`

NewDetectedObjectWithDefaults instantiates a new DetectedObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DetectedObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DetectedObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DetectedObject) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DetectedObject) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBoundingBox

`func (o *DetectedObject) GetBoundingBox() BoundingBox`

GetBoundingBox returns the BoundingBox field if non-nil, zero value otherwise.

### GetBoundingBoxOk

`func (o *DetectedObject) GetBoundingBoxOk() (*BoundingBox, bool)`

GetBoundingBoxOk returns a tuple with the BoundingBox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundingBox

`func (o *DetectedObject) SetBoundingBox(v BoundingBox)`

SetBoundingBox sets BoundingBox field to given value.


### GetTags

`func (o *DetectedObject) GetTags() []ContentTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DetectedObject) GetTagsOk() (*[]ContentTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DetectedObject) SetTags(v []ContentTag)`

SetTags sets Tags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


