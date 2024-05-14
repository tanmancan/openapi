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

// checks if the DetectedPerson type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DetectedPerson{}

// DetectedPerson A person detected in an image.
type DetectedPerson struct {
	BoundingBox BoundingBox `json:"boundingBox"`
	// Confidence score of having observed the person in the image. Confidence scores span the range of 0.0 to 1.0 (inclusive), with higher values indicating a higher confidence of a match.
	Confidence float64 `json:"confidence"`
}

type _DetectedPerson DetectedPerson

// NewDetectedPerson instantiates a new DetectedPerson object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDetectedPerson(boundingBox BoundingBox, confidence float64) *DetectedPerson {
	this := DetectedPerson{}
	this.BoundingBox = boundingBox
	this.Confidence = confidence
	return &this
}

// NewDetectedPersonWithDefaults instantiates a new DetectedPerson object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDetectedPersonWithDefaults() *DetectedPerson {
	this := DetectedPerson{}
	return &this
}

// GetBoundingBox returns the BoundingBox field value
func (o *DetectedPerson) GetBoundingBox() BoundingBox {
	if o == nil {
		var ret BoundingBox
		return ret
	}

	return o.BoundingBox
}

// GetBoundingBoxOk returns a tuple with the BoundingBox field value
// and a boolean to check if the value has been set.
func (o *DetectedPerson) GetBoundingBoxOk() (*BoundingBox, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BoundingBox, true
}

// SetBoundingBox sets field value
func (o *DetectedPerson) SetBoundingBox(v BoundingBox) {
	o.BoundingBox = v
}

// GetConfidence returns the Confidence field value
func (o *DetectedPerson) GetConfidence() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value
// and a boolean to check if the value has been set.
func (o *DetectedPerson) GetConfidenceOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Confidence, true
}

// SetConfidence sets field value
func (o *DetectedPerson) SetConfidence(v float64) {
	o.Confidence = v
}

func (o DetectedPerson) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DetectedPerson) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["boundingBox"] = o.BoundingBox
	toSerialize["confidence"] = o.Confidence
	return toSerialize, nil
}

func (o *DetectedPerson) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"boundingBox",
		"confidence",
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

	varDetectedPerson := _DetectedPerson{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDetectedPerson)

	if err != nil {
		return err
	}

	*o = DetectedPerson(varDetectedPerson)

	return err
}

type NullableDetectedPerson struct {
	value *DetectedPerson
	isSet bool
}

func (v NullableDetectedPerson) Get() *DetectedPerson {
	return v.value
}

func (v *NullableDetectedPerson) Set(val *DetectedPerson) {
	v.value = val
	v.isSet = true
}

func (v NullableDetectedPerson) IsSet() bool {
	return v.isSet
}

func (v *NullableDetectedPerson) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDetectedPerson(val *DetectedPerson) *NullableDetectedPerson {
	return &NullableDetectedPerson{value: val, isSet: true}
}

func (v NullableDetectedPerson) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDetectedPerson) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


