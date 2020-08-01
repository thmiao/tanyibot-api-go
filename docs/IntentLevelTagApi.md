# \IntentLevelTagApi

All URIs are relative to *https://openapi.tanyibot.com/apiOpen/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUsedIntentLevelTagList**](IntentLevelTagApi.md#GetUsedIntentLevelTagList) | **Get** /intentLevelTag/getUsedIntentLevelTagList | 获取已使用的意向标签组列表接口
[**IntentLevelTagGetIntentLevelTagGet**](IntentLevelTagApi.md#IntentLevelTagGetIntentLevelTagGet) | **Get** /intentLevelTag/getIntentLevelTag | 获取意向标签内容



## GetUsedIntentLevelTagList

> GetUsedIntentLevelTagListRsp GetUsedIntentLevelTagList(ctx, )

获取已使用的意向标签组列表接口

通过此接口可以获取已使用的所有意向标签组

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**GetUsedIntentLevelTagListRsp**](GetUsedIntentLevelTagListRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntentLevelTagGetIntentLevelTagGet

> GetIntentLevelTagRsp IntentLevelTagGetIntentLevelTagGet(ctx, intentLevelTagId)

获取意向标签内容

通过接口可以获取指定意向标签分组的内容

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**intentLevelTagId** | **int32**| 标签分组Id | [default to 0]

### Return type

[**GetIntentLevelTagRsp**](GetIntentLevelTagRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

