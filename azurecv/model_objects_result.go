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

// checks if the ObjectsResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectsResult{}

// ObjectsResult Describes detected objects in an image.
type ObjectsResult struct {
	// An array of detected objects.
	Values []DetectedObject `json:"values"`
}

type _ObjectsResult ObjectsResult

// NewObjectsResult instantiates a new ObjectsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectsResult(values []DetectedObject) *ObjectsResult {
	this := ObjectsResult{}
	this.Values = values
	return &this
}

// NewObjectsResultWithDefaults instantiates a new ObjectsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectsResultWithDefaults() *ObjectsResult {
	this := ObjectsResult{}
	return &this
}

// GetValues returns the Values field value
func (o *ObjectsResult) GetValues() []DetectedObject {
	if o == nil {
		var ret []DetectedObject
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *ObjectsResult) GetValuesOk() ([]DetectedObject, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *ObjectsResult) SetValues(v []DetectedObject) {
	o.Values = v
}

func (o ObjectsResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObjectsResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["values"] = o.Values
	return toSerialize, nil
}

func (o *ObjectsResult) UnmarshalJSON(data []byte) (err error) {
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

	varObjectsResult := _ObjectsResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varObjectsResult)

	if err != nil {
		return err
	}

	*o = ObjectsResult(varObjectsResult)

	return err
}

type NullableObjectsResult struct {
	value *ObjectsResult
	isSet bool
}

func (v NullableObjectsResult) Get() *ObjectsResult {
	return v.value
}

func (v *NullableObjectsResult) Set(val *ObjectsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectsResult(val *ObjectsResult) *NullableObjectsResult {
	return &NullableObjectsResult{value: val, isSet: true}
}

func (v NullableObjectsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


