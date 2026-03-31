// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UmsMemberStatisticsInfo is the golang structure of table ums_member_statistics_info for DAO operations like Where/Data.
type UmsMemberStatisticsInfo struct {
	g.Meta              `orm:"table:ums_member_statistics_info, do:true"`
	Id                  any         //
	MemberId            any         //
	ConsumeAmount       any         // 累计消费金额
	OrderCount          any         // 订单数量
	CouponCount         any         // 优惠券数量
	CommentCount        any         // 评价数
	ReturnOrderCount    any         // 退货数量
	LoginCount          any         // 登录次数
	AttendCount         any         // 关注数量
	FansCount           any         // 粉丝数量
	CollectProductCount any         //
	CollectSubjectCount any         //
	CollectTopicCount   any         //
	CollectCommentCount any         //
	InviteFriendCount   any         //
	RecentOrderTime     *gtime.Time // 最后一次下订单时间
}
