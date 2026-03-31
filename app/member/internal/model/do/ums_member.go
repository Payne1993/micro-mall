// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UmsMember is the golang structure of table ums_member for DAO operations like Where/Data.
type UmsMember struct {
	g.Meta                `orm:"table:ums_member, do:true"`
	Id                    any         //
	MemberLevelId         any         //
	Username              any         // 用户名
	Password              any         // 密码
	Nickname              any         // 昵称
	Phone                 any         // 手机号码
	Status                any         // 帐号启用状态:0->禁用；1->启用
	CreateTime            *gtime.Time // 注册时间
	Icon                  any         // 头像
	Gender                any         // 性别：0->未知；1->男；2->女
	Birthday              *gtime.Time // 生日
	City                  any         // 所做城市
	Job                   any         // 职业
	PersonalizedSignature any         // 个性签名
	SourceType            any         // 用户来源
	Integration           any         // 积分
	Growth                any         // 成长值
	LuckeyCount           any         // 剩余抽奖次数
	HistoryIntegration    any         // 历史积分数量
}
