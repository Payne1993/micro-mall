// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UmsMemberDao is the data access object for the table ums_member.
type UmsMemberDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  UmsMemberColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// UmsMemberColumns defines and stores column names for the table ums_member.
type UmsMemberColumns struct {
	Id                    string //
	MemberLevelId         string //
	Username              string // 用户名
	Password              string // 密码
	Nickname              string // 昵称
	Phone                 string // 手机号码
	Status                string // 帐号启用状态:0->禁用；1->启用
	CreateTime            string // 注册时间
	Icon                  string // 头像
	Gender                string // 性别：0->未知；1->男；2->女
	Birthday              string // 生日
	City                  string // 所做城市
	Job                   string // 职业
	PersonalizedSignature string // 个性签名
	SourceType            string // 用户来源
	Integration           string // 积分
	Growth                string // 成长值
	LuckeyCount           string // 剩余抽奖次数
	HistoryIntegration    string // 历史积分数量
}

// umsMemberColumns holds the columns for the table ums_member.
var umsMemberColumns = UmsMemberColumns{
	Id:                    "id",
	MemberLevelId:         "member_level_id",
	Username:              "username",
	Password:              "password",
	Nickname:              "nickname",
	Phone:                 "phone",
	Status:                "status",
	CreateTime:            "create_time",
	Icon:                  "icon",
	Gender:                "gender",
	Birthday:              "birthday",
	City:                  "city",
	Job:                   "job",
	PersonalizedSignature: "personalized_signature",
	SourceType:            "source_type",
	Integration:           "integration",
	Growth:                "growth",
	LuckeyCount:           "luckey_count",
	HistoryIntegration:    "history_integration",
}

// NewUmsMemberDao creates and returns a new DAO object for table data access.
func NewUmsMemberDao(handlers ...gdb.ModelHandler) *UmsMemberDao {
	return &UmsMemberDao{
		group:    "default",
		table:    "ums_member",
		columns:  umsMemberColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UmsMemberDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UmsMemberDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UmsMemberDao) Columns() UmsMemberColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UmsMemberDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UmsMemberDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *UmsMemberDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
