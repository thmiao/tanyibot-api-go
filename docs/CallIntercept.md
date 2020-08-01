# CallIntercept

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantCallInterceptId** | **int64** | 拦截组Id | [optional] 
**Name** | **string** | 拦截组名称 | [optional] 
**CallOutRestrict** | **string** | 是否开启外呼上限设置(YES 开启 NO 不开启) | [optional] 
**CallOutCount** | **int32** | 外呼次数上限限制 | [optional] 
**CallOutDays** | **int32** | 外呼次数天数内限制 | [optional] 
**NotExistDays** | **int32** | 空号天数限制 0为取消 | [optional] 
**NotServiceDays** | **int32** | 停机天数限制 0为取消 | [optional] 
**Type** | **string** | 生效范围 ALL 全网 TENANT 公司内部 | [optional] 
**Source** | **string** | 拦截组来源 SYSTEM 系统规则 CUSTOM 自定义规则 | [optional] 
**DefaultStatus** | **bool** | 是否默认规则 | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


