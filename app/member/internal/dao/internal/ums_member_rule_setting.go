// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UmsMemberRuleSettingDao is the data access object for the table ums_member_rule_setting.
type UmsMemberRuleSettingDao struct {
	table    string                      // table is the underlying table name of the DAO.
	group    string                      // group is the database configuration group name of the current DAO.
	columns  UmsMemberRuleSettingColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler          // handlers for customized model modification.
}

// UmsMemberRuleSettingColumns defines and stores column names for the table ums_member_rule_setting.
type UmsMemberRuleSettingColumns struct {
	Id                string //
	ContinueSignDay   string // 连续签到天数
	ContinueSignPoint string // 连续签到赠送数量
	ConsumePerPoint   string // 每消费多少元获取1个点
	LowOrderAmount    string // 最低获取点数的订单金额
	MaxPointPerOrder  string // 每笔订单最高获取点数
	Type              string // 类型：0->积分规则；1->成长值规则
}

// umsMemberRuleSettingColumns holds the columns for the table ums_member_rule_setting.
var umsMemberRuleSettingColumns = UmsMemberRuleSettingColumns{
	Id:                "id",
	ContinueSignDay:   "continue_sign_day",
	ContinueSignPoint: "continue_sign_point",
	ConsumePerPoint:   "consume_per_point",
	LowOrderAmount:    "low_order_amount",
	MaxPointPerOrder:  "max_point_per_order",
	Type:              "type",
}

// NewUmsMemberRuleSettingDao creates and returns a new DAO object for table data access.
func NewUmsMemberRuleSettingDao(handlers ...gdb.ModelHandler) *UmsMemberRuleSettingDao {
	return &UmsMemberRuleSettingDao{
		group:    "default",
		table:    "ums_member_rule_setting",
		columns:  umsMemberRuleSettingColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UmsMemberRuleSettingDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UmsMemberRuleSettingDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UmsMemberRuleSettingDao) Columns() UmsMemberRuleSettingColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UmsMemberRuleSettingDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UmsMemberRuleSettingDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *UmsMemberRuleSettingDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
