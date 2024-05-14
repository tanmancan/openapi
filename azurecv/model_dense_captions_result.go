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

// checks if the DenseCaptionsResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DenseCaptionsResult{}

// DenseCaptionsResult A list of captions.
type DenseCaptionsResult struct {
	// A list of captions.
	Values []DenseCaption `json:"values"`
}

type _DenseCaptionsResult DenseCaptionsResult

// NewDenseCaptionsResult instantiates a new DenseCaptionsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDenseCaptionsResult(values []DenseCaption) *DenseCaptionsResult {
	this := DenseCaptionsResult{}
	this.Values = values
	return &this
}

// NewDenseCaptionsResultWithDefaults instantiates a new DenseCaptionsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDenseCaptionsResultWithDefaults() *DenseCaptionsResult {
	this := DenseCaptionsResult{}
	return &this
}

// GetValues returns the Values field value
func (o *DenseCaptionsResult) GetValues() []DenseCaption {
	if o == nil {
		var ret []DenseCaption
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *DenseCaptionsResult) GetValuesOk() ([]DenseCaption, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *DenseCaptionsResult) SetValues(v []DenseCaption) {
	o.Values = v
}

func (o DenseCaptionsResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DenseCaptionsResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["values"] = o.Values
	return toSerialize, nil
}

func (o *DenseCaptionsResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"values",
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

	varDenseCaptionsResult := _DenseCaptionsResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDenseCaptionsResult)

	if err != nil {
		return err
	}

	*o = DenseCaptionsResult(varDenseCaptionsResult)

	return err
}

type NullableDenseCaptionsResult struct {
	value *DenseCaptionsResult
	isSet bool
}

func (v NullableDenseCaptionsResult) Get() *DenseCaptionsResult {
	return v.value
}

func (v *NullableDenseCaptionsResult) Set(val *DenseCaptionsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableDenseCaptionsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableDenseCaptionsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDenseCaptionsResult(val *DenseCaptionsResult) *NullableDenseCaptionsResult {
	return &NullableDenseCaptionsResult{value: val, isSet: true}
}

func (v NullableDenseCaptionsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDenseCaptionsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


