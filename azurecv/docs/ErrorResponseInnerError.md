# ErrorResponseInnerError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Error code. | 
**Message** | **string** | Error message. | 
**Innererror** | Pointer to [**ErrorResponseInnerError**](ErrorResponseInnerError.md) |  | [optional] 

## Methods

### NewErrorResponseInnerError

`func NewErrorResponseInnerError(code string, message string, ) *ErrorResponseInnerError`

NewErrorResponseInnerError instantiates a new ErrorResponseInnerError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorResponseInnerErrorWithDefaults

`func NewErrorResponseInnerErrorWithDefaults() *ErrorResponseInnerError`

NewErrorResponseInnerErrorWithDefaults instantiates a new ErrorResponseInnerError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ErrorResponseInnerError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorResponseInnerError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorResponseInnerError) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *ErrorResponseInnerError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorResponseInnerError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorResponseInnerError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetInnererror

`func (o *ErrorResponseInnerError) GetInnererror() ErrorResponseInnerError`

GetInnererror returns the Innererror field if non-nil, zero value otherwise.

### GetInnererrorOk

`func (o *ErrorResponseInnerError) GetInnererrorOk() (*ErrorResponseInnerError, bool)`

GetInnererrorOk returns a tuple with the Innererror field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnererror

`func (o *ErrorResponseInnerError) SetInnererror(v ErrorResponseInnerError)`

SetInnererror sets Innererror field to given value.

### HasInnererror

`func (o *ErrorResponseInnerError) HasInnererror() bool`

HasInnererror returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


