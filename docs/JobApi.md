# \JobApi

All URIs are relative to *https://openapi.tanyibot.com/apiOpen/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](JobApi.md#Create) | **Post** /job/create | 通过此接口可以创建新的任务
[**ImportCustomer**](JobApi.md#ImportCustomer) | **Post** /job/importCustomer | 
[**JobCheckAllPost**](JobApi.md#JobCheckAllPost) | **Post** /job/checkAll | 
[**JobDeletePost**](JobApi.md#JobDeletePost) | **Post** /job/delete | 
[**JobExecuteJobsPost**](JobApi.md#JobExecuteJobsPost) | **Post** /job/executeJobs | 
[**JobGetJobDetailGet**](JobApi.md#JobGetJobDetailGet) | **Get** /job/getJobDetail | 
[**JobGetJobStatsInfoListPost**](JobApi.md#JobGetJobStatsInfoListPost) | **Post** /job/getJobStatsInfoList | 
[**JobGetJobsGet**](JobApi.md#JobGetJobsGet) | **Get** /job/getJobs | 获取任务列表接口
[**JobModifyPatch**](JobApi.md#JobModifyPatch) | **Patch** /job/modify | 修改任务接口
[**JobPausePost**](JobApi.md#JobPausePost) | **Post** /job/pause | 
[**JobReAddCustomerToJobPost**](JobApi.md#JobReAddCustomerToJobPost) | **Post** /job/reAddCustomerToJob | 
[**JobUpdateAiCountPost**](JobApi.md#JobUpdateAiCountPost) | **Post** /job/updateAiCount | 
[**Start**](JobApi.md#Start) | **Post** /job/start | 



## Create

> JobCreateRsp Create(ctx, body)

通过此接口可以创建新的任务

创建任务,支持使用多个无主叫固话,只需设置总并发数,由系统自动分配每个线路的并发。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Job**](Job.md)|  | 

### Return type

[**JobCreateRsp**](JobCreateRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportCustomer

> ApiResponse ImportCustomer(ctx, body)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**ImportCustomerReq**](ImportCustomerReq.md)|  | 

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


## JobCheckAllPost

> CheckAllRsp JobCheckAllPost(ctx, body)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**InlineObject4**](InlineObject4.md)|  | 

### Return type

[**CheckAllRsp**](CheckAllRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JobDeletePost

> ApiResponse JobDeletePost(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***JobDeletePostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a JobDeletePostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of InlineObject2**](InlineObject2.md)|  | 

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


## JobExecuteJobsPost

> ExecuteJobsRsp JobExecuteJobsPost(ctx, body)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**ExecuteJobsReq**](ExecuteJobsReq.md)|  | 

### Return type

[**ExecuteJobsRsp**](ExecuteJobsRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JobGetJobDetailGet

> GetJobDetailRsp JobGetJobDetailGet(ctx, robotCallJobId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**robotCallJobId** | **int64**| 任务Id | 

### Return type

[**GetJobDetailRsp**](GetJobDetailRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JobGetJobStatsInfoListPost

> GetJobStatsInfoListRsp JobGetJobStatsInfoListPost(ctx, body)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**InlineObject5**](InlineObject5.md)|  | 

### Return type

[**GetJobStatsInfoListRsp**](GetJobStatsInfoListRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JobGetJobsGet

> GetJobsRsp JobGetJobsGet(ctx, optional)

获取任务列表接口

通过此接口可以获取指定公司的任务列表

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***JobGetJobsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a JobGetJobsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| 任务名称 | 
 **status** | **optional.String**| 任务状态,(NOT_STARTED, \&quot;未开始\&quot;),(IN_PROCESS, \&quot;进行中\&quot;),(COMPLETED, \&quot;已完成\&quot;),(RUNNABLE, \&quot;可运行\&quot;),(USER_PAUSE, \&quot;用户暂停\&quot;),(SYSTEM_SUSPENDED, \&quot;系统暂停\&quot;),(TERMINATE, \&quot;已终止\&quot;),(IN_QUEUE, \&quot;排队中\&quot;),(SYSTEM_HANG_UP,\&quot;系统挂起\&quot;),(WAITING_FOR_REDIAL,\&quot;等待重呼\&quot;),(ACCOUNT_DISABLE,\&quot;账户禁用\&quot;),(MAINTAIN,\&quot;系统维护\&quot;); | 
 **pageNum** | **optional.Int32**| 第几页,默认1 | 
 **pageSize** | **optional.Int32**| 页面大小,选填,默认20,最大值100 | 

### Return type

[**GetJobsRsp**](GetJobsRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JobModifyPatch

> JobUpdateRsp JobModifyPatch(ctx, optional)

修改任务接口

通过此接口可以修改任务

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***JobModifyPatchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a JobModifyPatchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Job**](Job.md)|  | 

### Return type

[**JobUpdateRsp**](JobUpdateRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JobPausePost

> ApiResponse JobPausePost(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***JobPausePostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a JobPausePostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of InlineObject1**](InlineObject1.md)|  | 

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


## JobReAddCustomerToJobPost

> ApiResponse JobReAddCustomerToJobPost(ctx, body)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**ReAddCustomerToJobReq**](ReAddCustomerToJobReq.md)|  | 

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


## JobUpdateAiCountPost

> ApiResponse JobUpdateAiCountPost(ctx, body)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**InlineObject3**](InlineObject3.md)|  | 

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


## Start

> ApiResponse Start(ctx, body)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**InlineObject**](InlineObject.md)|  | 

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

