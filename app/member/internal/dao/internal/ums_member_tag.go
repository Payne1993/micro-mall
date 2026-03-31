// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UmsMemberTagDao is the data access object for the table ums_member_tag.
type UmsMemberTagDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  UmsMemberTagColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// UmsMemberTagColumns defines and stores column names for the table ums_member_tag.
type UmsMemberTagColumns struct {
	Id                string //
	Name              string //
	FinishOrderCount  string // 自动打标签完成订单数量
	FinishOrderAmount string // 自动打标签完成订单金额
}

// umsMemberTagColumns holds the columns for the table ums_member_tag.
var umsMemberTagColumns = UmsMemberTagColumns{
	Id:                "id",
	Name:              "name",
	FinishOrderCount:  "finish_order_count",
	FinishOrderAmount: "finish_order_amount",
}

// NewUmsMemberTagDao creates and returns a new DAO object for table data access.
func NewUmsMemberTagDao(handlers ...gdb.ModelHandler) *UmsMemberTagDao {
	return &UmsMemberTagDao{
		group:    "default",
		table:    "ums_member_tag",
		columns:  umsMemberTagColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UmsMemberTagDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UmsMemberTagDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UmsMemberTagDao) Columns() UmsMemberTagColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UmsMemberTagDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UmsMemberTagDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *UmsMemberTagDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
