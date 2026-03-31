// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// UmsMemberTask is the golang structure of table ums_member_task for DAO operations like Where/Data.
type UmsMemberTask struct {
	g.Meta       `orm:"table:ums_member_task, do:true"`
	Id           any //
	Name         any //
	Growth       any // 赠送成长值
	Intergration any // 赠送积分
	Type         any // 任务类型：0->新手任务；1->日常任务
}
