# \DialogFlowApi

All URIs are relative to *https://openapi.tanyibot.com/apiOpen/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DialogFlowBatchUploadPost**](DialogFlowApi.md#DialogFlowBatchUploadPost) | **Post** /dialogFlow/batchUpload | 上传话术录音
[**DialogFlowCopyDialogFlowPost**](DialogFlowApi.md#DialogFlowCopyDialogFlowPost) | **Post** /dialogFlow/copyDialogFlow | 话术复制
[**DialogFlowGetDialogFlowCallJobRelatedInfoGet**](DialogFlowApi.md#DialogFlowGetDialogFlowCallJobRelatedInfoGet) | **Get** /dialogFlow/getDialogFlowCallJobRelatedInfo | 获取话术中存在人工介入和转人工等标识
[**DialogFlowGetTotalDialogFlowListPost**](DialogFlowApi.md#DialogFlowGetTotalDialogFlowListPost) | **Post** /dialogFlow/getTotalDialogFlowList | 话术列表
[**ExportDialogFlowWordDoc**](DialogFlowApi.md#ExportDialogFlowWordDoc) | **Get** /dialogFlow/exportDialogFlowWordDoc | 获取话术主流程word文档
[**GetDialogContentInfo**](DialogFlowApi.md#GetDialogContentInfo) | **Get** /dialogFlow/getDialogContentInfo | 获取话术中所有需要录音的文本和录音文件
[**GetDialogFlowList**](DialogFlowApi.md#GetDialogFlowList) | **Get** /dialogFlow/getDialogFlowList | 获取公司的机器人话术列表接口



## DialogFlowBatchUploadPost

> ApiResponse DialogFlowBatchUploadPost(ctx, dialogFlowId, file)

上传话术录音

通过接口可以上传话术中指定文本的录音

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dialogFlowId** | **int64**| 话术ID | 
**file** | ***os.File*****os.File**| 音频文件(MP3或者wav格式)，或者多个录音文件的zip包，录音文件名称为对应录音文本的label | 

### Return type

[**ApiResponse**](APIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DialogFlowCopyDialogFlowPost

> CopyDialogFlowRsp DialogFlowCopyDialogFlowPost(ctx, optional)

话术复制

通过接口可以复制指定话术

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DialogFlowCopyDialogFlowPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DialogFlowCopyDialogFlowPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CopyDialogFlowReq**](CopyDialogFlowReq.md)|  | 

### Return type

[**CopyDialogFlowRsp**](CopyDialogFlowRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DialogFlowGetDialogFlowCallJobRelatedInfoGet

> DialogFlowCallJobRelatedInfo DialogFlowGetDialogFlowCallJobRelatedInfoGet(ctx, dialogFlowId)

获取话术中存在人工介入和转人工等标识

通过接口可以获取指定话术的标识

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dialogFlowId** | **int64**|  | 

### Return type

[**DialogFlowCallJobRelatedInfo**](DialogFlowCallJobRelatedInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DialogFlowGetTotalDialogFlowListPost

> TotalDialogFlowListRsp DialogFlowGetTotalDialogFlowListPost(ctx, optional)

话术列表

通过接口可以获取所有话术列表

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DialogFlowGetTotalDialogFlowListPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DialogFlowGetTotalDialogFlowListPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of GetTotalDialogFlowListReq**](GetTotalDialogFlowListReq.md)|  | 

### Return type

[**TotalDialogFlowListRsp**](TotalDialogFlowListRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportDialogFlowWordDoc

> InlineResponse200 ExportDialogFlowWordDoc(ctx, dialogFlowId)

获取话术主流程word文档

通过接口可以获取指定话术的主流程文档

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dialogFlowId** | **int64**|  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDialogContentInfo

> DialogContentInfo GetDialogContentInfo(ctx, dialogFlowId)

获取话术中所有需要录音的文本和录音文件

通过接口可以获取指定话术所有需要录音的文本和录音文件

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dialogFlowId** | **int64**| 话术ID | 

### Return type

[**DialogContentInfo**](DialogContentInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDialogFlowList

> DialogFlowList GetDialogFlowList(ctx, )

获取公司的机器人话术列表接口

通过接口可以获取指定公司的所有配置完成的机器人话术

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**DialogFlowList**](DialogFlowList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

