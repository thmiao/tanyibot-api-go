# \PolicyGroupApi

All URIs are relative to *https://openapi.tanyibot.com/apiOpen/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PolicyGroupCreatePost**](PolicyGroupApi.md#PolicyGroupCreatePost) | **Post** /policyGroup/create | 创建外呼策略组接口
[**PolicyGroupGetIdAndNamePairListGet**](PolicyGroupApi.md#PolicyGroupGetIdAndNamePairListGet) | **Get** /policyGroup/getIdAndNamePairList | 获取外呼策略组选择接口
[**PolicyGroupListGet**](PolicyGroupApi.md#PolicyGroupListGet) | **Get** /policyGroup/list | 获取外呼策略组列表接口
[**PolicyGroupUpdatePost**](PolicyGroupApi.md#PolicyGroupUpdatePost) | **Post** /policyGroup/update | 更新外呼策略组接口



## PolicyGroupCreatePost

> ApiResponse PolicyGroupCreatePost(ctx, optional)

创建外呼策略组接口

创建新的外呼策略组

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyGroupCreatePostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PolicyGroupCreatePostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CreatePolicyGroupReq**](CreatePolicyGroupReq.md)|  | 

### Return type

[**ApiResponse**](APIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

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


## PolicyGroupUpdatePost

> ApiResponse PolicyGroupUpdatePost(ctx, optional)

更新外呼策略组接口

更新新的外呼策略组

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyGroupUpdatePostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PolicyGroupUpdatePostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of UpdatePolicyGroupReq**](UpdatePolicyGroupReq.md)|  | 

### Return type

[**ApiResponse**](APIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

