# ErrorResponseDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Error code. | 
**Message** | **string** | Error message. | 
**Target** | Pointer to **string** | Target of the error. | [optional] 
**Details** | Pointer to [**[]ErrorResponseDetails**](ErrorResponseDetails.md) | List of detailed errors. | [optional] 
**Innererror** | Pointer to [**ErrorResponseInnerError**](ErrorResponseInnerError.md) |  | [optional] 

## Methods

### NewErrorResponseDetails

`func NewErrorResponseDetails(code string, message string, ) *ErrorResponseDetails`

NewErrorResponseDetails instantiates a new ErrorResponseDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorResponseDetailsWithDefaults

`func NewErrorResponseDetailsWithDefaults() *ErrorResponseDetails`

NewErrorResponseDetailsWithDefaults instantiates a new ErrorResponseDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ErrorResponseDetails) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorResponseDetails) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorResponseDetails) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *ErrorResponseDetails) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorResponseDetails) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorResponseDetails) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetTarget

`func (o *ErrorResponseDetails) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ErrorResponseDetails) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ErrorResponseDetails) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *ErrorResponseDetails) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetDetails

`func (o *ErrorResponseDetails) GetDetails() []ErrorResponseDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ErrorResponseDetails) GetDetailsOk() (*[]ErrorResponseDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ErrorResponseDetails) SetDetails(v []ErrorResponseDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ErrorResponseDetails) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetInnererror

`func (o *ErrorResponseDetails) GetInnererror() ErrorResponseInnerError`

GetInnererror returns the Innererror field if non-nil, zero value otherwise.

### GetInnererrorOk

`func (o *ErrorResponseDetails) GetInnererrorOk() (*ErrorResponseInnerError, bool)`

GetInnererrorOk returns a tuple with the Innererror field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnererror

`func (o *ErrorResponseDetails) SetInnererror(v ErrorResponseInnerError)`

SetInnererror sets Innererror field to given value.

### HasInnererror

`func (o *ErrorResponseDetails) HasInnererror() bool`

HasInnererror returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


