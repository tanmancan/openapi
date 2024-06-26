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

// checks if the ErrorResponseDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorResponseDetails{}

// ErrorResponseDetails Error info.
type ErrorResponseDetails struct {
	// Error code.
	Code string `json:"code"`
	// Error message.
	Message string `json:"message"`
	// Target of the error.
	Target *string `json:"target,omitempty"`
	// List of detailed errors.
	Details []ErrorResponseDetails `json:"details,omitempty"`
	Innererror *ErrorResponseInnerError `json:"innererror,omitempty"`
}

type _ErrorResponseDetails ErrorResponseDetails

// NewErrorResponseDetails instantiates a new ErrorResponseDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorResponseDetails(code string, message string) *ErrorResponseDetails {
	this := ErrorResponseDetails{}
	this.Code = code
	this.Message = message
	return &this
}

// NewErrorResponseDetailsWithDefaults instantiates a new ErrorResponseDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorResponseDetailsWithDefaults() *ErrorResponseDetails {
	this := ErrorResponseDetails{}
	return &this
}

// GetCode returns the Code field value
func (o *ErrorResponseDetails) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *ErrorResponseDetails) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *ErrorResponseDetails) SetCode(v string) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *ErrorResponseDetails) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ErrorResponseDetails) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ErrorResponseDetails) SetMessage(v string) {
	o.Message = v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *ErrorResponseDetails) GetTarget() string {
	if o == nil || IsNil(o.Target) {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponseDetails) GetTargetOk() (*string, bool) {
	if o == nil || IsNil(o.Target) {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *ErrorResponseDetails) HasTarget() bool {
	if o != nil && !IsNil(o.Target) {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *ErrorResponseDetails) SetTarget(v string) {
	o.Target = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *ErrorResponseDetails) GetDetails() []ErrorResponseDetails {
	if o == nil || IsNil(o.Details) {
		var ret []ErrorResponseDetails
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponseDetails) GetDetailsOk() ([]ErrorResponseDetails, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *ErrorResponseDetails) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []ErrorResponseDetails and assigns it to the Details field.
func (o *ErrorResponseDetails) SetDetails(v []ErrorResponseDetails) {
	o.Details = v
}

// GetInnererror returns the Innererror field value if set, zero value otherwise.
func (o *ErrorResponseDetails) GetInnererror() ErrorResponseInnerError {
	if o == nil || IsNil(o.Innererror) {
		var ret ErrorResponseInnerError
		return ret
	}
	return *o.Innererror
}

// GetInnererrorOk returns a tuple with the Innererror field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponseDetails) GetInnererrorOk() (*ErrorResponseInnerError, bool) {
	if o == nil || IsNil(o.Innererror) {
		return nil, false
	}
	return o.Innererror, true
}

// HasInnererror returns a boolean if a field has been set.
func (o *ErrorResponseDetails) HasInnererror() bool {
	if o != nil && !IsNil(o.Innererror) {
		return true
	}

	return false
}

// SetInnererror gets a reference to the given ErrorResponseInnerError and assigns it to the Innererror field.
func (o *ErrorResponseDetails) SetInnererror(v ErrorResponseInnerError) {
	o.Innererror = &v
}

func (o ErrorResponseDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorResponseDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	if !IsNil(o.Target) {
		toSerialize["target"] = o.Target
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !IsNil(o.Innererror) {
		toSerialize["innererror"] = o.Innererror
	}
	return toSerialize, nil
}

func (o *ErrorResponseDetails) UnmarshalJSON(data []byte) (err error) {
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

	varErrorResponseDetails := _ErrorResponseDetails{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varErrorResponseDetails)

	if err != nil {
		return err
	}

	*o = ErrorResponseDetails(varErrorResponseDetails)

	return err
}

type NullableErrorResponseDetails struct {
	value *ErrorResponseDetails
	isSet bool
}

func (v NullableErrorResponseDetails) Get() *ErrorResponseDetails {
	return v.value
}

func (v *NullableErrorResponseDetails) Set(val *ErrorResponseDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorResponseDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorResponseDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorResponseDetails(val *ErrorResponseDetails) *NullableErrorResponseDetails {
	return &NullableErrorResponseDetails{value: val, isSet: true}
}

func (v NullableErrorResponseDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorResponseDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


