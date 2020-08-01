# RobotCallJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DialogFlowId** | **int32** |  | [optional] 
**IsPrior** | **bool** |  | [optional] 
**RobotCallJobId** | **int64** | 任务id | [optional] 
**DailyEndTime** | **string** |  | [optional] 
**DailyStartTime** | **string** |  | [optional] 
**Description** | **string** |  | [optional] 
**EarlyWarningAlertUsers** | **[]string** |  | [optional] 
**InactiveTimeList** | [**[]RobotCallJobInactiveTimeList**](RobotCallJob_inactiveTimeList.md) |  | [optional] 
**Name** | **string** | 任务名称 | [optional] 
**Mode** | **string** | 任务类型 (AUTO, \&quot;自动任务\&quot;),(MANUAL, \&quot;手动任务\&quot;); | [optional] 
**PhoneType** | **string** | 号码类型 (MOBILE, \&quot;手机号码\&quot;),(LANDLINE, \&quot;固话\&quot;),(UNFIXED_CALL, \&quot;无主叫\&quot;), (CALL_POLICY_GROUP, \&quot;外呼策略组\&quot;) | [optional] 
**RobotCount** | **int32** |  | [optional] 
**SmsAlertLevel** | **[]string** |  | [optional] 
**SmsTemplateId** | **int32** |  | [optional] 
**StartTime** | **string** |  | [optional] 
**TenantId** | **int64** |  | [optional] 
**SmsAlertLevelCode** | **[]int32** |  | [optional] 
**WechatPushConditionList** | [**[]WechatPushCondition**](WechatPushCondition.md) | 微信推送条件 | [optional] 
**RedialCondition** | **[]string** | 重拨条件(CALL_LOSS,\&quot;呼损客户\&quot;),(NO_ANSWER,\&quot;无应答\&quot;),(BUSY,\&quot;忙线中\&quot;),(REFUSED,\&quot;拒接\&quot;),(POWER_OFF,\&quot;关机\&quot;),(OUT_OF_SERVICE,\&quot;停机\&quot;),(CAN_NOT_CONNECT,“无法接通“),(FROM_PHONE_ERROR,\&quot;主叫欠费\&quot;),(SYSTEM_ERROR,\&quot;外呼失败\&quot;) | [optional] 
**RedialInterval** | **int32** | 重拨间隔(分钟，取值范围6分钟~24 x 60分钟) | [optional] 
**RedialTimes** | **int32** | 重拨次数(取值范围1~10） | [optional] 
**TenantCallInterceptId** | **int64** | 拦截组id | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


