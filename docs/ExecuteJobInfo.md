# ExecuteJobInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RobotCallJobId** | **int64** | 任务id | [optional] 
**RobotCallJobName** | **string** | 任务名 | [optional] 
**CreateTime** | **string** | 创建时间 | [optional] 
**CreateUserName** | **string** | 任务创建人 | [optional] 
**Status** | **string** | 任务状态，NOT_STARTED(0, \&quot;未开始\&quot;),IN_PROCESS(1, \&quot;进行中\&quot;)，COMPLETED(2, \&quot;已完成\&quot;),RUNNABLE(3, \&quot;可运行\&quot;),USER_PAUSE(4, \&quot;用户暂停\&quot;),SYSTEM_SUSPENDED(5, \&quot;系统暂停\&quot;)，TERMINATE(6, \&quot;已终止\&quot;),IN_QUEUE(7, \&quot;排队中\&quot;),SYSTEM_HANG_UP(10, \&quot;系统挂起\&quot;),WAITING_FOR_REDIAL(11, \&quot;等待重呼\&quot;),ACCOUNT_DISABLE(12, \&quot;账户禁用\&quot;),MAINTAIN(13, \&quot;系统维护\&quot;); | [optional] 
**HangUpType** | **string** | 系统挂起类型 ACCOUNT_DEBT(0, \&quot;账户欠费\&quot;, \&quot;使用的线路账户已欠费\&quot;),NO_ROBOT_AVAILABLE(1, \&quot;没有可用坐席\&quot;, \&quot;当前没有可用坐席\&quot;),PHONE_UNBIND(2, \&quot;线路已解绑\&quot;, \&quot;使用的线路已解绑\&quot;),LINE_BREAKDOWN(3, \&quot;线路故障\&quot;, \&quot;使用的线路状态均为故障\&quot;),AVAILABLE_ROBOT_NOT_ENOUGH(4,\&quot;有效坐席数不足\&quot;,\&quot;有效坐席数不足，请检查有效AI坐席数量！\&quot;); | [optional] 
**HangUpReason** | **string** | 系统挂起原因 | [optional] 
**StatsInfo** | [**StatsInfo**](StatsInfo.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


