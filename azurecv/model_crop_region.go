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

// checks if the CropRegion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CropRegion{}

// CropRegion A region identified for smart cropping. There will be one region returned for each requested aspect ratio.
type CropRegion struct {
	// The aspect ratio of the crop region.
	AspectRatio float64 `json:"aspectRatio"`
	BoundingBox BoundingBox `json:"boundingBox"`
}

type _CropRegion CropRegion

// NewCropRegion instantiates a new CropRegion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCropRegion(aspectRatio float64, boundingBox BoundingBox) *CropRegion {
	this := CropRegion{}
	this.AspectRatio = aspectRatio
	this.BoundingBox = boundingBox
	return &this
}

// NewCropRegionWithDefaults instantiates a new CropRegion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCropRegionWithDefaults() *CropRegion {
	this := CropRegion{}
	return &this
}

// GetAspectRatio returns the AspectRatio field value
func (o *CropRegion) GetAspectRatio() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.AspectRatio
}

// GetAspectRatioOk returns a tuple with the AspectRatio field value
// and a boolean to check if the value has been set.
func (o *CropRegion) GetAspectRatioOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AspectRatio, true
}

// SetAspectRatio sets field value
func (o *CropRegion) SetAspectRatio(v float64) {
	o.AspectRatio = v
}

// GetBoundingBox returns the BoundingBox field value
func (o *CropRegion) GetBoundingBox() BoundingBox {
	if o == nil {
		var ret BoundingBox
		return ret
	}

	return o.BoundingBox
}

// GetBoundingBoxOk returns a tuple with the BoundingBox field value
// and a boolean to check if the value has been set.
func (o *CropRegion) GetBoundingBoxOk() (*BoundingBox, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BoundingBox, true
}

// SetBoundingBox sets field value
func (o *CropRegion) SetBoundingBox(v BoundingBox) {
	o.BoundingBox = v
}

func (o CropRegion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CropRegion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["aspectRatio"] = o.AspectRatio
	toSerialize["boundingBox"] = o.BoundingBox
	return toSerialize, nil
}

func (o *CropRegion) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"aspectRatio",
		"boundingBox",
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

	varCropRegion := _CropRegion{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCropRegion)

	if err != nil {
		return err
	}

	*o = CropRegion(varCropRegion)

	return err
}

type NullableCropRegion struct {
	value *CropRegion
	isSet bool
}

func (v NullableCropRegion) Get() *CropRegion {
	return v.value
}

func (v *NullableCropRegion) Set(val *CropRegion) {
	v.value = val
	v.isSet = true
}

func (v NullableCropRegion) IsSet() bool {
	return v.isSet
}

func (v *NullableCropRegion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCropRegion(val *CropRegion) *NullableCropRegion {
	return &NullableCropRegion{value: val, isSet: true}
}

func (v NullableCropRegion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCropRegion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


