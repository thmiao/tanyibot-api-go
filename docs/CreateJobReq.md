# CreateJobReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConcurrencyQuota** | **int32** | 并发数 开启一线多并发时可比robotCount大，不开启一线多并发是与robotCount相同（线路类型为手机号的时候可不传） | [optional] 
**JobPhoneNumberIdList** | **[]int32** | 通过获取外呼线路接口获取 tenant_phone_number_id,当类型是手机号的时候他的size代表机器人的个数，当类型非手机号的时候他的size只能是1；如果外呼方式选择的是外呼策略组，则里面内容为外呼策略组的id（size只能为1） | [optional] 
**TransferPhoneNumber** | **[]string** | 转人工号码,触发转人工时轮训号码列表 | [optional] 
**RobotCallJob** | [**map[string]interface{}**](object.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


