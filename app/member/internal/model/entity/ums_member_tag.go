// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// UmsMemberTag is the golang structure for table ums_member_tag.
type UmsMemberTag struct {
	Id                int64   `json:"id"                orm:"id"                  description:""`            //
	Name              string  `json:"name"              orm:"name"                description:""`            //
	FinishOrderCount  int     `json:"finishOrderCount"  orm:"finish_order_count"  description:"自动打标签完成订单数量"` // 自动打标签完成订单数量
	FinishOrderAmount float64 `json:"finishOrderAmount" orm:"finish_order_amount" description:"自动打标签完成订单金额"` // 自动打标签完成订单金额
}
