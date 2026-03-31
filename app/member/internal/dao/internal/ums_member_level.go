// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UmsMemberLevelDao is the data access object for the table ums_member_level.
type UmsMemberLevelDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  UmsMemberLevelColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// UmsMemberLevelColumns defines and stores column names for the table ums_member_level.
type UmsMemberLevelColumns struct {
	Id                    string //
	Name                  string //
	GrowthPoint           string //
	DefaultStatus         string // 是否为默认等级：0->不是；1->是
	FreeFreightPoint      string // 免运费标准
	CommentGrowthPoint    string // 每次评价获取的成长值
	PriviledgeFreeFreight string // 是否有免邮特权
	PriviledgeSignIn      string // 是否有签到特权
	PriviledgeComment     string // 是否有评论获奖励特权
	PriviledgePromotion   string // 是否有专享活动特权
	PriviledgeMemberPrice string // 是否有会员价格特权
	PriviledgeBirthday    string // 是否有生日特权
	Note                  string //
}

// umsMemberLevelColumns holds the columns for the table ums_member_level.
var umsMemberLevelColumns = UmsMemberLevelColumns{
	Id:                    "id",
	Name:                  "name",
	GrowthPoint:           "growth_point",
	DefaultStatus:         "default_status",
	FreeFreightPoint:      "free_freight_point",
	CommentGrowthPoint:    "comment_growth_point",
	PriviledgeFreeFreight: "priviledge_free_freight",
	PriviledgeSignIn:      "priviledge_sign_in",
	PriviledgeComment:     "priviledge_comment",
	PriviledgePromotion:   "priviledge_promotion",
	PriviledgeMemberPrice: "priviledge_member_price",
	PriviledgeBirthday:    "priviledge_birthday",
	Note:                  "note",
}

// NewUmsMemberLevelDao creates and returns a new DAO object for table data access.
func NewUmsMemberLevelDao(handlers ...gdb.ModelHandler) *UmsMemberLevelDao {
	return &UmsMemberLevelDao{
		group:    "default",
		table:    "ums_member_level",
		columns:  umsMemberLevelColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UmsMemberLevelDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UmsMemberLevelDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UmsMemberLevelDao) Columns() UmsMemberLevelColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UmsMemberLevelDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UmsMemberLevelDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *UmsMemberLevelDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
