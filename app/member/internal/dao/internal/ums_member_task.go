// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UmsMemberTaskDao is the data access object for the table ums_member_task.
type UmsMemberTaskDao struct {
	table    string               // table is the underlying table name of the DAO.
	group    string               // group is the database configuration group name of the current DAO.
	columns  UmsMemberTaskColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler   // handlers for customized model modification.
}

// UmsMemberTaskColumns defines and stores column names for the table ums_member_task.
type UmsMemberTaskColumns struct {
	Id           string //
	Name         string //
	Growth       string // 赠送成长值
	Intergration string // 赠送积分
	Type         string // 任务类型：0->新手任务；1->日常任务
}

// umsMemberTaskColumns holds the columns for the table ums_member_task.
var umsMemberTaskColumns = UmsMemberTaskColumns{
	Id:           "id",
	Name:         "name",
	Growth:       "growth",
	Intergration: "intergration",
	Type:         "type",
}

// NewUmsMemberTaskDao creates and returns a new DAO object for table data access.
func NewUmsMemberTaskDao(handlers ...gdb.ModelHandler) *UmsMemberTaskDao {
	return &UmsMemberTaskDao{
		group:    "default",
		table:    "ums_member_task",
		columns:  umsMemberTaskColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UmsMemberTaskDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UmsMemberTaskDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UmsMemberTaskDao) Columns() UmsMemberTaskColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UmsMemberTaskDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UmsMemberTaskDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *UmsMemberTaskDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
