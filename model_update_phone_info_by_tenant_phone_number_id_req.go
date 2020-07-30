/*
 * AIBot open api
 *
 * aibot api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package aibot
// UpdatePhoneInfoByTenantPhoneNumberIdReq 
type UpdatePhoneInfoByTenantPhoneNumberIdReq struct {
	// 线路归属地
	AreaCode string `json:"areaCode,omitempty"`
	// 线路归属地省
	Province string `json:"province,omitempty"`
	// 线路归属地市
	City string `json:"city,omitempty"`
	// (FINANCE, \"金融\"),(OTHER, \"其他\")
	CallOutIndustry string `json:"callOutIndustry,omitempty"`
	DeadArea UpdatePhoneInfoByTenantPhoneNumberIdReqDeadArea `json:"deadArea,omitempty"`
}
