// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UmsMemberReceiveAddressDao is the data access object for the table ums_member_receive_address.
type UmsMemberReceiveAddressDao struct {
	table    string                         // table is the underlying table name of the DAO.
	group    string                         // group is the database configuration group name of the current DAO.
	columns  UmsMemberReceiveAddressColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler             // handlers for customized model modification.
}

// UmsMemberReceiveAddressColumns defines and stores column names for the table ums_member_receive_address.
type UmsMemberReceiveAddressColumns struct {
	Id            string //
	MemberId      string //
	Name          string // 收货人名称
	PhoneNumber   string //
	DefaultStatus string // 是否为默认
	PostCode      string // 邮政编码
	Province      string // 省份/直辖市
	City          string // 城市
	Region        string // 区
	DetailAddress string // 详细地址(街道)
}

// umsMemberReceiveAddressColumns holds the columns for the table ums_member_receive_address.
var umsMemberReceiveAddressColumns = UmsMemberReceiveAddressColumns{
	Id:            "id",
	MemberId:      "member_id",
	Name:          "name",
	PhoneNumber:   "phone_number",
	DefaultStatus: "default_status",
	PostCode:      "post_code",
	Province:      "province",
	City:          "city",
	Region:        "region",
	DetailAddress: "detail_address",
}

// NewUmsMemberReceiveAddressDao creates and returns a new DAO object for table data access.
func NewUmsMemberReceiveAddressDao(handlers ...gdb.ModelHandler) *UmsMemberReceiveAddressDao {
	return &UmsMemberReceiveAddressDao{
		group:    "default",
		table:    "ums_member_receive_address",
		columns:  umsMemberReceiveAddressColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UmsMemberReceiveAddressDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UmsMemberReceiveAddressDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UmsMemberReceiveAddressDao) Columns() UmsMemberReceiveAddressColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UmsMemberReceiveAddressDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UmsMemberReceiveAddressDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *UmsMemberReceiveAddressDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
