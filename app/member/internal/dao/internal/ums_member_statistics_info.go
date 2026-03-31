// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UmsMemberStatisticsInfoDao is the data access object for the table ums_member_statistics_info.
type UmsMemberStatisticsInfoDao struct {
	table    string                         // table is the underlying table name of the DAO.
	group    string                         // group is the database configuration group name of the current DAO.
	columns  UmsMemberStatisticsInfoColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler             // handlers for customized model modification.
}

// UmsMemberStatisticsInfoColumns defines and stores column names for the table ums_member_statistics_info.
type UmsMemberStatisticsInfoColumns struct {
	Id                  string //
	MemberId            string //
	ConsumeAmount       string // 累计消费金额
	OrderCount          string // 订单数量
	CouponCount         string // 优惠券数量
	CommentCount        string // 评价数
	ReturnOrderCount    string // 退货数量
	LoginCount          string // 登录次数
	AttendCount         string // 关注数量
	FansCount           string // 粉丝数量
	CollectProductCount string //
	CollectSubjectCount string //
	CollectTopicCount   string //
	CollectCommentCount string //
	InviteFriendCount   string //
	RecentOrderTime     string // 最后一次下订单时间
}

// umsMemberStatisticsInfoColumns holds the columns for the table ums_member_statistics_info.
var umsMemberStatisticsInfoColumns = UmsMemberStatisticsInfoColumns{
	Id:                  "id",
	MemberId:            "member_id",
	ConsumeAmount:       "consume_amount",
	OrderCount:          "order_count",
	CouponCount:         "coupon_count",
	CommentCount:        "comment_count",
	ReturnOrderCount:    "return_order_count",
	LoginCount:          "login_count",
	AttendCount:         "attend_count",
	FansCount:           "fans_count",
	CollectProductCount: "collect_product_count",
	CollectSubjectCount: "collect_subject_count",
	CollectTopicCount:   "collect_topic_count",
	CollectCommentCount: "collect_comment_count",
	InviteFriendCount:   "invite_friend_count",
	RecentOrderTime:     "recent_order_time",
}

// NewUmsMemberStatisticsInfoDao creates and returns a new DAO object for table data access.
func NewUmsMemberStatisticsInfoDao(handlers ...gdb.ModelHandler) *UmsMemberStatisticsInfoDao {
	return &UmsMemberStatisticsInfoDao{
		group:    "default",
		table:    "ums_member_statistics_info",
		columns:  umsMemberStatisticsInfoColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UmsMemberStatisticsInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UmsMemberStatisticsInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UmsMemberStatisticsInfoDao) Columns() UmsMemberStatisticsInfoColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UmsMemberStatisticsInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UmsMemberStatisticsInfoDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *UmsMemberStatisticsInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
