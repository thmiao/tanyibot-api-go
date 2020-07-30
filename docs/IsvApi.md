# \IsvApi

All URIs are relative to *https://openapi.tanyibot.com/apiOpen/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenSubAccountIsv**](IsvApi.md#GenSubAccountIsv) | **Post** /isv/genSubAccountIsv | 通过此接口可生成子账号的ISV账号
[**GetIsvList**](IsvApi.md#GetIsvList) | **Post** /isv/getIsvList | 获取isv列表
[**UpdateIsvInfo**](IsvApi.md#UpdateIsvInfo) | **Post** /isv/updateIsvInfo | 修改ISV对象的公司签名和回调地址



## GenSubAccountIsv

> GetSubAccountIsvResp GenSubAccountIsv(ctx, body)

通过此接口可生成子账号的ISV账号

通过此接口可生成子账号的ISV账号

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**GetSubAccountIsvReq**](GetSubAccountIsvReq.md)|  | 

### Return type

[**GetSubAccountIsvResp**](GetSubAccountIsvResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIsvList

> GetIsvListRsp GetIsvList(ctx, )

获取isv列表

获取isv列表

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**GetIsvListRsp**](GetIsvListRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIsvInfo

> ApiResponse UpdateIsvInfo(ctx, body)

修改ISV对象的公司签名和回调地址

修改ISV对象的公司签名和回调地址

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**UpdateIsvInfoReq**](UpdateIsvInfoReq.md)|  | 

### Return type

[**ApiResponse**](APIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

