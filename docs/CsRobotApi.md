# \CsRobotApi

All URIs are relative to *https://openapi.tanyibot.com/apiOpen/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StaffGroupList**](CsRobotApi.md#StaffGroupList) | **Get** /csRobot/staffGroupList | 获取人工坐席列表接口



## StaffGroupList

> StaffGroupListRsp StaffGroupList(ctx, pageNum, pageSize)

获取人工坐席列表接口

通过此接口可以获取用户所有的有效坐席组列表信息

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageNum** | **int32**|  | 
**pageSize** | **int32**|  | 

### Return type

[**StaffGroupListRsp**](StaffGroupListRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

