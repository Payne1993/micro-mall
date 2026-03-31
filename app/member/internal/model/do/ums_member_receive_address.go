// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// UmsMemberReceiveAddress is the golang structure of table ums_member_receive_address for DAO operations like Where/Data.
type UmsMemberReceiveAddress struct {
	g.Meta        `orm:"table:ums_member_receive_address, do:true"`
	Id            any //
	MemberId      any //
	Name          any // 收货人名称
	PhoneNumber   any //
	DefaultStatus any // 是否为默认
	PostCode      any // 邮政编码
	Province      any // 省份/直辖市
	City          any // 城市
	Region        any // 区
	DetailAddress any // 详细地址(街道)
}
