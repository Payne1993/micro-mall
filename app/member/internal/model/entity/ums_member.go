// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UmsMember is the golang structure for table ums_member.
type UmsMember struct {
	Id                    int64       `json:"id"                    orm:"id"                     description:""`                   //
	MemberLevelId         int64       `json:"memberLevelId"         orm:"member_level_id"        description:""`                   //
	Username              string      `json:"username"              orm:"username"               description:"用户名"`                // 用户名
	Password              string      `json:"password"              orm:"password"               description:"密码"`                 // 密码
	Nickname              string      `json:"nickname"              orm:"nickname"               description:"昵称"`                 // 昵称
	Phone                 string      `json:"phone"                 orm:"phone"                  description:"手机号码"`               // 手机号码
	Status                int         `json:"status"                orm:"status"                 description:"帐号启用状态:0->禁用；1->启用"` // 帐号启用状态:0->禁用；1->启用
	CreateTime            *gtime.Time `json:"createTime"            orm:"create_time"            description:"注册时间"`               // 注册时间
	Icon                  string      `json:"icon"                  orm:"icon"                   description:"头像"`                 // 头像
	Gender                int         `json:"gender"                orm:"gender"                 description:"性别：0->未知；1->男；2->女"` // 性别：0->未知；1->男；2->女
	Birthday              *gtime.Time `json:"birthday"              orm:"birthday"               description:"生日"`                 // 生日
	City                  string      `json:"city"                  orm:"city"                   description:"所做城市"`               // 所做城市
	Job                   string      `json:"job"                   orm:"job"                    description:"职业"`                 // 职业
	PersonalizedSignature string      `json:"personalizedSignature" orm:"personalized_signature" description:"个性签名"`               // 个性签名
	SourceType            int         `json:"sourceType"            orm:"source_type"            description:"用户来源"`               // 用户来源
	Integration           int         `json:"integration"           orm:"integration"            description:"积分"`                 // 积分
	Growth                int         `json:"growth"                orm:"growth"                 description:"成长值"`                // 成长值
	LuckeyCount           int         `json:"luckeyCount"           orm:"luckey_count"           description:"剩余抽奖次数"`             // 剩余抽奖次数
	HistoryIntegration    int         `json:"historyIntegration"    orm:"history_integration"    description:"历史积分数量"`             // 历史积分数量
}
