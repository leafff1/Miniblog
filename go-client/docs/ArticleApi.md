# \ArticleApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddArticle**](ArticleApi.md#AddArticle) | **Post** /api/article/ | Add a new article
[**DeleteArticle**](ArticleApi.md#DeleteArticle) | **Delete** /api/article/{articleID} | Delete an existing article
[**GetAllArticle**](ArticleApi.md#GetAllArticle) | **Get** /api/article/ | Get all articles
[**GetArticleByID**](ArticleApi.md#GetArticleByID) | **Get** /api/article/{articleID} | Get article by ID
[**GetArticleByTag**](ArticleApi.md#GetArticleByTag) | **Get** /api/article/tag/{tagID} | Get article by Tag
[**GetArticleByUser**](ArticleApi.md#GetArticleByUser) | **Get** /api/article/user/{userID} | Get article by User
[**UpdateArticle**](ArticleApi.md#UpdateArticle) | **Put** /api/article/{articleID} | Update an existing article


# **AddArticle**
> interface{} AddArticle(ctx, body)
Add a new article

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Article**](Article.md)| The body of article. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteArticle**
> DeleteArticle(ctx, articleID)
Delete an existing article

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleID** | **int32**| The id of article. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllArticle**
> []Article GetAllArticle(ctx, )
Get all articles

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Article**](Article.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetArticleByID**
> []Article GetArticleByID(ctx, articleID)
Get article by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleID** | **int32**| The id of article. | 

### Return type

[**[]Article**](Article.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetArticleByTag**
> []Article GetArticleByTag(ctx, tagID)
Get article by Tag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tagID** | **int32**| The tagid of article. | 

### Return type

[**[]Article**](Article.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetArticleByUser**
> []Article GetArticleByUser(ctx, userID)
Get article by User

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userID** | **int32**| The userid of article. | 

### Return type

[**[]Article**](Article.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateArticle**
> UpdateArticle(ctx, articleID, body)
Update an existing article

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleID** | **int32**| The id of article. | 
  **body** | [**Article**](Article.md)| The body of article. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

