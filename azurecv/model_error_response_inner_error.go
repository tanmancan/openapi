/*
Computer Vision APIs (2024-02-01)

The Computer Vision API provides state-of-the-art algorithms to process images and return information. For example, it can be used to determine if an image contains mature content, or it can be used to find all the people in an image.  It also has other features like categorizing the content of images, and describing an image with complete English sentences.

API version: 2024-02-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package azurecv

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ErrorResponseInnerError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorResponseInnerError{}

// ErrorResponseInnerError Detailed error.
type ErrorResponseInnerError struct {
	// Error code.
	Code string `json:"code"`
	// Error message.
	Message string `json:"message"`
	Innererror *ErrorResponseInnerError `json:"innererror,omitempty"`
}

type _ErrorResponseInnerError ErrorResponseInnerError

// NewErrorResponseInnerError instantiates a new ErrorResponseInnerError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorResponseInnerError(code string, message string) *ErrorResponseInnerError {
	this := ErrorResponseInnerError{}
	this.Code = code
	this.Message = message
	return &this
}

// NewErrorResponseInnerErrorWithDefaults instantiates a new ErrorResponseInnerError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorResponseInnerErrorWithDefaults() *ErrorResponseInnerError {
	this := ErrorResponseInnerError{}
	return &this
}

// GetCode returns the Code field value
func (o *ErrorResponseInnerError) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *ErrorResponseInnerError) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *ErrorResponseInnerError) SetCode(v string) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *ErrorResponseInnerError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ErrorResponseInnerError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ErrorResponseInnerError) SetMessage(v string) {
	o.Message = v
}

// GetInnererror returns the Innererror field value if set, zero value otherwise.
func (o *ErrorResponseInnerError) GetInnererror() ErrorResponseInnerError {
	if o == nil || IsNil(o.Innererror) {
		var ret ErrorResponseInnerError
		return ret
	}
	return *o.Innererror
}

// GetInnererrorOk returns a tuple with the Innererror field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponseInnerError) GetInnererrorOk() (*ErrorResponseInnerError, bool) {
	if o == nil || IsNil(o.Innererror) {
		return nil, false
	}
	return o.Innererror, true
}

// HasInnererror returns a boolean if a field has been set.
func (o *ErrorResponseInnerError) HasInnererror() bool {
	if o != nil && !IsNil(o.Innererror) {
		return true
	}

	return false
}

// SetInnererror gets a reference to the given ErrorResponseInnerError and assigns it to the Innererror field.
func (o *ErrorResponseInnerError) SetInnererror(v ErrorResponseInnerError) {
	o.Innererror = &v
}

func (o ErrorResponseInnerError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorResponseInnerError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	if !IsNil(o.Innererror) {
		toSerialize["innererror"] = o.Innererror
	}
	return toSerialize, nil
}

func (o *ErrorResponseInnerError) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
		"message",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varErrorResponseInnerError := _ErrorResponseInnerError{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varErrorResponseInnerError)

	if err != nil {
		return err
	}

	*o = ErrorResponseInnerError(varErrorResponseInnerError)

	return err
}

type NullableErrorResponseInnerError struct {
	value *ErrorResponseInnerError
	isSet bool
}

func (v NullableErrorResponseInnerError) Get() *ErrorResponseInnerError {
	return v.value
}

func (v *NullableErrorResponseInnerError) Set(val *ErrorResponseInnerError) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorResponseInnerError) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorResponseInnerError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorResponseInnerError(val *ErrorResponseInnerError) *NullableErrorResponseInnerError {
	return &NullableErrorResponseInnerError{value: val, isSet: true}
}

func (v NullableErrorResponseInnerError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorResponseInnerError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


