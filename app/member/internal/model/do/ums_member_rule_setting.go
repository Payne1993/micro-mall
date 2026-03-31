// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// UmsMemberRuleSetting is the golang structure of table ums_member_rule_setting for DAO operations like Where/Data.
type UmsMemberRuleSetting struct {
	g.Meta            `orm:"table:ums_member_rule_setting, do:true"`
	Id                any //
	ContinueSignDay   any // 连续签到天数
	ContinueSignPoint any // 连续签到赠送数量
	ConsumePerPoint   any // 每消费多少元获取1个点
	LowOrderAmount    any // 最低获取点数的订单金额
	MaxPointPerOrder  any // 每笔订单最高获取点数
	Type              any // 类型：0->积分规则；1->成长值规则
}
