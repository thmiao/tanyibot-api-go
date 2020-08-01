# ReAddCustomerToJobReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RobotCallJobId** | **int64** |  | 
**CustomerPersonIds** | **[]int64** |  | 
**LastCallRecord** | **bool** | 是否过滤重复通话记录只显示最近的一条 | [optional] 
**CalledPhoneNumber** | **string** | 被叫号码 | [optional] 
**CallRecordId** | **int64** | 通话记录id | [optional] 
**CustomerPersonName** | **string** | 客户姓名 | [optional] 
**CustomerGroupId** | **int64** | 分组id | [optional] 
**ResultStatuses** | **[]string** | ANSWERED(0, \&quot;已接听\&quot;, null), NO_ANSWER(1, \&quot;无应答\&quot;, \&quot;呼叫号码未接听\&quot;), BUSY(2, \&quot;忙线中\&quot;, \&quot;呼叫号码占线\&quot;), POWER_OFF(3, \&quot;关机\&quot;, \&quot;呼叫号码关机\&quot;), OUT_OF_SERVICE(4, \&quot;停机\&quot;, \&quot;呼叫号码停机\&quot;), REFUSED(5, \&quot;拒接\&quot;, \&quot;呼叫号码拒接\&quot;), VACANT_NUMBER(6, \&quot;空号\&quot;, \&quot;呼叫的号码是空号\&quot;), CAN_NOT_CONNECT(7, \&quot;无法接通\&quot;, \&quot;呼叫的号码无法接通\&quot;), FROM_PHONE_ERROR(8, \&quot;主叫停机\&quot;, \&quot;主叫号码不可用\&quot;), SYSTEM_ERROR(9, \&quot;系统错误\&quot;, \&quot;外呼失败，系统错误\&quot;), CALL_LOSS(10,\&quot;多并发呼损\&quot;,\&quot;等待服务中用户挂断\&quot;), TRANSFER_ARTIFICIAL(11,\&quot;转人工呼损\&quot;,\&quot;等待转人工服务中用户挂断\&quot;); [\&quot;ANSWERED\&quot;, | [optional] 
**IntentLevels** | **[]string** |  | [optional] 
**FollowStatus** | **string** | 跟进状态CLUE(0, \&quot;线索\&quot;), AI_INITIAL_VISIT(1, \&quot;AI初访\&quot;), PEOPLE_INITIAL_VISIT(2, \&quot;人工初访\&quot;), AI_FOLLOW_UP(3, \&quot;AI持续跟进\&quot;), PEOPLE_FOLLOW_UP(4, \&quot;人工持续跟进\&quot;), QUOTE(5, \&quot;报价\&quot;), DEAL(6, \&quot;成交\&quot;), LOSS(7, \&quot;流失\&quot;); | [optional] 
**DialogFlowId** | **int64** | 话术ID | [optional] 
**EarliestCreationTime** | **string** | 最早创建时间，日期标准格式，请不要包含时间。可以为空 | [optional] 
**LatestCreationTime** | **string** | 最晚创建时间，日期标准格式，请不要包含时间。可以为空 | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


