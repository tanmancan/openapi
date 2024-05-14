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

// checks if the ImageMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageMetadata{}

// ImageMetadata The image metadata information such as height and width.
type ImageMetadata struct {
	// The width of the image in pixels.
	Width int32 `json:"width"`
	// The height of the image in pixels.
	Height int32 `json:"height"`
}

type _ImageMetadata ImageMetadata

// NewImageMetadata instantiates a new ImageMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageMetadata(width int32, height int32) *ImageMetadata {
	this := ImageMetadata{}
	this.Width = width
	this.Height = height
	return &this
}

// NewImageMetadataWithDefaults instantiates a new ImageMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageMetadataWithDefaults() *ImageMetadata {
	this := ImageMetadata{}
	return &this
}

// GetWidth returns the Width field value
func (o *ImageMetadata) GetWidth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Width
}

// GetWidthOk returns a tuple with the Width field value
// and a boolean to check if the value has been set.
func (o *ImageMetadata) GetWidthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Width, true
}

// SetWidth sets field value
func (o *ImageMetadata) SetWidth(v int32) {
	o.Width = v
}

// GetHeight returns the Height field value
func (o *ImageMetadata) GetHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Height
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
func (o *ImageMetadata) GetHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Height, true
}

// SetHeight sets field value
func (o *ImageMetadata) SetHeight(v int32) {
	o.Height = v
}

func (o ImageMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImageMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["width"] = o.Width
	toSerialize["height"] = o.Height
	return toSerialize, nil
}

func (o *ImageMetadata) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"width",
		"height",
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

	varImageMetadata := _ImageMetadata{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varImageMetadata)

	if err != nil {
		return err
	}

	*o = ImageMetadata(varImageMetadata)

	return err
}

type NullableImageMetadata struct {
	value *ImageMetadata
	isSet bool
}

func (v NullableImageMetadata) Get() *ImageMetadata {
	return v.value
}

func (v *NullableImageMetadata) Set(val *ImageMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableImageMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableImageMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageMetadata(val *ImageMetadata) *NullableImageMetadata {
	return &NullableImageMetadata{value: val, isSet: true}
}

func (v NullableImageMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


