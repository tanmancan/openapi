# ReadResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blocks** | [**[]DetectedTextBlock**](DetectedTextBlock.md) | A list of text blocks. | 

## Methods

### NewReadResult

`func NewReadResult(blocks []DetectedTextBlock, ) *ReadResult`

NewReadResult instantiates a new ReadResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadResultWithDefaults

`func NewReadResultWithDefaults() *ReadResult`

NewReadResultWithDefaults instantiates a new ReadResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlocks

`func (o *ReadResult) GetBlocks() []DetectedTextBlock`

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

`func (o *ReadResult) GetBlocksOk() (*[]DetectedTextBlock, bool)`

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

`func (o *ReadResult) SetBlocks(v []DetectedTextBlock)`

SetBlocks sets Blocks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


