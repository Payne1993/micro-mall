// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UmsMemberLoginLog is the golang structure of table ums_member_login_log for DAO operations like Where/Data.
type UmsMemberLoginLog struct {
	g.Meta     `orm:"table:ums_member_login_log, do:true"`
	Id         any         //
	MemberId   any         //
	CreateTime *gtime.Time //
	Ip         any         //
	City       any         //
	LoginType  any         // 登录类型：0->PC；1->android;2->ios;3->小程序
	Province   any         //
}
