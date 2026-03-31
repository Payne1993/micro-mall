// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UmsMemberLoginLog is the golang structure for table ums_member_login_log.
type UmsMemberLoginLog struct {
	Id         int64       `json:"id"         orm:"id"          description:""`                                    //
	MemberId   int64       `json:"memberId"   orm:"member_id"   description:""`                                    //
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:""`                                    //
	Ip         string      `json:"ip"         orm:"ip"          description:""`                                    //
	City       string      `json:"city"       orm:"city"        description:""`                                    //
	LoginType  int         `json:"loginType"  orm:"login_type"  description:"登录类型：0->PC；1->android;2->ios;3->小程序"` // 登录类型：0->PC；1->android;2->ios;3->小程序
	Province   string      `json:"province"   orm:"province"    description:""`                                    //
}
