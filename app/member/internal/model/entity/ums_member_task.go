// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// UmsMemberTask is the golang structure for table ums_member_task.
type UmsMemberTask struct {
	Id           int64  `json:"id"           orm:"id"           description:""`                     //
	Name         string `json:"name"         orm:"name"         description:""`                     //
	Growth       int    `json:"growth"       orm:"growth"       description:"赠送成长值"`                // 赠送成长值
	Intergration int    `json:"intergration" orm:"intergration" description:"赠送积分"`                 // 赠送积分
	Type         int    `json:"type"         orm:"type"         description:"任务类型：0->新手任务；1->日常任务"` // 任务类型：0->新手任务；1->日常任务
}
