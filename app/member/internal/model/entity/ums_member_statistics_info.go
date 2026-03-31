// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UmsMemberStatisticsInfo is the golang structure for table ums_member_statistics_info.
type UmsMemberStatisticsInfo struct {
	Id                  int64       `json:"id"                  orm:"id"                    description:""`          //
	MemberId            int64       `json:"memberId"            orm:"member_id"             description:""`          //
	ConsumeAmount       float64     `json:"consumeAmount"       orm:"consume_amount"        description:"累计消费金额"`    // 累计消费金额
	OrderCount          int         `json:"orderCount"          orm:"order_count"           description:"订单数量"`      // 订单数量
	CouponCount         int         `json:"couponCount"         orm:"coupon_count"          description:"优惠券数量"`     // 优惠券数量
	CommentCount        int         `json:"commentCount"        orm:"comment_count"         description:"评价数"`       // 评价数
	ReturnOrderCount    int         `json:"returnOrderCount"    orm:"return_order_count"    description:"退货数量"`      // 退货数量
	LoginCount          int         `json:"loginCount"          orm:"login_count"           description:"登录次数"`      // 登录次数
	AttendCount         int         `json:"attendCount"         orm:"attend_count"          description:"关注数量"`      // 关注数量
	FansCount           int         `json:"fansCount"           orm:"fans_count"            description:"粉丝数量"`      // 粉丝数量
	CollectProductCount int         `json:"collectProductCount" orm:"collect_product_count" description:""`          //
	CollectSubjectCount int         `json:"collectSubjectCount" orm:"collect_subject_count" description:""`          //
	CollectTopicCount   int         `json:"collectTopicCount"   orm:"collect_topic_count"   description:""`          //
	CollectCommentCount int         `json:"collectCommentCount" orm:"collect_comment_count" description:""`          //
	InviteFriendCount   int         `json:"inviteFriendCount"   orm:"invite_friend_count"   description:""`          //
	RecentOrderTime     *gtime.Time `json:"recentOrderTime"     orm:"recent_order_time"     description:"最后一次下订单时间"` // 最后一次下订单时间
}
