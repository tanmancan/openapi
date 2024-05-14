# \DefaultAPI

All URIs are relative to *https://cognitiveuseprod.cognitiveservices.azure.com/computervision*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ImageAnalysisAnalyze**](DefaultAPI.md#ImageAnalysisAnalyze) | **Post** /imageanalysis:analyze | Analyze
[**RetrievalVectorizeImage**](DefaultAPI.md#RetrievalVectorizeImage) | **Post** /retrieval:vectorizeImage | VectorizeImage
[**RetrievalVectorizeText**](DefaultAPI.md#RetrievalVectorizeText) | **Post** /retrieval:vectorizeText | VectorizeText



## ImageAnalysisAnalyze

> ImageAnalysisResult ImageAnalysisAnalyze(ctx).ApiVersion(apiVersion).Features(features).ModelVersion(modelVersion).Language(language).SmartcropsAspectRatios(smartcropsAspectRatios).GenderNeutralCaption(genderNeutralCaption).Body(body).Execute()

Analyze



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tanmancan/openapi/azurecv"
)

func main() {
	apiVersion := "apiVersion_example" // string |  (optional) (default to "2024-02-01")
	features := []string{"Inner_example"} // []string | The visual features requested: tags, objects, caption, denseCaptions, read, smartCrops, people. At least one visual feature must be specified. (optional)
	modelVersion := "modelVersion_example" // string | Optional parameter to specify the version of the AI model, or \"latest\" to use the latest available model. Defaults to \"latest\". (optional) (default to "latest")
	language := "language_example" // string | The desired language for output generation. If this parameter is not specified, the default value is \"en\". See https://aka.ms/cv-languages for a list of supported languages. (optional) (default to "en")
	smartcropsAspectRatios := []string{"Inner_example"} // []string | A list of aspect ratios to use for smartCrops feature. Aspect ratios are calculated by dividing the target crop width by the height. Supported values are between 0.75 and 1.8 (inclusive). Multiple values should be comma-separated. If this parameter is not specified, the service will return one crop suggestion with an aspect ratio it sees fit between 0.5 and 2.0 (inclusive). (optional)
	genderNeutralCaption := true // bool | Boolean flag for enabling gender-neutral captioning for caption and denseCaptions features. If this parameter is not specified, the default value is \"false\". (optional) (default to false)
	body := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ImageAnalysisAnalyze(context.Background()).ApiVersion(apiVersion).Features(features).ModelVersion(modelVersion).Language(language).SmartcropsAspectRatios(smartcropsAspectRatios).GenderNeutralCaption(genderNeutralCaption).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ImageAnalysisAnalyze``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImageAnalysisAnalyze`: ImageAnalysisResult
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ImageAnalysisAnalyze`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImageAnalysisAnalyzeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | [default to &quot;2024-02-01&quot;]
 **features** | **[]string** | The visual features requested: tags, objects, caption, denseCaptions, read, smartCrops, people. At least one visual feature must be specified. | 
 **modelVersion** | **string** | Optional parameter to specify the version of the AI model, or \&quot;latest\&quot; to use the latest available model. Defaults to \&quot;latest\&quot;. | [default to &quot;latest&quot;]
 **language** | **string** | The desired language for output generation. If this parameter is not specified, the default value is \&quot;en\&quot;. See https://aka.ms/cv-languages for a list of supported languages. | [default to &quot;en&quot;]
 **smartcropsAspectRatios** | **[]string** | A list of aspect ratios to use for smartCrops feature. Aspect ratios are calculated by dividing the target crop width by the height. Supported values are between 0.75 and 1.8 (inclusive). Multiple values should be comma-separated. If this parameter is not specified, the service will return one crop suggestion with an aspect ratio it sees fit between 0.5 and 2.0 (inclusive). | 
 **genderNeutralCaption** | **bool** | Boolean flag for enabling gender-neutral captioning for caption and denseCaptions features. If this parameter is not specified, the default value is \&quot;false\&quot;. | [default to false]
 **body** | ***os.File** |  | 

### Return type

[**ImageAnalysisResult**](ImageAnalysisResult.md)

### Authorization

[apiKeyQuery](../README.md#apiKeyQuery), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrievalVectorizeImage

> SingleVectorResult RetrievalVectorizeImage(ctx).ModelVersion(modelVersion).ImageUrl(imageUrl).Execute()

VectorizeImage



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tanmancan/openapi/azurecv"
)

func main() {
	modelVersion := "modelVersion_example" // string | Model version. Please refer https://aka.ms/image-retrieval for supported model versions.
	imageUrl := *openapiclient.NewImageUrl("Url_example") // ImageUrl | A JSON document with a URL pointing to the image that is to be analyzed. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.RetrievalVectorizeImage(context.Background()).ModelVersion(modelVersion).ImageUrl(imageUrl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RetrievalVectorizeImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrievalVectorizeImage`: SingleVectorResult
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.RetrievalVectorizeImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrievalVectorizeImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelVersion** | **string** | Model version. Please refer https://aka.ms/image-retrieval for supported model versions. | 
 **imageUrl** | [**ImageUrl**](ImageUrl.md) | A JSON document with a URL pointing to the image that is to be analyzed. | 

### Return type

[**SingleVectorResult**](SingleVectorResult.md)

### Authorization

[apiKeyQuery](../README.md#apiKeyQuery), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrievalVectorizeText

> SingleVectorResult RetrievalVectorizeText(ctx).ModelVersion(modelVersion).VectorizeTextRequest(vectorizeTextRequest).Execute()

VectorizeText



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tanmancan/openapi/azurecv"
)

func main() {
	modelVersion := "modelVersion_example" // string | Model version. Please refer https://aka.ms/image-retrieval for supported model versions.
	vectorizeTextRequest := *openapiclient.NewVectorizeTextRequest("Text_example") // VectorizeTextRequest | Request of VectorizeText. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.RetrievalVectorizeText(context.Background()).ModelVersion(modelVersion).VectorizeTextRequest(vectorizeTextRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RetrievalVectorizeText``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrievalVectorizeText`: SingleVectorResult
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.RetrievalVectorizeText`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrievalVectorizeTextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelVersion** | **string** | Model version. Please refer https://aka.ms/image-retrieval for supported model versions. | 
 **vectorizeTextRequest** | [**VectorizeTextRequest**](VectorizeTextRequest.md) | Request of VectorizeText. | 

### Return type

[**SingleVectorResult**](SingleVectorResult.md)

### Authorization

[apiKeyQuery](../README.md#apiKeyQuery), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

