# \ReviewApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateReview**](ReviewApi.md#CreateReview) | **Post** /api/review/{articleID}/{userID} | Submit review


# **CreateReview**
> CreateReview(ctx, userID, articleID, content)
Submit review

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userID** | **int32**| The id of user. | 
  **articleID** | **int32**| The id of article. | 
  **content** | [**Review**](Review.md)| The content of review. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

