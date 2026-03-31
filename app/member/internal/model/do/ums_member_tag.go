// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// UmsMemberTag is the golang structure of table ums_member_tag for DAO operations like Where/Data.
type UmsMemberTag struct {
	g.Meta            `orm:"table:ums_member_tag, do:true"`
	Id                any //
	Name              any //
	FinishOrderCount  any // 自动打标签完成订单数量
	FinishOrderAmount any // 自动打标签完成订单金额
}
