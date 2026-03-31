// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// UmsMemberLevel is the golang structure of table ums_member_level for DAO operations like Where/Data.
type UmsMemberLevel struct {
	g.Meta                `orm:"table:ums_member_level, do:true"`
	Id                    any //
	Name                  any //
	GrowthPoint           any //
	DefaultStatus         any // 是否为默认等级：0->不是；1->是
	FreeFreightPoint      any // 免运费标准
	CommentGrowthPoint    any // 每次评价获取的成长值
	PriviledgeFreeFreight any // 是否有免邮特权
	PriviledgeSignIn      any // 是否有签到特权
	PriviledgeComment     any // 是否有评论获奖励特权
	PriviledgePromotion   any // 是否有专享活动特权
	PriviledgeMemberPrice any // 是否有会员价格特权
	PriviledgeBirthday    any // 是否有生日特权
	Note                  any //
}
