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

// checks if the ImageAnalysisResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageAnalysisResult{}

// ImageAnalysisResult Describe the combined results of different types of image analysis.
type ImageAnalysisResult struct {
	// Model Version.
	ModelVersion string `json:"modelVersion"`
	CaptionResult *CaptionResult `json:"captionResult,omitempty"`
	DenseCaptionsResult *DenseCaptionsResult `json:"denseCaptionsResult,omitempty"`
	Metadata ImageMetadata `json:"metadata"`
	TagsResult *TagsResult `json:"tagsResult,omitempty"`
	ObjectsResult *ObjectsResult `json:"objectsResult,omitempty"`
	ReadResult *ReadResult `json:"readResult,omitempty"`
	SmartCropsResult *SmartCropsResult `json:"smartCropsResult,omitempty"`
	PeopleResult *PeopleResult `json:"peopleResult,omitempty"`
}

type _ImageAnalysisResult ImageAnalysisResult

// NewImageAnalysisResult instantiates a new ImageAnalysisResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageAnalysisResult(modelVersion string, metadata ImageMetadata) *ImageAnalysisResult {
	this := ImageAnalysisResult{}
	this.ModelVersion = modelVersion
	this.Metadata = metadata
	return &this
}

// NewImageAnalysisResultWithDefaults instantiates a new ImageAnalysisResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageAnalysisResultWithDefaults() *ImageAnalysisResult {
	this := ImageAnalysisResult{}
	return &this
}

// GetModelVersion returns the ModelVersion field value
func (o *ImageAnalysisResult) GetModelVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModelVersion
}

// GetModelVersionOk returns a tuple with the ModelVersion field value
// and a boolean to check if the value has been set.
func (o *ImageAnalysisResult) GetModelVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelVersion, true
}

// SetModelVersion sets field value
func (o *ImageAnalysisResult) SetModelVersion(v string) {
	o.ModelVersion = v
}

// GetCaptionResult returns the CaptionResult field value if set, zero value otherwise.
func (o *ImageAnalysisResult) GetCaptionResult() CaptionResult {
	if o == nil || IsNil(o.CaptionResult) {
		var ret CaptionResult
		return ret
	}
	return *o.CaptionResult
}

// GetCaptionResultOk returns a tuple with the CaptionResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageAnalysisResult) GetCaptionResultOk() (*CaptionResult, bool) {
	if o == nil || IsNil(o.CaptionResult) {
		return nil, false
	}
	return o.CaptionResult, true
}

// HasCaptionResult returns a boolean if a field has been set.
func (o *ImageAnalysisResult) HasCaptionResult() bool {
	if o != nil && !IsNil(o.CaptionResult) {
		return true
	}

	return false
}

// SetCaptionResult gets a reference to the given CaptionResult and assigns it to the CaptionResult field.
func (o *ImageAnalysisResult) SetCaptionResult(v CaptionResult) {
	o.CaptionResult = &v
}

// GetDenseCaptionsResult returns the DenseCaptionsResult field value if set, zero value otherwise.
func (o *ImageAnalysisResult) GetDenseCaptionsResult() DenseCaptionsResult {
	if o == nil || IsNil(o.DenseCaptionsResult) {
		var ret DenseCaptionsResult
		return ret
	}
	return *o.DenseCaptionsResult
}

// GetDenseCaptionsResultOk returns a tuple with the DenseCaptionsResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageAnalysisResult) GetDenseCaptionsResultOk() (*DenseCaptionsResult, bool) {
	if o == nil || IsNil(o.DenseCaptionsResult) {
		return nil, false
	}
	return o.DenseCaptionsResult, true
}

// HasDenseCaptionsResult returns a boolean if a field has been set.
func (o *ImageAnalysisResult) HasDenseCaptionsResult() bool {
	if o != nil && !IsNil(o.DenseCaptionsResult) {
		return true
	}

	return false
}

// SetDenseCaptionsResult gets a reference to the given DenseCaptionsResult and assigns it to the DenseCaptionsResult field.
func (o *ImageAnalysisResult) SetDenseCaptionsResult(v DenseCaptionsResult) {
	o.DenseCaptionsResult = &v
}

// GetMetadata returns the Metadata field value
func (o *ImageAnalysisResult) GetMetadata() ImageMetadata {
	if o == nil {
		var ret ImageMetadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *ImageAnalysisResult) GetMetadataOk() (*ImageMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *ImageAnalysisResult) SetMetadata(v ImageMetadata) {
	o.Metadata = v
}

// GetTagsResult returns the TagsResult field value if set, zero value otherwise.
func (o *ImageAnalysisResult) GetTagsResult() TagsResult {
	if o == nil || IsNil(o.TagsResult) {
		var ret TagsResult
		return ret
	}
	return *o.TagsResult
}

// GetTagsResultOk returns a tuple with the TagsResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageAnalysisResult) GetTagsResultOk() (*TagsResult, bool) {
	if o == nil || IsNil(o.TagsResult) {
		return nil, false
	}
	return o.TagsResult, true
}

// HasTagsResult returns a boolean if a field has been set.
func (o *ImageAnalysisResult) HasTagsResult() bool {
	if o != nil && !IsNil(o.TagsResult) {
		return true
	}

	return false
}

// SetTagsResult gets a reference to the given TagsResult and assigns it to the TagsResult field.
func (o *ImageAnalysisResult) SetTagsResult(v TagsResult) {
	o.TagsResult = &v
}

// GetObjectsResult returns the ObjectsResult field value if set, zero value otherwise.
func (o *ImageAnalysisResult) GetObjectsResult() ObjectsResult {
	if o == nil || IsNil(o.ObjectsResult) {
		var ret ObjectsResult
		return ret
	}
	return *o.ObjectsResult
}

