# InlineObject4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operation** | **string** | 操作类型，必填 PAUSE:暂停 START:开始 | 
**RobotCallJobIdSet** | **[]string** | 任务id集合 | [optional] 
**StatusSet** | **string** | 任务状态，NOT_STARTED(0, \&quot;未开始\&quot;),IN_PROCESS(1, \&quot;进行中\&quot;)，COMPLETED(2, \&quot;已完成\&quot;),RUNNABLE(3, \&quot;可运行\&quot;),USER_PAUSE(4, \&quot;用户暂停\&quot;),SYSTEM_SUSPENDED(5, \&quot;系统暂停\&quot;)，TERMINATE(6, \&quot;已终止\&quot;),IN_QUEUE(7, \&quot;排队中\&quot;),SYSTEM_HANG_UP(10, \&quot;系统挂起\&quot;),WAITING_FOR_REDIAL(11, \&quot;等待重呼\&quot;),ACCOUNT_DISABLE(12, \&quot;账户禁用\&quot;),MAINTAIN(13, \&quot;系统维护\&quot;); | [optional] 
**Name** | **string** | 任务名称 | [optional] 
**BeginCreateDate** | **string** | 开始任务创建日期 | [optional] 
**EndCreateDate** | **string** | 结束任务创建日期 | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


