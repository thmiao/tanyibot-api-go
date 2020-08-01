/*
 * Tanyi Bot open api
 *
 * tanyibot api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tanyibot
// CopyDialogFlowReq struct for CopyDialogFlowReq
type CopyDialogFlowReq struct {
	// 被复制话术ID
	DialogFlowTemplateId int64 `json:"dialogFlowTemplateId"`
	// 话术复制
	Name string `json:"name"`
	// 测试话术
	Description string `json:"description,omitempty"`
	// 意向规则分组id
	IntentLevelTagId int64 `json:"intentLevelTagId,omitempty"`
	// 打断灵敏度 [1-9]
	VadGateMute int32 `json:"vadGateMute,omitempty"`
	// 反应灵敏度 [1-9]
	MaxSentenceSilence int32 `json:"maxSentenceSilence,omitempty"`
	// 是否启用问法
	EnableAskService bool `json:"enableAskService,omitempty"`
	BranchLevel []string `json:"branchLevel,omitempty"`
}