// GetObjectsResultOk returns a tuple with the ObjectsResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageAnalysisResult) GetObjectsResultOk() (*ObjectsResult, bool) {
	if o == nil || IsNil(o.ObjectsResult) {
		return nil, false
	}
	return o.ObjectsResult, true
}

// HasObjectsResult returns a boolean if a field has been set.
func (o *ImageAnalysisResult) HasObjectsResult() bool {
	if o != nil && !IsNil(o.ObjectsResult) {
		return true
	}

	return false
}

// SetObjectsResult gets a reference to the given ObjectsResult and assigns it to the ObjectsResult field.
func (o *ImageAnalysisResult) SetObjectsResult(v ObjectsResult) {
	o.ObjectsResult = &v
}

// GetReadResult returns the ReadResult field value if set, zero value otherwise.
func (o *ImageAnalysisResult) GetReadResult() ReadResult {
	if o == nil || IsNil(o.ReadResult) {
		var ret ReadResult
		return ret
	}
	return *o.ReadResult
}

// GetReadResultOk returns a tuple with the ReadResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageAnalysisResult) GetReadResultOk() (*ReadResult, bool) {
	if o == nil || IsNil(o.ReadResult) {
		return nil, false
	}
	return o.ReadResult, true
}

// HasReadResult returns a boolean if a field has been set.
func (o *ImageAnalysisResult) HasReadResult() bool {
	if o != nil && !IsNil(o.ReadResult) {
		return true
	}

	return false
}

// SetReadResult gets a reference to the given ReadResult and assigns it to the ReadResult field.
func (o *ImageAnalysisResult) SetReadResult(v ReadResult) {
	o.ReadResult = &v
}

// GetSmartCropsResult returns the SmartCropsResult field value if set, zero value otherwise.
func (o *ImageAnalysisResult) GetSmartCropsResult() SmartCropsResult {
	if o == nil || IsNil(o.SmartCropsResult) {
		var ret SmartCropsResult
		return ret
	}
	return *o.SmartCropsResult
}

// GetSmartCropsResultOk returns a tuple with the SmartCropsResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageAnalysisResult) GetSmartCropsResultOk() (*SmartCropsResult, bool) {
	if o == nil || IsNil(o.SmartCropsResult) {
		return nil, false
	}
	return o.SmartCropsResult, true
}

// HasSmartCropsResult returns a boolean if a field has been set.
func (o *ImageAnalysisResult) HasSmartCropsResult() bool {
	if o != nil && !IsNil(o.SmartCropsResult) {
		return true
	}

	return false
}

// SetSmartCropsResult gets a reference to the given SmartCropsResult and assigns it to the SmartCropsResult field.
func (o *ImageAnalysisResult) SetSmartCropsResult(v SmartCropsResult) {
	o.SmartCropsResult = &v
}

// GetPeopleResult returns the PeopleResult field value if set, zero value otherwise.
func (o *ImageAnalysisResult) GetPeopleResult() PeopleResult {
	if o == nil || IsNil(o.PeopleResult) {
		var ret PeopleResult
		return ret
	}
	return *o.PeopleResult
}

// GetPeopleResultOk returns a tuple with the PeopleResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageAnalysisResult) GetPeopleResultOk() (*PeopleResult, bool) {
	if o == nil || IsNil(o.PeopleResult) {
		return nil, false
	}
	return o.PeopleResult, true
}

// HasPeopleResult returns a boolean if a field has been set.
func (o *ImageAnalysisResult) HasPeopleResult() bool {
	if o != nil && !IsNil(o.PeopleResult) {
		return true
	}

	return false
}

// SetPeopleResult gets a reference to the given PeopleResult and assigns it to the PeopleResult field.
func (o *ImageAnalysisResult) SetPeopleResult(v PeopleResult) {
	o.PeopleResult = &v
}

func (o ImageAnalysisResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImageAnalysisResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["modelVersion"] = o.ModelVersion
	if !IsNil(o.CaptionResult) {
		toSerialize["captionResult"] = o.CaptionResult
	}
	if !IsNil(o.DenseCaptionsResult) {
		toSerialize["denseCaptionsResult"] = o.DenseCaptionsResult
	}
	toSerialize["metadata"] = o.Metadata
	if !IsNil(o.TagsResult) {
		toSerialize["tagsResult"] = o.TagsResult
	}
	if !IsNil(o.ObjectsResult) {
		toSerialize["objectsResult"] = o.ObjectsResult
	}
	if !IsNil(o.ReadResult) {
		toSerialize["readResult"] = o.ReadResult
	}
	if !IsNil(o.SmartCropsResult) {
		toSerialize["smartCropsResult"] = o.SmartCropsResult
	}
	if !IsNil(o.PeopleResult) {
		toSerialize["peopleResult"] = o.PeopleResult
	}
	return toSerialize, nil
}

func (o *ImageAnalysisResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"modelVersion",
		"metadata",
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

	varImageAnalysisResult := _ImageAnalysisResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varImageAnalysisResult)

	if err != nil {
		return err
	}

	*o = ImageAnalysisResult(varImageAnalysisResult)

	return err
}

type NullableImageAnalysisResult struct {
	value *ImageAnalysisResult
	isSet bool
}

func (v NullableImageAnalysisResult) Get() *ImageAnalysisResult {
	return v.value
}

func (v *NullableImageAnalysisResult) Set(val *ImageAnalysisResult) {
	v.value = val
	v.isSet = true
}

func (v NullableImageAnalysisResult) IsSet() bool {
	return v.isSet
}

func (v *NullableImageAnalysisResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageAnalysisResult(val *ImageAnalysisResult) *NullableImageAnalysisResult {
	return &NullableImageAnalysisResult{value: val, isSet: true}
}

func (v NullableImageAnalysisResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageAnalysisResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


