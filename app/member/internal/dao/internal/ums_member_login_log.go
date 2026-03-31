// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UmsMemberLoginLogDao is the data access object for the table ums_member_login_log.
type UmsMemberLoginLogDao struct {
	table    string                   // table is the underlying table name of the DAO.
	group    string                   // group is the database configuration group name of the current DAO.
	columns  UmsMemberLoginLogColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler       // handlers for customized model modification.
}

// UmsMemberLoginLogColumns defines and stores column names for the table ums_member_login_log.
type UmsMemberLoginLogColumns struct {
	Id         string //
	MemberId   string //
	CreateTime string //
	Ip         string //
	City       string //
	LoginType  string // 登录类型：0->PC；1->android;2->ios;3->小程序
	Province   string //
}

// umsMemberLoginLogColumns holds the columns for the table ums_member_login_log.
var umsMemberLoginLogColumns = UmsMemberLoginLogColumns{
	Id:         "id",
	MemberId:   "member_id",
	CreateTime: "create_time",
	Ip:         "ip",
	City:       "city",
	LoginType:  "login_type",
	Province:   "province",
}

// NewUmsMemberLoginLogDao creates and returns a new DAO object for table data access.
func NewUmsMemberLoginLogDao(handlers ...gdb.ModelHandler) *UmsMemberLoginLogDao {
	return &UmsMemberLoginLogDao{
		group:    "default",
		table:    "ums_member_login_log",
		columns:  umsMemberLoginLogColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UmsMemberLoginLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UmsMemberLoginLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UmsMemberLoginLogDao) Columns() UmsMemberLoginLogColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UmsMemberLoginLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UmsMemberLoginLogDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *UmsMemberLoginLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
