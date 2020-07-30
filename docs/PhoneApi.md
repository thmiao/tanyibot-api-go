# \PhoneApi

All URIs are relative to *https://openapi.tanyibot.com/apiOpen/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPhoneList**](PhoneApi.md#GetPhoneList) | **Get** /phone/getPhoneList | 通过接口可以获取所有可用的外呼线路的列表
[**UpdatePhoneInfoByTenantPhoneNumberId**](PhoneApi.md#UpdatePhoneInfoByTenantPhoneNumberId) | **Post** /phone/updatePhoneInfoByTenantPhoneNumberId | 通过此接口修改线路的归属地、行业、黑名单，只能修改归属客户自己的线路
[**UpdatePhonePriceByTenantPhoneNumberId**](PhoneApi.md#UpdatePhonePriceByTenantPhoneNumberId) | **Post** /phone/updatePhonePriceByTenantPhoneNumberId | 修改绑定客户线路的价格



## GetPhoneList

> GetPhoneListRsp GetPhoneList(ctx, )

通过接口可以获取所有可用的外呼线路的列表

通过接口可以获取所有可用的外呼线路的列表

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**GetPhoneListRsp**](GetPhoneListRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePhoneInfoByTenantPhoneNumberId

> ApiResponse UpdatePhoneInfoByTenantPhoneNumberId(ctx, body)

通过此接口修改线路的归属地、行业、黑名单，只能修改归属客户自己的线路

通过此接口修改线路的归属地、行业、黑名单，只能修改归属客户自己的线路

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**UpdatePhoneInfoByTenantPhoneNumberIdReq**](UpdatePhoneInfoByTenantPhoneNumberIdReq.md)|  | 

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


## UpdatePhonePriceByTenantPhoneNumberId

> ApiResponse UpdatePhonePriceByTenantPhoneNumberId(ctx, )

修改绑定客户线路的价格

修改绑定客户线路的价格

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**ApiResponse**](APIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

