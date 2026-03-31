// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// UmsMemberRuleSetting is the golang structure for table ums_member_rule_setting.
type UmsMemberRuleSetting struct {
	Id                int64   `json:"id"                orm:"id"                  description:""`                    //
	ContinueSignDay   int     `json:"continueSignDay"   orm:"continue_sign_day"   description:"连续签到天数"`              // 连续签到天数
	ContinueSignPoint int     `json:"continueSignPoint" orm:"continue_sign_point" description:"连续签到赠送数量"`            // 连续签到赠送数量
	ConsumePerPoint   float64 `json:"consumePerPoint"   orm:"consume_per_point"   description:"每消费多少元获取1个点"`         // 每消费多少元获取1个点
	LowOrderAmount    float64 `json:"lowOrderAmount"    orm:"low_order_amount"    description:"最低获取点数的订单金额"`         // 最低获取点数的订单金额
	MaxPointPerOrder  int     `json:"maxPointPerOrder"  orm:"max_point_per_order" description:"每笔订单最高获取点数"`          // 每笔订单最高获取点数
	Type              int     `json:"type"              orm:"type"                description:"类型：0->积分规则；1->成长值规则"` // 类型：0->积分规则；1->成长值规则
}
