# CropRegion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AspectRatio** | **float64** | The aspect ratio of the crop region. | 
**BoundingBox** | [**BoundingBox**](BoundingBox.md) |  | 

## Methods

### NewCropRegion

`func NewCropRegion(aspectRatio float64, boundingBox BoundingBox, ) *CropRegion`

NewCropRegion instantiates a new CropRegion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCropRegionWithDefaults

`func NewCropRegionWithDefaults() *CropRegion`

NewCropRegionWithDefaults instantiates a new CropRegion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAspectRatio

`func (o *CropRegion) GetAspectRatio() float64`

GetAspectRatio returns the AspectRatio field if non-nil, zero value otherwise.

### GetAspectRatioOk

`func (o *CropRegion) GetAspectRatioOk() (*float64, bool)`

GetAspectRatioOk returns a tuple with the AspectRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspectRatio

`func (o *CropRegion) SetAspectRatio(v float64)`

SetAspectRatio sets AspectRatio field to given value.


### GetBoundingBox

`func (o *CropRegion) GetBoundingBox() BoundingBox`

GetBoundingBox returns the BoundingBox field if non-nil, zero value otherwise.

### GetBoundingBoxOk

`func (o *CropRegion) GetBoundingBoxOk() (*BoundingBox, bool)`

GetBoundingBoxOk returns a tuple with the BoundingBox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundingBox

`func (o *CropRegion) SetBoundingBox(v BoundingBox)`

SetBoundingBox sets BoundingBox field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


