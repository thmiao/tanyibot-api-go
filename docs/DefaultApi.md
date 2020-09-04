# \DefaultApi

All URIs are relative to *https://openapi.tanyibot.com/apiOpen/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CallInterceptListGet**](DefaultApi.md#CallInterceptListGet) | **Get** /callIntercept/list | 获取拦截组列表接口
[**DialogFlowBatchUploadPost**](DefaultApi.md#DialogFlowBatchUploadPost) | **Post** /dialogFlow/batchUpload | 上传话术录音
[**DialogFlowCopyDialogFlowPost**](DefaultApi.md#DialogFlowCopyDialogFlowPost) | **Post** /dialogFlow/copyDialogFlow | 话术复制
[**DialogFlowGetDialogFlowCallJobRelatedInfoGet**](DefaultApi.md#DialogFlowGetDialogFlowCallJobRelatedInfoGet) | **Get** /dialogFlow/getDialogFlowCallJobRelatedInfo | 获取话术中存在人工介入和转人工等标识
[**DialogFlowGetDialogFlowListGet**](DefaultApi.md#DialogFlowGetDialogFlowListGet) | **Get** /dialogFlow/getDialogFlowList | 获取公司的机器人话术列表接口
[**DialogFlowGetTotalDialogFlowListPost**](DefaultApi.md#DialogFlowGetTotalDialogFlowListPost) | **Post** /dialogFlow/getTotalDialogFlowList | 话术列表
[**ExportDialogFlowWordDoc**](DefaultApi.md#ExportDialogFlowWordDoc) | **Get** /dialogFlow/exportDialogFlowWordDoc | 获取话术主流程word文档
[**GetDialogContentInfo**](DefaultApi.md#GetDialogContentInfo) | **Get** /dialogFlow/getDialogContentInfo | 获取话术中所有需要录音的文本和录音文件
[**PolicyGroupGetIdAndNamePairListGet**](DefaultApi.md#PolicyGroupGetIdAndNamePairListGet) | **Get** /policyGroup/getIdAndNamePairList | 获取外呼策略组选择接口
[**PolicyGroupListGet**](DefaultApi.md#PolicyGroupListGet) | **Get** /policyGroup/list | 获取外呼策略组列表接口



## CallInterceptListGet

> InterceptListRsp CallInterceptListGet(ctx, )

获取拦截组列表接口

通过此接口可以获取所有拦截组

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**InterceptListRsp**](InterceptListRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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

> CopyDialogFlowReq DialogFlowCopyDialogFlowPost(ctx, optional)

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
 **body** | **optional.Map[string]interface{}**|  | 

### Return type

[**CopyDialogFlowReq**](CopyDialogFlowReq.md)

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


## DialogFlowGetDialogFlowListGet

> DialogFlowList DialogFlowGetDialogFlowListGet(ctx, )

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


## PolicyGroupGetIdAndNamePairListGet

> GetIdAndNamePairListRsp PolicyGroupGetIdAndNamePairListGet(ctx, )

获取外呼策略组选择接口

通过此接口可以获取外呼策略组id和名称对应的键值对列表

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**GetIdAndNamePairListRsp**](GetIdAndNamePairListRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PolicyGroupListGet

> GetPolicyGroupListRsp PolicyGroupListGet(ctx, )

获取外呼策略组列表接口

通过此接口可以获取用户所有的外呼策略组信息

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**GetPolicyGroupListRsp**](GetPolicyGroupListRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

