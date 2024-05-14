# BoundingBox

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**X** | **int32** | Left-coordinate of the top left point of the area, in pixels. | 
**Y** | **int32** | Top-coordinate of the top left point of the area, in pixels. | 
**W** | **int32** | Width measured from the top-left point of the area, in pixels. | 
**H** | **int32** | Height measured from the top-left point of the area, in pixels. | 

## Methods

### NewBoundingBox

`func NewBoundingBox(x int32, y int32, w int32, h int32, ) *BoundingBox`

NewBoundingBox instantiates a new BoundingBox object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoundingBoxWithDefaults

`func NewBoundingBoxWithDefaults() *BoundingBox`

NewBoundingBoxWithDefaults instantiates a new BoundingBox object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetX

`func (o *BoundingBox) GetX() int32`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *BoundingBox) GetXOk() (*int32, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *BoundingBox) SetX(v int32)`

SetX sets X field to given value.


### GetY

`func (o *BoundingBox) GetY() int32`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *BoundingBox) GetYOk() (*int32, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *BoundingBox) SetY(v int32)`

SetY sets Y field to given value.


### GetW

`func (o *BoundingBox) GetW() int32`

GetW returns the W field if non-nil, zero value otherwise.

### GetWOk

`func (o *BoundingBox) GetWOk() (*int32, bool)`

GetWOk returns a tuple with the W field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetW

`func (o *BoundingBox) SetW(v int32)`

SetW sets W field to given value.


### GetH

`func (o *BoundingBox) GetH() int32`

GetH returns the H field if non-nil, zero value otherwise.

### GetHOk

`func (o *BoundingBox) GetHOk() (*int32, bool)`

GetHOk returns a tuple with the H field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetH

`func (o *BoundingBox) SetH(v int32)`

SetH sets H field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


