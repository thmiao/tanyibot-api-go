# Phone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantPhoneNumberId** | **int64** | 用户线路ID（代表绑给用户的虚拟线路） | [optional] 
**PhoneNumber** | **string** | 电话号码 | [optional] 
**PhoneName** | **string** | 电话号码名称 | [optional] 
**PhoneType** | **string** | 电话类型 (MOBILE，\&quot;手机\&quot;),(LANDLINE，\&quot;固话\&quot;),(UNFIXED_CALL，\&quot;无主叫\&quot;),(VOIP_DEVICE，\&quot;网关设备\&quot;), (CS_SEAT，\&quot;人工外呼\&quot;),(MESSAGE，\&quot;短信\&quot;),(CALL_POLICY_GROUP，\&quot;外呼策略组\&quot;) | [optional] 
**LocalBillRate** | **int32** | 本地话费（单位：厘） | [optional] 
**OtherBillRate** | **int32** | 外地话费（单位：厘） | [optional] 
**CallOutIndustry** | **string** | (FINANCE, \&quot;金融\&quot;),(OTHER, \&quot;其他\&quot;) | [optional] 
**AreaCode** | **string** | 归属地区号 | [optional] 
**Province** | **string** | 归属地省 | [optional] 
**City** | **string** | 归属地市 | [optional] 
**DeadArea** | [**[]PhoneDeadArea**](Phone_deadArea.md) | 无法外呼地区 | [optional] 
**TenantId** | **int64** |  | [optional] 
**Concurrency** | **int32** |  | [optional] 
**PhoneNumberId** | **int32** | 线路ID（代表实际的线路） | [optional] 
**AccountFare** | **int32** | 账户余额，单位（厘） | [optional] 
**EnabledStatus** | **int32** | 是否启用，1为启用，0为禁用 | [optional] 
**MonthlyBillRate** | **int32** | 呼入月租费率，单位（厘） | [optional] 
**CallInLocalBillRate** | **int32** | 呼入本地通话费用，单位（厘） | [optional] 
**CallInOtherBillRate** | **int32** | 呼入外地通话费用，单位（厘） | [optional] 
**CallInBillMode** | **string** | 呼入计费方式 (MONTHLY，\&quot;按月租计费\&quot;),(CHAT_TIME，\&quot;按通话时长(分钟)计费\&quot;) | [optional] 
**LastHeartBeatTime** | **string** | 最后外呼时间 | [optional] 
**ConcurrenceLimit** | **int32** | 并发数限制 | [optional] 
**Remark** | **string** | 备注 | [optional] 
**SipAccount** | **string** | SIP账号 | [optional] 
**OwnerName** | **string** | 归属公司名称 | [optional] 
**LocationDisplayType** | **string** | 线路归属地显示方式（DEFAULT，\&quot;默认显示配置归属地\&quot;）,（RANDOM，\&quot;全国随机显示\&quot;）,（NO_DISPLAY，\&quot;不显示\&quot;） | [optional] 
**OwnerTenantId** | **int64** | 归属公司ID | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


